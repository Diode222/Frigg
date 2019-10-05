package chatMessageListServer

import (
	"context"
	pb "github.com/Diode222/Frigg/proto_gen"
	"sync"
)

type wordFreqListServer struct {}

var serverWordFreqListServer *wordFreqListServer
var wordFreqListServerOnce sync.Once

func NewWordFreqListServer() *wordFreqListServer {
	wordFreqListServerOnce.Do(func() {
		serverWordFreqListServer = &wordFreqListServer{}
	})
	return serverWordFreqListServer;
}

func (s *wordFreqListServer) GetWordFreqList(context.Context, *pb.PartOfSpeech) (*pb.WordFreqList, error) {
	// TODO request from MongoDB
	word1 := "nono"
	var count1 int32 = 50
	word2 := "masiwei"
	var count2 int32 = 30
	return &pb.WordFreqList{
		WordFreqs: []*pb.WordFreq{
			&pb.WordFreq{
				Word:                 &word1,
				Count:                &count1,
			},
			&pb.WordFreq{
				Word:                 &word2,
				Count:                &count2,
			},
		},
	}, nil
}
