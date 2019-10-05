package main

import (
	"fmt"
	"github.com/Diode222/Frigg/conf"
	"github.com/Diode222/Frigg/manager"
	"github.com/Diode222/Frigg/model"
	pb "github.com/Diode222/Frigg/proto_gen"
	chatMessageListServer "github.com/Diode222/Frigg/serviceServer"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"log"
)

func main() {
	initMysql()
	initService()
}

func initMysql() {
	// 必须拥有数据配置文件才能运行
	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", conf.MYSQL_USERNAME, conf.MYSQL_PASSWORD, conf.MYSQL_IP, conf.MYSQL_PORT, conf.MYSQL_DBNAME)
	friggDB, err := gorm.Open("mysql", sqlInfo)
	if err != nil {
		panic(err.Error())
	}
	defer friggDB.Close()

	friggDB.AutoMigrate(&model.WordItem{})
}

func initService() {
	grpcServer := grpc.NewServer()
	defer grpcServer.GracefulStop()

	pb.RegisterWordFreqListServiceServer(grpcServer, chatMessageListServer.NewWordFreqListServer())
	pb.RegisterChatMessageListServiceServer(grpcServer, chatMessageListServer.NewChatMessageListServer())
	err := manager.GetServiceManger(conf.ETCD_ADDRESS).Register(conf.SERVICE_NAME, conf.SERVICE_IP, conf.SERVICE_PORT, grpcServer, conf.TTL)
	if err != nil {
		log.Printf(fmt.Sprintf("Register service to etcd failed, err: %s", err.Error()))
		panic(err)
	}
}
