package nsq_queue

import (
	"fmt"
	// "os"
	// "os/signal"
	// "syscall"
	"github.com/nsqio/go-nsq"
)

// NSQ Consumer

// MyHandler 是一个消费者类型

type MyHandler struct {
	Title string
}

type ConsumerInterface interface {
    HandleMsg(msg string) error
}

var Consumer ConsumerInterface;

// HandleMessage 是需要实现的处理消息的方法
func (m *MyHandler) HandleMessage(msg *nsq.Message) (err error) {
	// fmt.Printf("%s recv from %v, msg:%v\n", m.Title, msg.NSQDAddress, string(msg.Body))
	Consumer.HandleMsg(string(msg.Body))
	return
}

// 初始化消费者
func InitConsumer(topic string, channel string, address string, config *nsq.Config, consumer ConsumerInterface) (err error) {

	c, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		fmt.Printf("create consumer failed, err:%v\n", err)
		return
	}
	
	Consumer = consumer
	
	consumerHandler := &MyHandler{
		Title: "handler",
	}
	
	
	c.AddHandler(consumerHandler)

	// if err := c.ConnectToNSQD(address); err != nil { // 直接连NSQD
	if err := c.ConnectToNSQLookupd(address); err != nil { // 通过lookupd查询
		return err
	}
	return nil

}

/*
func main() {
	address := "127.0.0.1:4161"
	err := initConsumer("test", "nsq_to_file", address)
	if err != nil {
		fmt.Printf("init consumer failed, err:%v\n", err)
		return
	}
	c := make(chan os.Signal)        // 定义一个信号的通道
	signal.Notify(c, syscall.SIGINT) // 转发键盘中断信号到c
	<-c                              // 阻塞
}
*/