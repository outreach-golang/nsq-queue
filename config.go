package nsq_queue

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
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
	//redis 客户端
	RedisCli *redis.Client
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

var NsqUrl string 

func setDefaultUrl(){
	url := "127.0.0.1:4150"
	SetNsqUrl(url)
}

func SetNsqUrl(url string){
	NsqUrl = url;
}

func GetNsqUrl() string {
	return NsqUrl;
}

func DefaultProducer() *NsqProducer {
	var (
		p             *NsqProducer
		config        *nsq.Config
		producer      *nsq.Producer
		db            *gorm.DB
		param         string
		clientOptions *options.ClientOptions
		mongoClient   *mongo.Client
		redisClient   *redis.Client
	)
	p = &NsqProducer{}

	config = NewConfig()
	// nsq 客户端
	if len(NsqUrl) == 0{
		setDefaultUrl()
	}
	url := GetNsqUrl()
	producer, _ = nsq.NewProducer(url, config)

	// mysql 客户端
	db, _ = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&timeout=5s&loc=Asia%%2FShanghai", "root", "root", "127.0.0.1:3306", "nsq"))

	// redis客户端
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// momgod 客户端
	param = fmt.Sprintf("mongodb://127.0.0.1:27017")
	clientOptions = options.Client().ApplyURI(param)
	mongoClient, _ = mongo.Connect(context.TODO(), clientOptions)

	p.Producer = producer
	p.MysqlDBS = map[string]*gorm.DB{
		"manage": db,
	}
	p.MongoCli = mongoClient
	p.RedisCli = redisClient

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

func SetRedis(redisclient *redis.Client) Option {
	return func(nsq *NsqProducer) {
		nsq.RedisCli = redisclient
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
