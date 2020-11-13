package model

// the mongo model
type SubscriptionLog struct {
	Message    string `bson:"message"`
	RegistryId int32  `bson:"registryId"`
	TopicName  string `bson:"topicName"`
	TopicTag   string `bson:"topicTag"`
	Handler    string `bson:"handler"`
	Status     int32  `bson:"status"`
	Remark     string `bson:"remark"`
}
