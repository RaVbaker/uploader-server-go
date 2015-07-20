package uploader

import (
	"fmt"
	"net/http"
)

func PrintJsonErrorString(w *http.ResponseWriter, reason, subject string) {
	err := &UploadError{reason, subject, 0}
	fmt.Fprint(*w, err.Json())
}

func PrintJsonError(w *http.ResponseWriter, reason string, subject error) {
	PrintJsonErrorString(w, reason, subject.Error())
}

func PrintJsonErrorDetails(w *http.ResponseWriter, reason error, subject string) {
	PrintJsonErrorString(w, reason.Error(), subject)
}
