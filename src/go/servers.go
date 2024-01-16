package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/taxfyle/lb-issue-repro/src/go/pb"
)

type demoServer struct {
	pb.UnimplementedDemoServer

	messages <-chan string
}

func (d *demoServer) StreamMessages(request *pb.StreamMessagesRequest, stream pb.Demo_StreamMessagesServer) error {
	log.Printf("[%s] got request named %s", time.Now(), request.Name)

	for message := range d.messages {
		log.Printf("[%s] got message", time.Now())

		err := stream.Send(&pb.StreamMessagesResponse{
			Message: message,
		})
		if err != nil {
			return err
		}

		log.Printf("[%s] message sent", time.Now())
	}

	return nil
}

type controlServer struct {
	messages chan<- string
}

func (h *controlServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	t := time.Now()
	log.Printf("[%s] got request at %s", time.Now(), t)
	h.messages <- fmt.Sprintf("Server time is %s", t)
	log.Printf("[%s] message sent on channel", time.Now())
}
