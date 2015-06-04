package transport

import (
	"github.com/gdamore/mangos"
)

type Socketer interface {
	Open() error
	Connect(address string) error
	Receive() ([]byte, error)
	Send(data []byte) error
	Close() error
}

type ClientSocket struct {
	socket mangos.Socket
}

type ServerSocket struct {
	socket mangos.Socket
}
