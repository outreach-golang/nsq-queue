package option

import (
	"github.com/jinzhu/gorm"
	"github.com/outreach-golang/nsq-queue/model"
	"strconv"
)

func AddSub(msg map[string]string, dbs map[string]*gorm.DB) error {
	//根据调用者发来的数据，拼装sub 对象，写入mysql
	var (
		sub       = model.Subscription{}
		status, _ = strconv.ParseInt(msg["status"], 10, 32)
	)

	sub = model.Subscription{
		TopicName: msg["TopicName"],
		TopicTag:  msg["TopicTag"],
		Handler:   msg["Handler"],
		Status:    int32(status),
		Remark:    msg["Remark"],
	}

	db := dbs["mbs_manage"].Table("nsq_mbs." + sub.TableName()).Create(&sub)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func AddRegistry(msg map[string]string, dbs map[string]*gorm.DB) error {
	var (
		sub       = model.Registry{}
		status, _ = strconv.ParseInt(msg["status"], 10, 32)
	)

	sub = model.Registry{
		TopicName: msg["TopicName"],
		TopicTag:  msg["TopicTag"],
		Status:    int32(status),
		Remark:    msg["Remark"],
	}

	db := dbs["mbs_manage"].Table("nsq_mbs." + sub.TableName()).Create(&sub)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
