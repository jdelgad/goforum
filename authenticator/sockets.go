package authenticator

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/jdelgad/goforum/protos"
	"github.com/jdelgad/goforum/transport"
	"log"
)

func SetupSocket(address string) *transport.ServerSocket {
	s := transport.NewServerSocket()
	s.Open()
	s.Connect(address)
	return s
}

func CreateReply(success bool) *protos.LoginReply {
	return &protos.LoginReply{}
}

func ParseLogin(b []byte) (*protos.Login, error) {
	login := &protos.Login{}
	err := proto.Unmarshal(b, login)

	if err != nil {
		return login, errors.New("could not parse login protobuf")
	}

	return login, nil
}

func isValidUser(l protos.Login) {
}

func ListenForRequest(s *transport.ServerSocket) {
	for {
		b, err := s.Receive()

		if err != nil {
			log.Fatal("could not parse incoming request")
		}

		fmt.Println(string(b))
	}

}
