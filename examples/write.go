package main

import (
	"log"
	"os"

	"github.com/google/uuid"
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
		Name:  "test:" + uuid.NewString(),
		Value: []byte("entry value"),
		Days:  30,
	})
	if err != nil {
		log.Fatalln(err)
	}
}
