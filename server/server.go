package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/Sai-Madhur-CH/gRPC-Master/proto/chatpb"
	"github.com/Sai-Madhur-CH/gRPC-Master/util"
)

var db *gorm.DB

// Run - start the gRPC Server
func Run() {
	// In case of run time errors show file name and line number in the code
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Initiated Chat Server:")

	// DB Connection
	db = util.Connect()
	// TCP listener
	listner, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Error while listining to TCP: %v", err)
	}

	// gRPC new server
	s := grpc.NewServer()

	// Register the service
	RegisterServices(s)

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

// RegisterServices - Register Existing Services
func RegisterServices(s *grpc.Server) {
	chatpb.RegisterChatServiceServer(s, &ChatServer{})
}
