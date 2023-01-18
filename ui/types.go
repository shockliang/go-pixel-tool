package ui

import (
	"fyne.io/fyne/v2"
	"go-pixel-tool/apptype"
	"go-pixel-tool/pxcanvas"
	"go-pixel-tool/swatch"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
