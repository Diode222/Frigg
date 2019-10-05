package serviceServer

import (
	"context"
	"fmt"
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

func (s *chatMessageListServer) PutChatMessageList(context context.Context, chatMessageList *pb.ChatMessageList) (*pb.ChatMessageListServiceStatus, error) {
	// TODO save data to db
	log.Printf("PutChatMessageList 收到了请求")
	for _, msg := range chatMessageList.GetChatMessages() {
		fmt.Println(msg.GetMessage())
		fmt.Println(msg.GetTime())
		fmt.Println(msg.GetChatPerson())
		fmt.Println(">>>>>>>>>>>>>")
		for _, wordAndPos := range msg.WordAndPosList {
			fmt.Println(wordAndPos.GetWord())
			fmt.Println(wordAndPos.GetPos().GetType().Enum().String())
		}
		fmt.Println(">>>>>>>>>>>>>")
	}
	ok := true
	return &pb.ChatMessageListServiceStatus{
		OK: &ok,
	}, nil
}
