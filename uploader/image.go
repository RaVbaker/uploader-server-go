package uploader

import (
	"encoding/json"
	"time"
)

type Image struct {
	Link     string `json:"link"`
	Filename string `json:"filename"`
	Time     int64
}

func (image *Image) Json() string {
	jsonBytes, _ := json.Marshal(image)

	return string(jsonBytes)
}

func NewImage(link, filename string) *Image {
	timestamp := time.Now().Unix()

	return &Image{link, filename, timestamp}
}
