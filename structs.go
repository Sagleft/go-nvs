package gonvs

import (
	rpcclient "github.com/bitlum/go-bitcoind-rpc/rpcclient"
)

// Client - Emercoin NVS client
type Client struct {
	RPC *rpcclient.Client
}

// CreateClientTask - new Client data
type CreateClientTask struct {
	Host        string `json:"host"`
	RPCUser     string `json:"user"`
	RPCPassword string `json:"password"`
	RPCPort     string `json:"port"`
	UseSSL      bool   `json:"useSSL"`
}

// WriteEntryTask - new blockchain NVS entry data
type WriteEntryTask struct {
	Name  string `json:"name"`
	Value []byte `json:"value"`
}
