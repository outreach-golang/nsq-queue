package nsq_queue

import (
	"errors"
	"fmt"
	"encoding/json"
	"go.uber.org/zap"
	"github.com/outreach-golang/nsq-queue/option"
	_ "github.com/outreach-golang/nsq-queue/service"
	_ "github.com/outreach-golang/logger"
)

var Producer *NsqProducer


var gLogger *zap.Logger

func SetGLogger(log *zap.Logger){
	gLogger = log
}

func GetGLogger() *zap.Logger{
	return gLogger
}


func NewNsqProducer(ops ...Option) (*NsqProducer, error) {

	producer := DefaultProducer()

	for _, op := range ops {
		op(producer)
	}

	Producer = producer

	return Producer, nil
}

/*
func (p *NsqProducer) MaintainSub(msg map[string]string) error {
	//维护 在redis 中的topic 和回调url对应关系
	var (
		topic   = msg["TopicName"]
		handurl = msg["Handler"]
	)
	err := service.MaintainSub(topic, handurl, p.RedisCli)
	return err
}
*/

func (p *NsqProducer) BindHander(msg map[string]string) error {
	//	绑定需要回调的url 等信息 入mongo
	fmt.Println("bind handler info to mongo ")
	err := option.AddSub(msg, p.MongoCli)
	return err
}

func (p *NsqProducer) AddLogs(msg map[string]string) error {
	//	将消息信息写入mongo
	fmt.Println("add the message log  to mongo")
	err := option.AddSubLog(msg, p.MongoCli)
	return err
}

func (p *NsqProducer) ProducerMessage(msg map[string]string) error {
	var (
		message = msg["message"]
		topic   = msg["TopicName"]
	)

	if topic == "" || message == "" {
		return errors.New("缺少必要的参数")
	}

	//mongo记录 回调等信息
	err := p.BindHander(msg)
	if err != nil {
		fmt.Println("bind handler info to mongo fail")
		return err
	}

	//在 redis维护 topic 对应的handlerurl, 做缓存使用
	/*
	err = p.MaintainSub(msg)
	if err != nil {
		fmt.Println("maintain handurl in redis fail")
		return err
	}
	*/

	//推送信息
	err = p.Producer.Publish(topic, []byte(message))
	if err != nil {
		fmt.Println("message producer fail")
		return err
	}

	//将消息日志写入mongo
	err = p.AddLogs(msg)
	if err != nil {
		fmt.Println("add the logs err:%s",err)
	}
	
	log := GetGLogger()
	if log != nil {
		jsonMessage , err := json.Marshal(msg)
		if err == nil {
			log.Info("nsq message:" + string(jsonMessage))
		}
	}
	
	fmt.Printf("%s producer message success!\n", message)
	return nil
}
