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
	Host        string `json:"host"` // default: 127.0.0.1
	RPCUser     string `json:"user"`
	RPCPassword string `json:"password"`
	RPCPort     string `json:"port"` // default: 6662
	UseSSL      bool   `json:"useSSL"`
}

// ValueType - NVS value type
type ValueType string

const (
	// ValueTypeHex - hex value
	ValueTypeHex string = "hex"

	// ValueTypeBase64 - base64 value
	ValueTypeBase64 string = "base64"

	// ValueTypePlain - plain (unicode) string. by default
	ValueTypePlain string = ""
)

// WriteEntryTask - new blockchain NVS entry data
type WriteEntryTask struct {
	// required
	Name  string `json:"name"`
	Value []byte `json:"value"`
	Days  int    `json:"days"`

	// optional
	ToAddress string    `json:"toAddress"`
	ValueType ValueType `json:"type"`
}

// GetEntrysByAddressTask - get NVS entrys by given address, data
type GetEntrysByAddressTask struct {
	// required
	Address string `json:"address"`

	// optional
	MaxValueLength int       `json:"maxValueLength"` // bytes
	ValueType      ValueType `json:"type"`
}

// Entry - NVS Entry
type Entry struct {
	Name      string
	Value     string
	TXID      string
	Address   string
	ExpiresIn int
	ExpiresAt int
	Time      int64
}

// GetEntryTask - get NVS entry task
type GetEntryTask struct {
	// required
	Name string `json:"name"`

	// optional
	ValueType ValueType `json:"type"`
	Filepath  string    `json:"filepath"`
}
