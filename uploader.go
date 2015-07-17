package main

import (
  "fmt"
  "net/http"
  "io"
  "os"
  "encoding/json"
  "strings"
  "regexp"
  "crypto/md5"
  "time"
)

type Image struct {
    Link     string `json:"link"`
    Filename string `json:"filename"`
    Time     int64
}

func saveImageHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method != "POST" {
    fmt.Fprintf(w, "use POST method for upload!")
    return
  }

  imageFile, imageHeader, err := r.FormFile("image")
  defer imageFile.Close()

  if err != nil {
    fmt.Fprintf(w, "error: %v", err)
    return
  }

  firstImageBytes := make([]byte, 512) // why 512 bytes ? see http://golang.org/pkg/net/http/#DetectContentType
  _, err = imageFile.Read(firstImageBytes)
  if err != nil {
    fmt.Fprintf(w, "error: ", err)
    return
  }

  md5Checksum := md5.Sum(firstImageBytes)

  extensionMatcher := regexp.MustCompile("\\.\\w+$")
  imageName := extensionMatcher.ReplaceAllString(imageHeader.Filename, "")

  filetype := http.DetectContentType(firstImageBytes)
  var extension string
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
      fmt.Fprintf(w, "unknown filetype: %v", filetype)
      return
  }

  uploadFilePath := fmt.Sprintf("./uploads/%x-%v.%v", md5Checksum, imageName, extension)
  uploadedFile, err := os.Create(uploadFilePath)
  if err != nil {
    fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
    return
   }
  defer uploadedFile.Close()

  imageFile.Seek(0, 0) // because we have already read first 512 bytes
  _, err = io.Copy(uploadedFile, imageFile)
  if err != nil {
    fmt.Fprintf(w, "filesave error: %v", err)
    return
  }

  link := strings.Replace(uploadFilePath, "./uploads/", "http://localhost:8080/image/", 1)
  timestamp := time.Now().UnixNano() / int64(time.Millisecond)
  image := Image { link, imageHeader.Filename, timestamp }

  jsonBytes, _ := json.Marshal(image)

  fmt.Fprintf(w, string(jsonBytes))
}

func showImageHandler(w http.ResponseWriter, r *http.Request) {
  filePath := strings.Replace(r.URL.Path, "/image/", "./uploads/", 1)
  http.ServeFile(w, r, filePath)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "ok")
}

func main() {
  var port = ":8080"
  http.HandleFunc("/upload/", saveImageHandler)
  http.HandleFunc("/image/", showImageHandler)
  http.HandleFunc("/ping", statusHandler)
  fmt.Printf("Server running on port: %v...", port)
  http.ListenAndServe(port, nil)
}
