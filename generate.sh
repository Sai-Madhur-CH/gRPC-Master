export PATH="$PATH:$(go env GOPATH)/bin"
protoc src/chatpb/chat.proto --go_out=plugins=grpc:.