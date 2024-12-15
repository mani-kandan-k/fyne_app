package widget

import (
	"fyne_app/constants"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func Lable(pText string) *widget.Label {
	return widget.NewLabel(pText)
}

func Btn(pText string, pAction func()) *widget.Button {
	return widget.NewButton(pText, pAction)
}

func IconBtn(pText string, pIcon fyne.Resource, pAction func()) *widget.Button {
	return widget.NewButtonWithIcon(pText, pIcon, pAction)
}

func Input(pPlaceHolder string, pType string) *widget.Entry {

	var lInput *widget.Entry

	if strings.EqualFold(strings.Trim(pType, " "), constants.Password) {
		lInput = widget.NewPasswordEntry()
		lInput.SetPlaceHolder(pPlaceHolder)
		return lInput
	} else {
		lInput = widget.NewEntry()
		lInput.SetPlaceHolder(pPlaceHolder)
		return lInput
	}
}
