package main

import (
	"fmt"
	pb "github.com/mponce/myproto/go-proto"
	"github.com/mponce/myrpc-server/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

// start gRPC server and wait for connection
func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9090))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := server.Server{}

	grpcServer := grpc.NewServer()

	// attach the Ping service to the server
	pb.RegisterPingServer(grpcServer, &s)

	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
