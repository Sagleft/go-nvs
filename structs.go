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
	Name      string `json:"name"`
	Value     string `json:"value"`
	TXID      string `json:"txid"`
	Address   string `json:"address"`
	ExpiresIn int    `json:"expiresIn"`
	ExpiresAt int    `json:"expiresAt"`
	Time      int64  `json:"time"`
}

// GetEntryTask - get NVS entry task
type GetEntryTask struct {
	// required
	Name string `json:"name"`

	// optional
	ValueType ValueType `json:"type"`
	Filepath  string    `json:"filepath"`
}

// GetEntryHistoryTask - get NVS entry history
type GetEntryHistoryTask struct {
	// required
	Name string `json:"name"`

	// optional
	LoadFullHistory bool      `json:"fullhistory"`
	ValueType       ValueType `json:"type"`
}

type rawTXInput struct {
	TransactionID string `json:"txid"`
	VOut          int    `json:"vout"`
}

type unspentData struct {
	TransactionID string  `json:"txid"`
	VOut          int     `json:"vout"`
	Address       string  `json:"address"`
	Account       string  `json:"account"`
	ScriptPubKey  string  `json:"scriptPubKey"`
	Amount        float64 `json:"amount"`
	Confirmations int     `json:"confirmations"`
	Spendable     bool    `json:"spendable"`
	Solvable      bool    `json:"solvable"`
}

type signRawTXResponse struct {
	Hex      string `json:"hex"`
	Complete bool   `json:"complete"`
}
