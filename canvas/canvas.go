package canvas

import (
	"fyne_app/constants"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func setContentToText(pCanvas fyne.Canvas) {
	blue := constants.ColorBlue()
	text := canvas.NewText("Text", blue)
	text.TextStyle.Bold = true
	pCanvas.SetContent(text)
}

func setContentToCircle(pCanvas fyne.Canvas) {
	red := constants.ColorRed()
	circle := canvas.NewCircle(color.White)
	circle.StrokeWidth = 4
	circle.StrokeColor = red
	pCanvas.SetContent(circle)
}
