// +build authenticator

package main

import (
	"github.com/jdelgad/goforum/authenticator"
)

func main() {

	s := authenticator.SetupSocket("tcp://127.0.0.1:13000")
	defer s.Close()

	authenticator.ServiceRequests(s)
}
