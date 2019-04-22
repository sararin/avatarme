package drawimg

import (
	"fmt"
	"bytes"
	"image"
	"image/color"
	"image/png"
)

func FromHash(u uint64) []byte {
	height, width := 5, 5
	img := image.NewRGBA(image.Rect(0, 0, height, width))

	shape := deduceShape(u)
	fgColor, bgColor := getColors(u)

	var currentColor color.RGBA
	var i int = 0
	for y, x := range shape {
		if x == 0 { 
			currentColor = bgColor 
		} else {
			currentColor = fgColor
		}

		if y%5 == 0 {
			i++
		}

		img.Set(i, y%5,  currentColor)

	}
	
	var buff bytes.Buffer
	png.Encode(&buff, img)
	return buff.Bytes()
}

func deduceShape(u uint64) [25]int {
	basicShape := [25]int{}
	x := 0
	f := u
	for i := uint64(0); i < 64; i++ {
		if (f & (1 << i)) != 0 {
			basicShape[x] = 1
		} else {
			basicShape[x] = 0
		}
		if x >= 14 { x = -1 }
		x++
	}
	
	for i := 15; i < 21; i++ {
		basicShape[i] = basicShape[i-10]
	}
	for i := 20; i < 25; i++ {
		basicShape[i] = basicShape[i-20]
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
