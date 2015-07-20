package uploader

import "encoding/json"

type Image struct {
	Link     string `json:"link"`
	Filename string `json:"filename"`
	Time     int64
}

func (image *Image) Json() string {
	jsonBytes, _ := json.Marshal(image)

	return string(jsonBytes)
}
