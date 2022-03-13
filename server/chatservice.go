package server

import (
	"context"
	"fmt"

	"github.com/Sai-Madhur-CH/gRPC-Master/proto/chatpb"
)

// ChatServer -
type ChatServer struct {
}

// Send -
func (s *ChatServer) Send(c context.Context, req *chatpb.ChatRequest) (*chatpb.ChatResponse, error) {
	fmt.Printf("Send Function us invoked with :%v \n", req)
	msg := req.GetReq()
	res := &chatpb.ChatResponse{Res: msg}
	return res, nil
}
