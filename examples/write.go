package main

import (
	"log"
	"os"

	gonvs ".."
)

func main() {
	client := gonvs.NewClient(gonvs.CreateClientTask{
		RPCUser:     os.Getenv("USER"),
		RPCPassword: os.Getenv("PASSWORD"),
	})

	err := client.Write(gonvs.WriteEntryTask{})
	if err != nil {
		log.Fatalln(err)
	}
}
