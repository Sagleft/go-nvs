package gonvs

import (
	"encoding/json"
	"errors"
	"strconv"

	rpcclient "github.com/bitlum/go-bitcoind-rpc/rpcclient"
)

// NewClient - create new Emercoin NVS Client
func NewClient(task CreateClientTask) (*Client, error) {
	if task.Host == "" {
		task.Host = "127.0.0.1"
	}
	if task.RPCPort == "" {
		task.RPCPort = "6662"
	}

	c := Client{}

	var err error
	c.RPC, err = rpcclient.New(&rpcclient.ConnConfig{
		HTTPPostMode: true,
		DisableTLS:   !task.UseSSL,
		Host:         task.Host + ":" + task.RPCPort,
		User:         task.RPCUser,
		Pass:         task.RPCPassword,
	}, nil)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (c *Client) Write(task WriteEntryTask) error {
	requestData := []json.RawMessage{
		wrapJSONParam(task.Name),
		wrapJSONParam(string(task.Value)),
		json.RawMessage(strconv.Itoa(task.Days)),
	}
	if task.ToAddress != "" {
		requestData = append(requestData, wrapJSONParam(task.ToAddress))
	}
	if task.ValueType != "" {
		requestData = append(requestData, wrapJSONParam(string(task.ValueType)))
	}

	_, err := c.RPC.RawRequest("name_new", requestData)
	if err != nil {
		return errors.New("failed to write entry: " + err.Error())
	}
	return nil
}

// GetEntrysByAddress -
func (c *Client) GetEntrysByAddress(task GetEntrysByAddressTask) ([]Entry, error) {
	requestData := []json.RawMessage{
		wrapJSONParam(task.Address),
	}
	if task.MaxValueLength > 0 {
		requestData = append(requestData, json.RawMessage(strconv.Itoa(task.MaxValueLength)))
	}
	if task.ValueType != "" {
		requestData = append(requestData, wrapJSONParam(string(task.ValueType)))
	}

	response, err := c.RPC.RawRequest("name_scan_address", requestData)
	if err != nil {
		return nil, errors.New("failed to get entrys: " + err.Error())
	}

	// Entry
	jsonBytes, err := response.MarshalJSON()
	if err != nil {
		return nil, errors.New("failed to json encode client response: " + err.Error())
	}

	result := []Entry{}
	err = json.Unmarshal(jsonBytes, &result)
	if err != nil {
		return nil, errors.New("failed to decode json response: " + err.Error())
	}
	return result, nil
}
