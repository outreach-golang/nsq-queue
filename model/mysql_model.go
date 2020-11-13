package model

// the mysql model
type Subscription struct {
	Id         int32  `gorm:"column:id;" json:"id"`
	RegistryId int32  `gorm:"column:registryId;" json:"registryId"`
	TopicName  string `gorm:"column:topicName;" json:"topicName"`
	TopicTag   string `gorm:"column:topicTag;" json:"topicTag"`
	Handler    string `gorm:"column:handler;" json:"handler"`
	Status     int32  `gorm:"column:status;" json:"status"`
	Remark     string `gorm:"column:remark;" json:"remark"`
}

func (s *Subscription) TableName() string {
	return "subscription"
}

type Registry struct {
	Id        int32  `gorm:"column:id;" json:"id"`
	TopicName string `gorm:"column:topicName;" json:"topicName"`
	TopicTag  string `gorm:"column:topicTag;" json:"topicTag"`
	Status    int32  `gorm:"column:status;" json:"status"`
	Remark    string `gorm:"column:remark;" json:"remark"`
}

func (r *Registry) TableName() string {
	return "registry"
}
