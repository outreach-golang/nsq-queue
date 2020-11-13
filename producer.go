package nsq_queue

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/outreach-golang/nsq-queue/option"
	"go.mongodb.org/mongo-driver/mongo"
)

var Producer *NsqProducer

func NewNsqProducer(ops ...Option) (*NsqProducer, error) {

	producer := DefaultProducer()

	for _, op := range ops {
		op(producer)
	}

	//if config.ServerName == "" {
	//	return nil, errors.New("ServerName参数必填！")
	//}
	//
	//Logger, err := logger(config)
	//if err != nil {
	//	return nil, err
	//}
	//
	Producer = producer

	return Producer, nil
}

func (p *NsqProducer) BindHander(msg map[string]string) error {
	//	绑定需要回调的url 等信息 入mysql
	fmt.Println("bind handler info to mysql ")
	err := option.AddSub(msg, p.MysqlDBS)
	return err
}

func (p *NsqProducer) AddLogs(msg map[string]string) error {
	//	将消息信息写入mongo
	fmt.Println("add the message to mongo")
	err := option.AddSubLog(msg, p.MongoCli)
	return err
}

func (p *NsqProducer) ProducerMessage(msg map[string]string, dbs map[string]*gorm.DB, client *mongo.Client) error {
	var (
		message = msg["message"]
		topic   = msg["TopicName"]
	)

	if topic == "" || message == "" {
		return errors.New("缺少必要的参数")
	}

	//mysql记录 回调等信息
	err := p.BindHander(msg)
	if err != nil {
		fmt.Println("bind handler info to mysql fail")
		return err
	}

	//推送信息
	err = p.Producer.Publish(topic, []byte(message))
	if err != nil {
		fmt.Println("message producer fail")
		return err
	}

	//将消息日志写入mongo
	p.AddLogs(msg)

	fmt.Printf("%s producer message success!\n", message)
	return nil
}
