package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"sync"
)

type db struct {
	Mysql *gorm.DB
}

var dbInstance *db
var dbInstanceOnce sync.Once

func InitDB(mySqlInfo string) {
	dbInstanceOnce.Do(func() {
		mysqlDB, err := gorm.Open("mysql", mySqlInfo)
		if err != nil {
			log.Printf(fmt.Sprintf("Open db instance failed when open Mysql, Mysql info: %s", mySqlInfo))
			panic(err)
		}
		dbInstance = &db{
			mysqlDB,
		}
	})
}

func GetDB() *db {
	return dbInstance
}

func (instance *db) Close() {
	instance.Mysql.Close()
}
