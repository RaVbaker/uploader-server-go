package main

import (
	_ "crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
	_ "regexp"
	"strings"
	"time"

	"github.com/ravbaker/uploader-server-go/uploader"
)

const (
	DefaultPort = ":8080"
	Protocol    = "http://"

	StorageDirectory = "./uploads/"

	ImagePath = "/image/"
)

func saveImageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		uploader.PrintJsonErrorString(&w, "use POST method for upload!", r.Method)
		return
	}

	imageFile, imageHeader, err := r.FormFile("image")
	if err != nil {
		uploader.PrintJsonError(&w, "Cannot fetch file from form field(image)", err)
		return
	}
	defer imageFile.Close()

	imageFileName := imageHeader.Filename
	uploadFilePath, err, errorSubject := uploader.BuildFilePath(&imageFile, StorageDirectory, imageFileName)
	if err != nil {
		uploader.PrintJsonErrorDetails(&w, err, errorSubject)
		return
	}

	uploadedFile, err := os.Create(uploadFilePath)
	if err != nil {
		uploader.PrintJsonErrorDetails(&w, err, uploadFilePath)
		return
	}
	defer uploadedFile.Close()

	_, err = io.Copy(uploadedFile, imageFile)
	if err != nil {
		uploader.PrintJsonErrorDetails(&w, err, uploadFilePath)
		return
	}

	imageResourceUrl := fmt.Sprint(Protocol, r.Host, ImagePath)
	link := strings.Replace(uploadFilePath, StorageDirectory, imageResourceUrl, 1)
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	image := uploader.Image{link, imageFileName, timestamp}

	fmt.Fprint(w, (&image).Json())
}

func showImageHandler(w http.ResponseWriter, r *http.Request) {
	filePath := strings.Replace(r.URL.Path, ImagePath, StorageDirectory, 1)
	http.ServeFile(w, r, filePath)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func main() {
	http.HandleFunc("/upload/", saveImageHandler)
	http.HandleFunc(ImagePath, showImageHandler)
	http.HandleFunc("/ping", statusHandler)

	var port = DefaultPort
	if envPort := os.Getenv("APP_PORT"); envPort != "" {
		port = fmt.Sprint(":", envPort)
	}

	fmt.Printf("Server running on port: %v...", port)
	http.ListenAndServe(port, nil)
}
