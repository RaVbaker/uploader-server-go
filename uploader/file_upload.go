package uploader

import (
  "errors"
)

var (
  unknownFiletype = errors.New("unknown filetype")
)

func FileExtension(filetype string) (extension string, err error) {
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
