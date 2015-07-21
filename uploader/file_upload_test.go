package uploader

import (
  "testing"
)

func TestFileExtensionDetection(t *testing.T) {
  var filetypeToExtension map[string]string
  filetypeToExtension = map[string]string {
    "image/jpeg": "jpg",
    "image/jpg": "jpg",
    "image/png": "png",
    "image/gif": "gif",
  }

  for filetype, expectedExtension := range filetypeToExtension {
    foundExtension, _ := fileExtension(filetype)
    if foundExtension != expectedExtension {
      t.Errorf("Didn't detect proper extension from %s. Found: %v", filetype, foundExtension)
    }
  }
}

func TestFileExtensionDetectionFailed(t *testing.T) {
  _, err := fileExtension("unknown/type")
  if err == nil {
    t.Fatal("We should return error when extension not detected!")
  }
}
