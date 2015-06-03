package transport

import (
	"github.com/gdamore/mangos"
	"github.com/gdamore/mangos/protocol/req"
	"github.com/gdamore/mangos/transport/tcp"
	"errors"
)

func OpenConnectSocket(address string) (mangos.Socket, error) {
	socket, err := req.NewSocket()

	if err != nil {
		return nil, errors.New("Could not create socket")
	}

	socket.AddTransport(tcp.NewTransport())
	err = socket.Dial(address)
	if err != nil {
		return nil, errors.New("Could not connect to address " + address)
	}

	return socket, nil
}

func ParseReply(s mangos.Socket) ([]byte, error) {
	return s.Recv()
}
