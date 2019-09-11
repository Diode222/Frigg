package main

import (
	"fmt"
	"github.com/Diode222/Frigg/conf"
	"github.com/Diode222/Frigg/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	// 必须拥有数据配置文件才能运行
	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", conf.MYSQL_USERNAME, conf.MYSQL_PASSWORD, conf.MYSQL_IP, conf.MYSQL_PORT, conf.MYSQL_DBNAME)
	friggDB, err := gorm.Open("mysql", sqlInfo)
	if err != nil {
		panic(err.Error())
	}
	defer friggDB.Close()

	friggDB.AutoMigrate(&model.WordItem{})
}
