package ui

import (
	"project/pixl/apptype"
	"project/pixl/pxcanvas"
	"project/pixl/swatch"

	"fyne.io/fyne/v2"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
