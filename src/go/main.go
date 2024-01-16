package main

import (
	"log"
	"net"
	"net/http"

	"github.com/taxfyle/lb-issue-repro/src/go/pb"
	"google.golang.org/grpc"
)

func runDemoServer(messages <-chan string) error {
	log.Println("booting demo server...")

	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// TODO: TLS
	var opts []grpc.ServerOption
	// if *tls {
	// 	if *certFile == "" {
	// 		*certFile = data.Path("x509/server_cert.pem")
	// 	}
	// 	if *keyFile == "" {
	// 		*keyFile = data.Path("x509/server_key.pem")
	// 	}
	// 	creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
	// 	if err != nil {
	// 		log.Fatalf("Failed to generate credentials: %v", err)
	// 	}
	// 	opts = []grpc.ServerOption{grpc.Creds(creds)}
	// }
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterDemoServer(grpcServer, &demoServer{
		messages: messages,
	})

	return grpcServer.Serve(lis)
}

func runControlServer(messages chan<- string) error {
	log.Println("booting control server...")

	s := &controlServer{
		messages: messages,
	}
	return http.ListenAndServe(":9998", s)
}

func main() {
	errch := make(chan error)
	msgch := make(chan string)

	go func(msgch <-chan string) {
		err := runDemoServer(msgch)
		if err != nil {
			log.Printf("error running demo server: %v", err)
		}

		errch <- err
	}(msgch)

	go func(msgch chan<- string) {
		err := runControlServer(msgch)
		if err != nil {
			log.Printf("error running control server: %v", err)
		}

		errch <- err
	}(msgch)

	<-errch
	log.Println("got an error, shutting down...")
}
