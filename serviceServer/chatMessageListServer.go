package chatMessageListServer

import (
	"context"
	pb "github.com/Diode222/Frigg/proto_gen"
	"log"
	"sync"
)

type chatMessageListServer struct {}

var serverChatMessageList *chatMessageListServer
var chatMessageListServerOnce sync.Once

func NewChatMessageListServer() *chatMessageListServer {
	chatMessageListServerOnce.Do(func() {
		serverChatMessageList = &chatMessageListServer{}
	})
	return serverChatMessageList;
}

func (s *chatMessageListServer) PutChatMessageList(context.Context, *pb.ChatMessageList) (*pb.ChatMessageListServiceStatus, error) {
	// TODO save data to MongoDB
	log.Printf("PutChatMessageList 收到了请求")
	ok := true
	return &pb.ChatMessageListServiceStatus{
		OK: &ok,
	}, nil
}
