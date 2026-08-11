package main

import (
	"fmt"
	"rolebasedprotection/data"
	"rolebasedprotection/image"
)

func main() {
	regestry := data.GetImageDataRegistry()
	regestry.Register("img1", []byte{1, 2, 3, 4})
	regestry.Register("img2", []byte{2, 3, 4, 5})
	regestry.Register("img3", []byte{3, 4, 5, 6})

	img1 := image.NewAccessCheck("img1", "ADMIN")
	img2 := image.NewAccessCheck("img2", "USER")
	img3 := image.NewAccessCheck("img3", "ADMIN")

	fmt.Println(img1.GetFileName())
	img1.Display()

	fmt.Println(img2.GetFileName())
	img2.Display()

	fmt.Println(img3.GetFileName())
	img3.Display()
}
