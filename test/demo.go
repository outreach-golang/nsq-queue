package main

import (
	"fmt"
	"context"
	//"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"github.com/jinzhu/gorm"
    nsq_queue "github.com/outreach-golang/nsq-queue"
	"github.com/outreach-golang/logger"
)
// 设置项目的mongo, redis配置

var producer *nsq_queue.NsqProducer

func InitProducer() {
	
	var (
		param            string
		clientOptions    *options.ClientOptions
		mongoClient      *mongo.Client
		// db               *gorm.DB
	)
	

	// mongo 客户端
	param = fmt.Sprintf("mongodb://127.0.0.1:27017")
	clientOptions = options.Client().ApplyURI(param)
	mongoClient, _ = mongo.Connect(context.TODO(), clientOptions)
	mongoOption := nsq_queue.SetMongo(mongoClient)

	/*
	// redis 客户端
	redisOption :=  nsq_queue.SetRedis(redis.NewClient(&redis.Options{
		Addr:     "127.0.01:6379",
		Password: "",
		DB:       0,
	}))
	*/
	/*
	// mysql 客户端
	db, _ = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&timeout=5s&loc=Asia%%2FShanghai", "root", "root", "127.0.0.1:3306", "nsq"))
	dbs := map[string]*gorm.DB{
		"manage": db,
	}
	mysqlOption := nsq_queue.SetMysql(dbs)
	*/
	
	nsq_queue.SetNsqUrl("127.0.0.1:4150")
	
    producer, _  = nsq_queue.NewNsqProducer(
		mongoOption,
		//redisOption,
		//mysqlOption,
    )
	
	if gLogger, err := logger.NewLogger(
		logger.ServerName("loggerServer"),
	); err != nil {
		fmt.Println(err.Error())
	} else {
		nsq_queue.SetGLogger(gLogger)
	}
}

func main() {
    //初始化

    InitProducer()
    //组织消息，消息体结构 为map[string]string 结构，主要包含字段 TopicName，Handler，message
    msg := map[string]string{
        "TopicName": "test",
        "Handler":   "127.0.0.1:9000",
        "message":   "bodys 1111",
    }
    //发送
    err := producer.ProducerMessage(msg)
    fmt.Println(err)
	//ch :=  make(chan struct{})
	//<- ch
}