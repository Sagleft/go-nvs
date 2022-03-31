
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

Create client:

```go
client, err := gonvs.NewClient(gonvs.CreateClientTask{
	RPCUser:     os.Getenv("USER"),
	RPCPassword: os.Getenv("PASSWORD"),
})
if err != nil {
	log.Fatalln(err)
}
```


Write new data to the blockchain:

```go
err = client.Write(gonvs.WriteEntryTask{
	Name:  "test:" + uuid.NewString(),
	Value: []byte("entry value"),
	Days:  30,
})
if err != nil {
	log.Fatalln(err)
}

```


Get a list of records at the address:

```go
address := os.Getenv("ADDRESS")
entrys, err := client.GetEntrysByAddress(gonvs.GetEntrysByAddressTask{
	Address: address,
})
if err != nil {
	log.Fatalln(err)
}

fmt.Println("NVS entrys for address `" + address + "`:")
for _, entry := range entrys {
	fmt.Println("    " + entry.Name)
}
```

## Donate

![logo](examples/donate.png)

```
Ed1epQAF8TW5VDhZUusf4MsMNpNh9UMzm2
```
