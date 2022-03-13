package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/Sai-Madhur-CH/gRPC-Master/src/chatpb"
)

func main() {
	fmt.Println("Chat Client Initiated...")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while Dialing ther server: %v \n", err)
	}

	defer cc.Close()

	c := chatpb.NewChatServiceClient(cc)
	req := &chatpb.ChatRequest{Req: &chatpb.Msg{
		Id:         "1",
		FromUserId: "1",
		ToUserId:   "2",
		Msg:        "Hi this is Sai",
		Date:       time.Now().String(),
	}}
	res, err := c.Send(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling send Service: %v \n", err)
	}

	log.Println("Response of send:", res.Res)
}
