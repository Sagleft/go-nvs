package gonvs

import (
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
	requestData := []interface{}{
		task.Name,
		string(task.Value),
		task.Days,
	}
	if task.ToAddress != "" {
		requestData = append(requestData, task.ToAddress)
	}
	if task.ValueType != "" {
		requestData = append(requestData, string(task.ValueType))
	}

	var result interface{}
	err := c.sendRequest("name_new", &result, requestData)
	if err != nil {
		return err
	}
	return nil
}

// GetEntrysByAddress - get NVS entrys by given EMC address
func (c *Client) GetEntrysByAddress(task GetEntrysByAddressTask) ([]Entry, error) {
	requestData := []interface{}{
		task.Address,
	}
	if task.MaxValueLength > 0 {
		requestData = append(requestData, task.MaxValueLength)
	}
	if task.ValueType != "" {
		requestData = append(requestData, string(task.ValueType))
	}

	result := []Entry{}
	err := c.sendRequest("name_scan_address", &result, requestData)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteEntry - delete NVS entry
func (c *Client) DeleteEntry(entryName string) error {
	var result interface{}
	return c.sendRequest("name_delete", &result, []interface{}{entryName})
}
