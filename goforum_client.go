package main

import (
	"github.com/jdelgad/goforum/transport"
	"log"
)

func main() {
	s, err := transport.OpenConnectSocket("tcp://127.0.0.1:4000")

	if err != nil {
		log.Fatal("Could not open port for outgoing requests on port 4000")
	}

	err = transport.SendReply(s, []byte("request"))

	if err != nil {
		log.Fatal("could not send outgoing request")
	}

	transport.CloseListenSocket(s)
}
