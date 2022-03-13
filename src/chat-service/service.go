package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"

	"github.com/Sai-Madhur-CH/gRPC-Master/src/chatpb"
)

type server struct {
}

func (s *server) Send(c context.Context, req *chatpb.ChatRequest) (*chatpb.ChatResponse, error) {
	fmt.Printf("Send Function us invoked with :%v \n", req)
	msg := req.GetReq()
	res := &chatpb.ChatResponse{Res: msg}
	return res, nil
}

func main() {
	// In case of run time errors show file name and line number in the code
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Initiated Chat Server:")

	// TCP listener
	listner, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Error while listining to TCP: %v", err)
	}

	// gRPC new server
	s := grpc.NewServer()

	// Register the service
	chatpb.RegisterChatServiceServer(s, &server{})

	// gRPC Serving
	go func() {
		fmt.Println("Starting the Server...")
		if err := s.Serve(listner); err != nil {
			log.Fatalf("Error while serving: %v", err)
		}
	}()

	// Wait for Control + c to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	// Stopping the server
	fmt.Println("Stopping the Server.")
	s.Stop()

	// Closing the listner
	fmt.Println("Closing the Listener.")
	listner.Close()

	fmt.Println("Server is Down.")
}
