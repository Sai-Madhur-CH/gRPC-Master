package main

import (
	"github.com/Sai-Madhur-CH/gRPC-Master/src/client"
	"github.com/Sai-Madhur-CH/gRPC-Master/src/server"
)

func main() {
	server.Run()
	client.Run()
}
