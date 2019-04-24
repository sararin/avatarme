package drawimg

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
)

func FromHash(u uint64) ([]byte, error) {
	height, width := 5, 5
	img := image.NewRGBA(image.Rect(0, 0, height, width))

	shape := deduceShape(u)
	fgColor, bgColor := getColors(u)
	var currentColor color.RGBA
	for y, x := range shape {
		for f, g := range x {
			currentColor = fgColor
			if g == 0 {
				currentColor = bgColor
			}
			img.Set(y, f, currentColor)
		}
	}
	
	var buff bytes.Buffer
	png.Encode(&buff, img)
	return buff.Bytes(), nil
}

func deduceShape(u uint64) [5][5]int {
	for i := uint64(0); i < 64; i++ {
		basicShape[x][y] = 0
		if (u & (1 << i)) != 0 {
			basicShape[x][y] = 1
		}
		x++
		if x > 3 { y++; x = 0 }
		if y > 3 { y = 0 }
	}
	for y, x := range basicShape[0] {
		basicShape[4][y] = x
	}
	for y, x := range basicShape[1] {
		basicShape[3][y] = x
	}
	return basicShape
}

// in theory:
// uint64 => 0000 1111 0000 1111 0000 1111 0000 1111 0000 1111 0000 1111 0000 1111 0000 1111
// fg => takes first 24bits
// bg => takes last 24bits
func getColors(u uint64) (color.RGBA, color.RGBA) {
	fg := 0xFFFFFF & u
	bg := 0xFFFFFF000000 & u >> 24
	redFg, greenFg, blueFg := deduceRGB(fg)
	redBg, greenBg, blueBg := deduceRGB(bg)
	return color.RGBA{redFg, greenFg, blueFg, 0xFF}, color.RGBA{redBg, greenBg, blueBg, 0xFF}
}

func deduceRGB(u uint64) (uint8, uint8, uint8) {
	return uint8(u & 255), uint8((u >> 8) & 255), uint8((u >> 16) & 255)
}