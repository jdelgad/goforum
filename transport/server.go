package transport

import (
	"errors"
	"github.com/gdamore/mangos/protocol/rep"
	"github.com/gdamore/mangos/transport/tcp"
)

func NewServerSocket() *ServerSocket {
	return &ServerSocket{}
}

func (s *ServerSocket) Open() error {
	socket, err := rep.NewSocket()

	if err != nil {
		return errors.New("Could not create reply socket")
	}

	socket.AddTransport(tcp.NewTransport())
	s.socket = socket
	return nil
}

func (s *ServerSocket) Connect(address string) error {
	return s.socket.Listen(address)
}

func (s *ServerSocket) Receive() ([]byte, error) {
	return s.socket.Recv()
}

func (s *ServerSocket) Send(data []byte) error {
	return s.socket.Send(data)
}

func (s *ServerSocket) Close() error {
	return s.socket.Close()
}
