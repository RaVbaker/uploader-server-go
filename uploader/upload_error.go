package uploader

import (
	"encoding/json"
	"time"
)

type UploadError struct {
	Reason  string `json:"error_message"`
	Subject string `json:"for"`
	When    int64  `json:"when"`
}

func (e *UploadError) Error() string {
	return e.Json()
}

func (e *UploadError) Json() string {
	if e.When == 0 {
		e.When = time.Now().Unix()
	}
	jsonBytes, _ := json.Marshal(e)
	return string(jsonBytes)
}
