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

// ValueType - NVS value type
type ValueType string

const (
	// ValueTypeHex - hex value
	ValueTypeHex string = "hex"

	// ValueTypeBase64 - base64 value
	ValueTypeBase64 string = "base64"

	// ValueTypePlain - plain value
	ValueTypePlain string = ""
)

// WriteEntryTask - new blockchain NVS entry data
type WriteEntryTask struct {
	//required
	Name  string `json:"name"`
	Value []byte `json:"value"`
	Days  int    `json:"days"`

	// optional
	ToAddress string    `json:"toAddress"`
	ValueType ValueType `json:"type"`
}
