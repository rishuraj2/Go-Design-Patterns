package main

import (
	"fmt"
	"highresimage/data"
	"highresimage/image"
)

func main() {
	regestry := data.GetImageDataRegistry()
	regestry.Register("img1", []byte{1, 2, 3, 4})
	regestry.Register("img2", []byte{2, 3, 4, 5})
	regestry.Register("img3", []byte{3, 4, 5, 6})

	img1 := image.NewImageProxy("img1")
	img2 := image.NewImageProxy("img2")
	img3 := image.NewImageProxy("img3")

	fmt.Println(img1.GetFileName())
	img1.Display()

	fmt.Println(img2.GetFileName())

	fmt.Println(img3.GetFileName())
	img3.Display()

	img1.Display()

}
