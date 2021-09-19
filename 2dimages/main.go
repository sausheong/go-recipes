package main

import (
	"image"
	"image/png"
	_ "image/png"
	"log"
	"math/rand"
	"os"
)

func main() {
	rect := image.Rect(0, 0, 100, 100)
	img := createRandomImage(rect)
	save("random.png", img)

}

// load the image
func load(filePath string) *image.NRGBA {
	imgFile, err := os.Open(filePath)
	defer imgFile.Close()
	if err != nil {
		log.Println("Cannot read file:", err)
	}

	img, _, err := image.Decode(imgFile)
	if err != nil {
		log.Println("Cannot decode file:", err)
	}
	return img.(*image.NRGBA)
}

// save the image
func save(filePath string, img *image.NRGBA) {
	imgFile, err := os.Create(filePath)
	defer imgFile.Close()
	if err != nil {
		log.Println("Cannot create file:", err)
	}
	png.Encode(imgFile, img.SubImage(img.Rect))
}

// create a random image
func createRandomImage(rect image.Rectangle) (created *image.NRGBA) {
	pix := make([]uint8, rect.Dx()*rect.Dy()*4)
	rand.Read(pix)
	created = &image.NRGBA{
		Pix:    pix,
		Stride: rect.Dx() * 4,
		Rect:   rect,
	}
	return
}
