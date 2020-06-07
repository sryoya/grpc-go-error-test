package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/sryoya/grpc-go-error-test/test"
)

func main() {
	// initialize clearinghandler server
	server := &test.Server{}

	// listen tcp port
	lis, err := net.Listen("tcp", "localhost:9080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	s := grpc.NewServer()

	// connect the grpc server and the clearinghandler service
	test.RegisterTestServer(s, server)
	// start the server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
