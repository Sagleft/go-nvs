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

	if err := client.BatchCreateAccounts(500); err != nil {
		log.Fatalln(err)
	}
}
