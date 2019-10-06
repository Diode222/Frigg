package main

import (
	"fmt"
	"github.com/Diode222/Frigg/conf"
	"github.com/Diode222/Frigg/db"
	"github.com/Diode222/Frigg/manager"
	"github.com/Diode222/Frigg/model"
	pb "github.com/Diode222/Frigg/proto_gen"
	"github.com/Diode222/Frigg/serviceServer"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"log"
)

func main() {
	initMysql()
	initService()
	defer db.GetDB().Close()
}

func initMysql() {
	// 必须拥有数据配置文件才能运行
	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", conf.MYSQL_USERNAME, conf.MYSQL_PASSWORD, conf.MYSQL_IP, conf.MYSQL_PORT, conf.MYSQL_DBNAME)
	db.InitDB(sqlInfo)
	dbInstance := db.GetDB()
	dbInstance.Mysql.AutoMigrate(&model.WordFreqItem{})
}

func initService() {
	grpcServer := grpc.NewServer()
	defer grpcServer.GracefulStop()

	pb.RegisterWordFreqListServiceServer(grpcServer, serviceServer.NewWordFreqListServer())
	pb.RegisterChatMessageListServiceServer(grpcServer, serviceServer.NewChatMessageListServer())
	err := manager.GetServiceManger(conf.ETCD_ADDRESS).Register(conf.SERVICE_NAME, conf.LISTEN_IP, conf.SERVICE_IP, conf.SERVICE_PORT, grpcServer, conf.TTL)
	if err != nil {
		log.Printf(fmt.Sprintf("Register service to etcd failed, err: %s", err.Error()))
		panic(err)
	}
}
