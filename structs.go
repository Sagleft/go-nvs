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

// [{"name":"test:0001","value":"entry value","txid":"bc50ad44b665ec05f4a42138181459bfb6c9357cedd973c99f1cc54254bc1315","address":"Edeve4DB1tn7epp796HV7WWWN3gHyzqqKy","expires_in":5250,"expires_at":538063,"time":1648763104}]

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
