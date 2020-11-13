package nsq_queue

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/nsqio/go-nsq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type Option func(*NsqProducer)

type NsqProducer struct {
	//消息生产者
	Producer *nsq.Producer
	//mysql 连接对象 maps
	MysqlDBS map[string]*gorm.DB
	//mongo 连接客户端
	MongoCli *mongo.Client
	//log 日志
	Log *zap.Logger
	//	消息体结构
	Message Message
}

type Message struct {
	Msg       string
	TopicName string
	TopicTag  string
	Status    int
	Remark    string
	Handler   string
}

func DefaultProducer() *NsqProducer {
	var (
		p             *NsqProducer
		config        *nsq.Config
		producer      *nsq.Producer
		db            *gorm.DB
		param         string
		clientOptions *options.ClientOptions
		client        *mongo.Client
	)
	p = &NsqProducer{}

	config = NewConfig()
	// nsq 客户端
	producer, _ = nsq.NewProducer("127.0.0.1:4150", config)

	// mysql 客户端
	db, _ = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&timeout=5s&loc=Asia%%2FShanghai", "root", "root", "127.0.0.1:3306", "nsq"))

	// momgod 客户端
	param = fmt.Sprintf("mongodb://127.0.0.1:27017")
	clientOptions = options.Client().ApplyURI(param)
	client, _ = mongo.Connect(context.TODO(), clientOptions)

	p.Producer = producer
	p.MysqlDBS = map[string]*gorm.DB{
		"manage": db,
	}
	p.MongoCli = client

	return p
}

func SetMysql(dbs map[string]*gorm.DB) Option {
	return func(nsq *NsqProducer) {
		nsq.MysqlDBS = dbs
	}
}

func SetMongo(mongoclient *mongo.Client) Option {
	return func(nsq *NsqProducer) {
		nsq.MongoCli = mongoclient
	}
}

func SetLogger(logger *zap.Logger) Option {
	return func(nsq *NsqProducer) {
		nsq.Log = logger
	}
}

func NewConfig() *nsq.Config {
	return nsq.NewConfig()
}
