package authenticator

import (
	"errors"
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

func CreateReply() *protos.LoginReply {
	return &protos.LoginReply{}
}

func CreateRequest() *protos.Login {
	return &protos.Login{}
}

func parseRequest(b []byte) (*protos.Login, error) {
	login := CreateRequest()
	err := proto.Unmarshal(b, login)

	if err != nil {
		return login, errors.New("could not parse login protobuf")
	}

	return login, nil
}

func isValidLogin(l protos.Login) bool {
	users, err := GetUserPasswordList("passwd")

	if err != nil {
		log.Fatal("Could not retrieve user password list")
	}

	return IsRegisteredUser(*l.Username, users) && IsPasswordValid(*l.Username, *l.Password, users)
}

func authFailure(r *protos.LoginReply) {
	authorization := protos.LoginReply_FAILED
	r.Authorized = &authorization
	sid := "-1"
	r.SessionID = &sid
}

func authSuccess(r *protos.LoginReply) {
	authorization := protos.LoginReply_SUCCESSFUL
	r.Authorized = &authorization
	sid := "1"
	r.SessionID = &sid
}

func sendReply(r *protos.LoginReply, s *transport.ServerSocket) {
	b, err := proto.Marshal(r)

	if err != nil {
		log.Fatal("could not serialize reply")
	}

	s.Send(b)
}

func ServiceRequests(s *transport.ServerSocket) {
	for {
		b, err := s.Receive()

		if err != nil {
			log.Fatal("could not parse incoming request")
		}

		l, err := parseRequest(b)

		r := CreateReply()
		if err != nil {
			authFailure(r)
		} else {

			if isValidLogin(*l) {
				authSuccess(r)
			} else {
				authFailure(r)
			}
		}

		sendReply(r, s)
	}
}
