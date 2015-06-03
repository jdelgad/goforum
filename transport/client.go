package transport

import (
	"github.com/gdamore/mangos"
	"github.com/gdamore/mangos/protocol/req"
	"github.com/gdamore/mangos/transport/tcp"
	"log"
	"os"
)

func OpenConnectSocket(address string) {
	socket, err := req.NewSocket()

	if err != nil {
		log.Fatal("Could not get request socket")
		os.Exit(-1)
	}

	socket.AddTransport(tcp.NewTransport())
	err = socket.Dial(address)
	if err != nil {
		log.Fatal("Could not open request socket on address: %s", address)
		os.Exit(-1)
	}
}

func ParseReply(s mangos.Socket) ([]byte, error) {
	return s.Recv()
}
