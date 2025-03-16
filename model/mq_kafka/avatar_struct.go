package mq_kafka

type Avatar struct {
	Id         uint   `json:"id"`
	UserID     uint   `json:"userID"`
	AvatarName string `json:"avatarName"`
	AvatarUuid string `json:"avatarUuid"`
	AvatarUrl  string `json:"avatarUrl"`
}
