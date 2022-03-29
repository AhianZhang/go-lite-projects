package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
)

const (
	IMG_1_NAME = "img1.png"
	IMG_2_NAME = "img2.png"
)

func main() {
	// read img1
	img1, err := os.Open(IMG_1_NAME)
	if err != nil {
		fmt.Println("read image 1 failed")
		fmt.Println(err)
		return
	}
	_img1, err := png.Decode(img1)
	if err != nil {
		fmt.Println(err)
	}

	defer img1.Close()
	// read img2
	img2, err := os.Open(IMG_2_NAME)
	_img2, err := png.Decode(img2)
	if err != nil {
		fmt.Println(err)
	}

	defer img2.Close()
	// merge
	// {x,y}
	offset := image.Point{200, 300}

	b := _img1.Bounds()
	img3 := image.NewRGBA(b)

	draw.Draw(img3, b, _img1, image.Point{}, draw.Src)
	draw.Draw(img3, _img2.Bounds().Add(offset), _img2, image.Point{}, draw.Over)

	_img3, err := os.Create("result.png")

	if err != nil {
		fmt.Println(err)
	}
	png.Encode(_img3, img3)
	defer _img3.Close()
}
