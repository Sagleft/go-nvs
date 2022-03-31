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
	})
	if err != nil {
		log.Fatalln(err)
	}

	err = client.Write(gonvs.WriteEntryTask{
		Name:  "test:0001",
		Value: []byte("entry value"),
	})
	if err != nil {
		log.Fatalln(err)
	}
}
