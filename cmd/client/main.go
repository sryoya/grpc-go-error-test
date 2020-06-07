// package main is a test grpc client to test auth code generation server
package main

import (
	"context"
	"errors"
	"log"
	"time"

	pb "github.com/sryoya/grpc-go-error-test/test"
	"google.golang.org/grpc"
)

// main stars grpc client and send request
func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTestClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Handle(ctx, &pb.Req{})
	if err != nil {
		if errors.Is(err, pb.CustomErr) {
			log.Printf("handled expected error")
			return
		}

		log.Fatalf("unexpected error: %+v", err)
	}
	log.Printf("Result: %v", r)

}
