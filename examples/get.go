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

	entry, err := client.GetEntry(gonvs.GetEntryTask{
		Name: "dns:example",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("value:" + entry.Value)
}
