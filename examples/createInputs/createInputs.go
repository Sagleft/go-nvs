package main

import (
	"log"

	gonvs "github.com/sagleft/go-nvs"
)

func main() {
	client, err := gonvs.NewClient(gonvs.CreateClientTask{
		RPCUser:     "emcuser",
		RPCPassword: "emcpass",
	})
	if err != nil {
		log.Fatalln(err)
	}

	addresses, err := client.GetAccountAddresses()
	if err != nil {
		log.Fatalln(err)
	}

	if err := client.CreateNewInputs(
		0.1,
		addresses,
	); err != nil {
		log.Fatalln(err)
	}
}
