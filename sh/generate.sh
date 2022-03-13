#!/bin/bash

export PATH="$PATH:$(go env GOPATH)/bin"
protoc proto/userpb/user.proto --go_out=plugins=grpc:.
protoc proto/chatpb/chat.proto --go_out=plugins=grpc:.
