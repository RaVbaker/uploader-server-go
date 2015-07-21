package uploader

import (
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
