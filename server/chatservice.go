package server

import (
	"context"
	"fmt"
	"time"

	"github.com/Sai-Madhur-CH/gRPC-Master/proto/chatpb"
)

// ChatServer -
type ChatServer struct {
}

// Send -
func (s *ChatServer) Send(c context.Context, req *chatpb.ChatRequest) (*chatpb.ChatResponse, error) {
	fmt.Printf("Send Function us invoked with :%v \n", req)
	msg := req.GetReq()
	msg.Date = time.Now().Format(time.RFC1123)
	res := &chatpb.ChatResponse{Res: msg}
	return res, nil
}
