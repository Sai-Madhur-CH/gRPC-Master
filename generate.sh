export PATH="$PATH:$(go env GOPATH)/bin"
protoc src/userpb/user.proto --go_out=plugins=grpc:.
protoc src/chatpb/chat.proto --go_out=plugins=grpc:.
chmod +x gorm.sh
./gorm.sh src/chatpb/chat.pb.go