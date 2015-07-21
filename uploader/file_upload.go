package uploader

import (
	"errors"
	"regexp"
	mimeMultipart "mime/multipart"
	"crypto/md5"
	"net/http"
	"fmt"
)

var (
	unknownFiletype = errors.New("unknown filetype")
)

func fileExtension(filetype string) (extension string, err error) {
	switch filetype {
	case "image/jpeg", "image/jpg":
		extension = "jpg"
	case "image/gif":
		extension = "gif"
	case "image/png":
		extension = "png"
	case "application/pdf":
		extension = "pdf"
	default:
		err = unknownFiletype
	}
	return
}

func filenameWithoutExtension(filename string) string {
	extensionMatcher := regexp.MustCompile("\\.\\w+$")
	return extensionMatcher.ReplaceAllString(filename, "")
}

func BuildFilePath(file *mimeMultipart.File, StorageDirectory string, filename string) (string, error, string) {
	firstBytes := make([]byte, 512) // why 512 bytes ? see http://golang.org/pkg/net/http/#DetectContentType
	_, err := (*file).Read(firstBytes)
	if err != nil {
		return "", err, "cannot read file"
	}
	defer (*file).Seek(0, 0) // because we have already read first 512 bytes

	md5Checksum := md5.Sum(firstBytes)

	filenamePrefix := filenameWithoutExtension(filename)

	filetype := http.DetectContentType(firstBytes)
	extension, err := fileExtension(filetype)
	if err != nil {
		return "", err, filetype
	}

	uploadFilePath := fmt.Sprintf("%v%x-%v.%v", StorageDirectory, md5Checksum, filenamePrefix, extension)
	return uploadFilePath, nil, ""
}
