package main

import (
	"fmt"
	"time"
	"os"
	"os/signal"
	"syscall"
	"github.com/nsqio/go-nsq"
	nsq_queue "github.com/outreach-golang/nsq-queue"
)

type MsgHandler struct {
	Title string
}

// HandleMsg 是需要实现的处理消息的方法
func (m MsgHandler) HandleMsg(msg string) (err error) {
	fmt.Printf("recv msg:%v\n", msg)
	return
}

func main() {
	var consumer MsgHandler
	consumer.Title = "title_test"
	address := "127.0.0.1:4161"
	config := nsq.NewConfig()
	config.LookupdPollInterval = 15 * time.Second
	
	err := nsq_queue.InitConsumer("test", "nsq_to_file", address, config, consumer)
	if err != nil {
		fmt.Printf("init consumer failed, err:%v\n", err)
		return
	}
	c := make(chan os.Signal)        
	signal.Notify(c, syscall.SIGINT) 
	<-c                              
}