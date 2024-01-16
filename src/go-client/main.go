package main

import (
	"context"
	"log"

	"github.com/taxfyle/lb-issue-repro/src/go-client/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	creds := credentials.NewClientTLSFromCert(nil, "")
	conn, err := grpc.Dial("l7-test-demo-go.example.com:443", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewDemoClient(conn)

	stream, err := client.StreamMessages(context.Background(), &pb.StreamMessagesRequest{
		Name: "foo",
	})
	if err != nil {
		log.Fatalf("error on stream messages: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Fatalf("error receiving message: %v", err)
		}

		log.Printf("Received message: %s", msg.GetMessage())
	}
}
