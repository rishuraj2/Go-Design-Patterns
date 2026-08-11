package image

import (
	"fmt"
	"highresimage/data"
)

type HighResolutionImage struct {
	filename  string
	imageData []byte
}

func NewHighResolutionImage(filename string) *HighResolutionImage {
	imgData, _ := data.GetImageDataRegistry().Fetch(filename)
	return &HighResolutionImage{
		filename:  filename,
		imageData: imgData,
	}
}

func (this *HighResolutionImage) Display() {
	fmt.Println(this.imageData)
}

func (this *HighResolutionImage) GetFileName() string {
	return this.filename
}
