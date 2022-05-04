package main

import (
	"fmt"
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
}
