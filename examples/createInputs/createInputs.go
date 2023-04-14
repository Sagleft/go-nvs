package main

import (
	"log"
	"os"

	gonvs "github.com/sagleft/go-nvs"
)

func main() {
	client, err := gonvs.NewClient(gonvs.CreateClientTask{
		RPCUser:     os.Getenv("USER"),
		RPCPassword: os.Getenv("PASSWORD"),
		RPCPort:     os.Getenv("PORT"),
	})
	if err != nil {
		log.Fatalln(err)
	}

	addresses, err := client.GetAccountAddresses()
	if err != nil {
		log.Fatalln(err)
	}

	if err := client.CreateNewInputs(
		os.Getenv("ADDRESS"),
		0.1,
		50,
		addresses,
	); err != nil {
		log.Fatalln(err)
	}
}
