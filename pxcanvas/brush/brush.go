package brush

import (
	"fyne.io/fyne/v2/driver/desktop"
	"go-pixel-tool/apptype"
)

const (
	Pixel = iota
)

func TryPaintPixel(appState *apptype.State, canvas apptype.Brushable, ev *desktop.MouseEvent) bool {
	x, y := canvas.MouseToCanvasXY(ev)
	if x != nil && y != nil && ev.Button == desktop.MouseButtonPrimary {
		canvas.SetColor(appState.BrushColor, *x, *y)
		return true
	}
	return false
}

func TryBrush(appState *apptype.State, canvas apptype.Brushable, ev *desktop.MouseEvent) bool {
	if appState.BrushType == Pixel {
		return TryPaintPixel(appState, canvas, ev)
	}

	return false
}
