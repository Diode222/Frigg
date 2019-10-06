package serviceServer

import (
	"context"
	"fmt"
	"github.com/Diode222/Frigg/conf"
	"github.com/Diode222/Frigg/db"
	"github.com/Diode222/Frigg/model"
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
	wordFreqList := &pb.WordFreqList{
		WordFreqs: []*pb.WordFreq{},
	}
	dbInstance := db.GetDB()

	posType := model.UNKNOWN
	switch pos.GetType() {
	case pb.PartOfSpeech_NOUN:
		posType = model.NOUN
	case pb.PartOfSpeech_VERB:
		posType = model.VERB
	case pb.PartOfSpeech_ADJECTIVE:
		posType = model.ADJECTIVE
	case pb.PartOfSpeech_PHRASE:
		posType = model.PHRASE
	}

	rows, err := dbInstance.Mysql.Table(conf.WORDFREQ_TABLE_NAME).Select("word, count(*) as count").
		Where("pos = ?", posType).Group("word").Limit(conf.MAX_RESPONSE_WORD_COUNT).Order("count DESC").Rows()
	if err != nil {
		log.Println(fmt.Sprintf("Mysql select from table %s failed, dbname: %s", conf.WORDFREQ_TABLE_NAME, conf.MYSQL_DBNAME))
		return wordFreqList, err
	}

	var word string
	var count int32
	for rows.Next() {
		if err := rows.Scan(&word, &count); err != nil {
			fmt.Println(err.Error())
			continue
		}

		wordFreqList.WordFreqs = append(wordFreqList.WordFreqs, &pb.WordFreq{
			Word:                 &word,
			Count:                &count,
		})
	}

	return wordFreqList, nil
}
