package uploader

type ImagesCollection struct {
  List []*Image
}

var Images ImagesCollection

func (images *ImagesCollection) Iter () <- chan *Image {
    ch := make(chan *Image);
    size := len(images.List)
    go func () {
        for i := 0; i < size; i++ {
            ch <- images.List[i]
        }
        close(ch)
    } ()
    return ch
}

func SaveImage(image *Image) {
  Images.List = append(Images.List, image)
}


func init() {
  Images := new(ImagesCollection)
  Images.List = make([]*Image, 10)
}
