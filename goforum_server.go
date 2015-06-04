// +build main3

package main

import (
	"fmt"
	"github.com/jdelgad/goforum/transport"
	"log"
)

func main() {
	s := transport.NewServerSocket()
	s.Open()
	s.Connect("tcp://127.0.0.1:4000")

	for {
		b, err := s.Receive()

		if err != nil {
			log.Fatal("could not parse incoming request")
		}

		fmt.Println(string(b))
	}

	s.Close()
}
