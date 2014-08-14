package main

import (
	"log"
	"os"

	"github.com/melraidin/iptc"
)

func main() {
	data, err := iptc.Open(os.Args[1])

	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	log.Printf("%v\n", data)
}
