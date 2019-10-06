package model

import "github.com/jinzhu/gorm"

// 词性
type PartOfSpeech = int

const (
	NOUN      PartOfSpeech = 0 // 名词
	VERB      PartOfSpeech = 1 // 动词
	ADJECTIVE PartOfSpeech = 2 // 形容词
	PHRASE    PartOfSpeech = 3 // 短语
	UNKNOWN PartOfSpeech = 4 // 未知词性
)

type WordFreqItem struct {
	gorm.Model
	Word       string       `gorm:"not null;index:idx_splitted_word"`
	Pos        PartOfSpeech `gorm:"not null"`
	ChatPerson string       `gorm:"not null"` // 聊天对象
	Sentence   string       `gorm:"not null"` // 所在句子
	SendTime   int64        `gorm:"not null"` // 发送时间
}
