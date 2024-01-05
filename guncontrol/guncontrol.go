package guncontrol

import (
	"image"
	"image/color"
)

type Recoil struct {
	X, Y int
}

type KeyPressCounter struct {
	Count int
	Reset chan struct{}
}

func ColorsMatch(c1, c2 color.Color, tolerance uint32) bool {
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()

	return absDiff(r1, r2) <= tolerance &&
		absDiff(g1, g2) <= tolerance &&
		absDiff(b1, b2) <= tolerance &&
		absDiff(a1, a2) <= tolerance
}

func absDiff(a, b uint32) uint32 {
	if a > b {
		return a - b
	}
	return b - a
}

func ProcessImage(img image.Image, targetColor color.Color, tolerance uint32) *Recoil {
	bounds := img.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixelColor := img.At(x, y)
			if ColorsMatch(pixelColor, targetColor, tolerance) {
				return &Recoil{x, y}
			}
		}
	}

	return nil
}
