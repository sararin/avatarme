package avatarme

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
)

// FromHash creates an byte array from any hash.
// By default image is 5x5 to ease the drawing process.
func FromHash(u uint64) []byte {
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
	return buff.Bytes()
}

// deduceShape returns a [5][5]int array made out of only 1s and 0s.
// 1 signifies that the foreground color will be used in the future, while 0 means a background color.
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

// getColors generates colors via using first 24bits as foreground color and last 24bits as background color.
// It returns two RGBA objects.
func getColors(u uint64) (color.RGBA, color.RGBA) {
	fg := 0xFFFFFF & u
	bg := 0xFFFFFF000000 & u >> 24
	redFg, greenFg, blueFg := deduceRGB(fg)
	redBg, greenBg, blueBg := deduceRGB(bg)
	return color.RGBA{redFg, greenFg, blueFg, 0xFF}, color.RGBA{redBg, greenBg, blueBg, 0xFF}
}

// deduceRGB splits 24bytes and returns 8 byte unsigned ints.
func deduceRGB(u uint64) (uint8, uint8, uint8) {
	return uint8(u & 255), uint8((u >> 8) & 255), uint8((u >> 16) & 255)
}