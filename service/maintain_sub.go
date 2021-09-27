package service

import (
	// "errors"
	// "context"
	// _ "github.com/go-redis/redis"
	
)

const NSQ_CONSUMER = "nsq_consumer"

/*
func CheckTopicExit(topicName string, redisclient *redis.Client) (bool, string) {
	//检查redis中是否存在存在值
	var (
		exist bool
		value string
	)
	exist = redisclient.HExists(context.TODO(),NSQ_CONSUMER, topicName).Val()
	if exist {
		value = redisclient.HGet(context.TODO(),NSQ_CONSUMER, topicName).Val()
		return exist, value
	}
	return exist, ""
}

func AssembleValue(topic, value string, redisclient *redis.Client) string {
	// 组织 需要维护的数据
	var (
		values string
		exist  bool
		v      string
	)
	exist, v = CheckTopicExit(topic, redisclient)
	if exist {
		values = v + "," + value
	} else {
		values = value
	}
	return values
}

func MaintainSub(topic, handurl string, redisclient *redis.Client) error {
	var (
		value  string
		number int64
	)
	value = AssembleValue(topic, handurl, redisclient)
	number = redisclient.HSet(context.TODO(),NSQ_CONSUMER, topic, value).Val()
	if number == -1 {
		return errors.New("maintain the handurl to the redis fail")
	}
	return nil
}
*/
