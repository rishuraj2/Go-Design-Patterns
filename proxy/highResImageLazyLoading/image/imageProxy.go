package image

import "fmt"

type ImageProxy struct {
	filename string
	image    *HighResolutionImage
}

func NewImageProxy(filename string) ImageProxy {
	return ImageProxy{
		filename: filename,
	}
}

func (this *ImageProxy) Display() {
	if this.image == nil {
		fmt.Printf("ImageProxy: display() requested for %s. Loading high-resolution image...\n", this.filename)
		this.image = NewHighResolutionImage(this.filename)
	} else {
		fmt.Printf("ImageProxy: Using cached high-resolution image for %s\n", this.filename)
	}

	this.image.Display()
}

func (this *ImageProxy) GetFileName() string {
	return this.filename
}
