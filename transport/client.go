package transport

import (
	"errors"
	"github.com/gdamore/mangos/protocol/req"
	"github.com/gdamore/mangos/transport/tcp"
)

func NewClientSocket() *ClientSocket {
	return &ClientSocket{}
}

func (s *ClientSocket) Open() error {
	sock, err := req.NewSocket()

	if err != nil {
		return errors.New("Could not create socket")
	}

	sock.AddTransport(tcp.NewTransport())
	s.socket = sock
	return nil
}

func (s *ClientSocket) Connect(address string) error {
	return s.socket.Dial(address)
}

func (s *ClientSocket) Receive() ([]byte, error) {
	return s.socket.Recv()
}

func (s *ClientSocket) Send(data []byte) error {
	return s.socket.Send(data)
}

func (s *ClientSocket) Close() error {
	return s.socket.Close()
}
