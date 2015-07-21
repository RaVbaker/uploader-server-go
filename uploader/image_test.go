package uploader

import (
	"fmt"
	"testing"
)

func TestJsonResponse(t *testing.T) {
	image := Image{"http://localhost:8080/image/123.png", "sample.png", 42}
	var expectedJsonResponse string = `{"link":"http://localhost:8080/image/123.png","filename":"sample.png","Time":42}`
	jsonResponse := image.Json()
	if jsonResponse != expectedJsonResponse {
		t.Errorf("Json response doesn't match requred pattern: '%v' expected: '%v'", jsonResponse, expectedJsonResponse)
	}
}

func TestNewImageCreation(t *testing.T) {
	link, filename := "url", "sample.jpg"
	image := NewImage(link, filename)
	if fmt.Sprintf("%T", image) != "*uploader.Image" {
		t.Fatal("It should return an image")
	}

	if image.Link != link {
		t.Errorf("image.Link(%v) is not equal to expected value: %v", image.Link, link)
	}

	if image.Filename != filename {
		t.Errorf("image.Filename(%v) is not equal to expected value: %v", image.Filename, filename)
	}
}
