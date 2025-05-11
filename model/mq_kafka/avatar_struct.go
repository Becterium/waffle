package mq_kafka

type Avatar struct {
	Id         uint   `json:"id,omitempty"`
	UserID     uint   `json:"userID,omitempty"`
	AvatarName string `json:"avatarName,omitempty"`
	AvatarUuid string `json:"avatarUuid,omitempty"`
	AvatarUrl  string `json:"avatarUrl,omitempty"`
}
