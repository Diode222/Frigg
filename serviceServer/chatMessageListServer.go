package serviceServer

import (
	"context"
	"github.com/Diode222/Frigg/db"
	"github.com/Diode222/Frigg/model"
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
	dbInstance := db.GetDB()

	log.Println("frigg get chatMessageList length: ", len(chatMessageList.GetChatMessages()))

	for _, chatMessage := range chatMessageList.GetChatMessages() {
		msg := chatMessage.GetMessage()
		time := chatMessage.GetTime()
		chatPerson := chatMessage.GetChatPerson()

		log.Println("get chat messsage wordAndPosList length: ", len(chatMessage.GetWordAndPosList()))

		for _, wordAndPos := range chatMessage.GetWordAndPosList() {
			word := wordAndPos.GetWord()
			pos := model.UNKNOWN
			switch wordAndPos.GetPos().GetType() {
			case pb.PartOfSpeech_NOUN:
				pos = model.NOUN
			case pb.PartOfSpeech_VERB:
				pos = model.VERB
			case pb.PartOfSpeech_ADJECTIVE:
				pos = model.ADJECTIVE
			case pb.PartOfSpeech_PHRASE:
				pos = model.PHRASE
			}

			wordFreqItem := model.WordFreqItem{
				Word:       word,
				Pos:        pos,
				ChatPerson: chatPerson,
				Sentence:   msg,
				SendTime:   time,
			}
			errs := dbInstance.Mysql.Create(&wordFreqItem).GetErrors()
			if len(errs) > 0 {
				log.Printf("Mysql insert got errors, ", errs)
			}
		}
	}

	ok := true
	return &pb.ChatMessageListServiceStatus{
		OK: &ok,
	}, nil
}
