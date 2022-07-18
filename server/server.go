package server

import (
	"log"

	api "github.com/mponce/myproto/go-proto"
	"golang.org/x/net/context"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, in *api.PingMessage) (*api.PingMessage, error) {
	log.Printf("Receive message %s", in.Greeting)
	return &api.PingMessage{Greeting: "bar"}, nil
}
