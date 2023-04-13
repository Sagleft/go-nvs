package gonvs

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (c *Client) ListUnspent() ([]unspentData, error) {
	r := []unspentData{}
	if err := c.sendRequest("listunspent", &r, []interface{}{}); err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) GetAccountAddress(account string) (string, error) {
	var address string
	if err := c.sendRequest(
		"getaccountaddress",
		&address,
		[]interface{}{account},
	); err != nil {
		return "", err
	}

	return address, nil
}

func (c *Client) GetAccounts() ([]string, error) {
	balances := map[string]float64{}
	if err := c.sendRequest("listaccounts", &balances, []interface{}{}); err != nil {
		return nil, err
	}

	accounts := []string{}
	for account := range balances {
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (c *Client) GetAccountAddresses() ([]string, error) {
	accounts, err := c.GetAccounts()
	if err != nil {
		return nil, fmt.Errorf("get accounts: %w", err)
	}

	addresses := []string{}
	for _, account := range accounts {
		address, err := c.GetAccountAddress(account)
		if err != nil {
			return nil, fmt.Errorf("get account %q address: %w", account, err)
		}

		addresses = append(addresses, address)
	}

	return addresses, nil
}

func (c *Client) BatchCreateAccounts(count int) error {
	/*accounts
	c.sendRequest("", )*/
	// TODO
	return nil
}

func (c *Client) CreateNewInputs(
	toAddress string,
	inputMinAmount float64,
	inputsCount int,
	addresses []string,
) error {
	inputs, err := c.ListUnspent()
	if err != nil {
		return fmt.Errorf("list unspent: %w", err)
	}

	var input unspentData
	for i, in := range inputs {
		if in.Amount > inputMinAmount {
			input = in
			break
		}

		if i == len(inputs)-1 {
			return errors.New("suitable inputs not found")
		}
	}

	txInput := rawTXInput{
		TransactionID: input.TransactionID,
		VOut:          input.VOut,
	}
	txInputs := []rawTXInput{txInput}
	txInputsJSONBytes, err := json.Marshal(txInputs)
	if err != nil {
		return fmt.Errorf("encode tx inputs: %w", err)
	}

	outputs := map[string]float64{} // TODO
	outputsJSONBytes, err := json.Marshal(outputs)
	if err != nil {
		return fmt.Errorf("encode addresses: %w", err)
	}

	// createrawtransaction "[{\"txid\" : \"mytxid\",\"vout\":0}]" "{\"myaddress\":0.01}"

	var rawTX string
	if err := c.sendRequest("createrawtransaction", &rawTX, []interface{}{
		string(txInputsJSONBytes),
		string(outputsJSONBytes),
	}); err != nil {
		return fmt.Errorf("create raw tx: %w", err)
	}

	var signTxResponse signRawTXResponse
	if err := c.sendRequest("signrawtransaction", &signTxResponse, []interface{}{
		rawTX,
	}); err != nil {
		return fmt.Errorf("sign raw tx: %w", err)
	}

	if !signTxResponse.Complete {
		return errors.New("failed to sign raw tx: incomplete")
	}

	var txID string
	if err := c.sendRequest("sendrawtransaction", &txID, []interface{}{
		signTxResponse.Hex,
	}); err != nil {
		return fmt.Errorf("send raw tx: %w", err)
	}

	return nil
}
