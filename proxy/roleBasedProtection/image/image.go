package image

import (
	"fmt"
	"rolebasedprotection/data"
)

type Image struct {
	filename string
	data     []byte
}

func NewImage(filename string) *Image {
	imgData, _ := data.GetImageDataRegistry().Fetch(filename)

	return &Image{
		filename: filename,
		data:     imgData,
	}
}

func (this Image) GetFileName() string {
	return this.filename
}

func (this Image) Display() {
	fmt.Println(this.data)
}
