package mq_kafka

type Image struct {
	ImageUuid string   `json:"imageUuid"`
	ImageName string   `json:"imageName"`
	ImageUrl  string   `json:"imageUrl"`
	Category  string   `json:"category"`
	Purity    string   `json:"purity"`
	Uploader  int64    `json:"uploader"`
	Size      int64    `json:"size"`
	Views     int64    `json:"views"`
	Tags      []uint64 `json:"tagsx"`
}
