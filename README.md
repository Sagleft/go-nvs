
![logo](logo.png)

Emercoin NVS lib written in Go

[![GoDoc](https://godoc.org/github.com/sagleft/go-nvs?status.svg)](https://godoc.org/gopkg.in/sagleft/go-nvs.v1)
[![go-report](https://goreportcard.com/badge/github.com/Sagleft/go-nvs)](https://goreportcard.com/report/github.com/Sagleft/go-nvs)

-----

## Install

```bash
go get github.com/sagleft/go-nvs
```

## Usage example

```go
client, err := gonvs.NewClient(gonvs.CreateClientTask{
	RPCUser:     os.Getenv("USER"),
	RPCPassword: os.Getenv("PASSWORD"),
})
if err != nil {
	log.Fatalln(err)
}

err = client.Write(gonvs.WriteEntryTask{
	Name:  "test:" + uuid.NewString(),
	Value: []byte("entry value"),
	Days:  30,
})
if err != nil {
	log.Fatalln(err)
}

```
