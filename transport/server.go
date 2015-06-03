package transport

import (
	"github.com/gdamore/mangos"
	"github.com/gdamore/mangos/protocol/rep"
	"github.com/gdamore/mangos/transport/tcp"
	"errors"
)

func OpenListenSocket(address string) (mangos.Socket, error) {
	socket, err := rep.NewSocket()

	if err != nil {
		return nil, errors.New("Could not create reply socket")
	}

	socket.AddTransport(tcp.NewTransport())
	err = socket.Listen(address)
	if err != nil {
		return nil, errors.New("Could not listen on address " + address)
	}

	return socket, nil
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
