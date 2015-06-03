package main

import (
	"fmt"
	"github.com/jdelgad/goforum/transport"
	"log"
)

func main() {
	s, err := transport.OpenListenSocket("tcp://127.0.0.1:4000")

	if err != nil {
		log.Fatal("Could not listen for incoming requests on port 4000")
	}

	for {
		b, err := transport.ParseRequest(s)

		if err != nil {
			log.Fatal("could not parse incoming request")
		}

		fmt.Println(string(b))
	}

	transport.CloseListenSocket(s)
}
