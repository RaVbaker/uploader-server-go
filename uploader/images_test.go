package uploader

import (
  _ "fmt"
  "testing"
)

func TestImageSaving(t *testing.T) {
  Images.List = Images.List[:0]

  img1 := NewImage("foo/bar", "img1")
  img2 := NewImage("bar/baz", "img2")
  SaveImage(img1)
  SaveImage(img2)

  if len(Images.List) != 2 {
    t.Error("Didn't manage to save both images")
  }

  if Images.List[1].Filename != "img2" {
    t.Error("Didn't manage to retrieve ")
  }
}


func TestImageSavingFromGoroutine(t *testing.T) {
  Images.List = Images.List[:0]

  img1 := NewImage("foo/bar", "img1")
  img2 := NewImage("bar/baz", "img2")
  done := make(chan bool)

  go func() {
    SaveImage(img1)
    done <- true
  }()
  SaveImage(img2)

  <- done

  if len(Images.List) != 2 {
    t.Error("Didn't manage to save both images. Found only:", len(Images.List))
  }

  if Images.List[1].Filename != "img2" && Images.List[1].Filename != "img1" { // because we don't know the order
    t.Error("Didn't manage to retrieve ", Images.List[1].Filename)
  }
}


func TestIteratingThroughImagesList(t *testing.T) {
  Images.List = Images.List[:0]

  SaveImage(&Image{ "foo/bar", "img1", 1})
  SaveImage(&Image{ "foo/bar", "img2", 2})
  SaveImage(&Image{ "foo/bar", "img2", 4})
  const expectedCounter = 7

  if len(Images.List) != 3 {
    t.Error("Didn't save all 3 images")
  }

  var counter int;
  for image := range Images.Iter() {
    counter = counter + int(image.Time)
  }

  if counter != expectedCounter {
    t.Error("Iterator failed to iterate through all elements. counter:", counter)
  }
}
