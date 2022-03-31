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
	// TODO
	return nil
}
