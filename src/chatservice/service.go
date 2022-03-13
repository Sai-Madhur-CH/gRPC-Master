package chatservice

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"github.com/Sai-Madhur-CH/gRPC-Master/src/chatpb"
)

var db *gorm.DB

// Server -
type Server struct {
}

// Send -
func (s *Server) Send(c context.Context, req *chatpb.ChatRequest) (*chatpb.ChatResponse, error) {
	fmt.Printf("Send Function us invoked with :%v \n", req)
	msg := req.GetReq()
	res := &chatpb.ChatResponse{Res: msg}
	return res, nil
}
