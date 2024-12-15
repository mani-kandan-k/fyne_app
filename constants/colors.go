package constants

import "image/color"

func ColorRed() color.NRGBA {
	return color.NRGBA{R: 0xff, G: 0x33, B: 0x33, A: 0xff}
}

func ColorGreen() color.NRGBA {
	return color.NRGBA{R: 0, G: 180, B: 0, A: 255}
}

func ColorBlue() color.NRGBA {
	return color.NRGBA{R: 0, G: 0, B: 180, A: 255}
}
