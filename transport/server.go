package transport

import (
	"github.com/gdamore/mangos"
	"github.com/gdamore/mangos/protocol/rep"
	"github.com/gdamore/mangos/transport/tcp"
	"log"
	"os"
)

func OpenListenSocket(address string) {
	socket, err := rep.NewSocket()

	if err != nil {
		log.Fatal("Could not create reply socket")
		os.Exit(-1)
	}

	socket.AddTransport(tcp.NewTransport())
	err = socket.Listen(address)
	if err != nil {
		log.Fatal("Could not open request socket on address: %s", address)
		os.Exit(-1)
	}
}

func CloseListenSocket(s mangos.Socket) {
	s.Close()
}

func ParseRequest(s mangos.Socket) ([]byte, error) {
	return s.Recv()
}

func SendReply(s mangos.Socket, reply []byte) error {
	return s.Send(reply)
}
