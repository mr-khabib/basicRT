package render

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

const width = 1024
const w_step = 255 / float32(width)
const height = 768
const h_step = 255 / float32(height)

func Render() {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			x_f := float32(x)
			y_f := float32(y)
			r := uint8(h_step * x_f)
			g := uint8(w_step * y_f)
			b := uint8(((r/2 + g/2) / 2))
			fmt.Println(r, g, b)
			col := color.RGBA{r, g, b, 0xff}
			img.Set(x, y, col)
		}
	}

	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
