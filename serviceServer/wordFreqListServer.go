package serviceServer

import (
	"context"
	pb "github.com/Diode222/Frigg/proto_gen"
	"log"
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

func (s *wordFreqListServer) GetWordFreqList(contex context.Context, pos *pb.PartOfSpeech) (*pb.WordFreqList, error) {
	// TODO request from db
	log.Println("获取到wordfreqlist请求")
	if pos.GetType() == pb.PartOfSpeech_NOUN {
		word1 := "noun"
		var count1 int32 = 50
		word2 := "名词有点尴尬"
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
	} else {
		word1 := "other, pos"
		var count1 int32 = 50
		word2 := "其他词性有点尴尬"
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
}
