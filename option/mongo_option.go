package option

import (
	"context"
	"fmt"
	"github.com/outreach-golang/nsq-queue/model"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
)

func AddSubLog(msg map[string]string, client *mongo.Client) error {
	//将发送到nsq 消息日志 存入mongo
	var (
		sublog     *model.SubscriptionLog
		database   *mongo.Database
		collection *mongo.Collection
		status     int64
	)
	status, _ = strconv.ParseInt(msg["status"], 10, 32)

	sublog = &model.SubscriptionLog{
		Message:   msg["message"],
		TopicName: msg["topicName"],
		TopicTag:  msg["topicTag"],
		Handler:   msg["Handler"],
		Status:    int32(status),
		Remark:    msg["remark"],
	}

	// 选择db
	database = client.Database("logs")

	// 选择表my_collection
	collection = database.Collection("subscription")

	if _, err := collection.InsertOne(context.TODO(), sublog); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
