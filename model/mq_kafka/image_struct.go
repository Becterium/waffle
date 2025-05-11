package mq_kafka

type Image struct {
	Id        uint     `json:"id,omitempty"`
	ImageUuid string   `json:"imageUuid,omitempty"`
	ImageName string   `json:"imageName,omitempty"`
	ImageUrl  string   `json:"imageUrl,omitempty"`
	Category  string   `json:"category,omitempty"`
	Purity    string   `json:"purity,omitempty"`
	Uploader  uint     `json:"uploader,omitempty"`
	Size      int64    `json:"size,omitempty"`
	Views     int64    `json:"views,omitempty"`
	Tags      []uint64 `json:"tags,omitempty"`
}
