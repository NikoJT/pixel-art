package ui
import (
	"fyne.io/fyne/v2"
	"NikoJT/pixel-art/apptype"
	"NikoJT/pixel-art/pxcanvas"
	"NikoJT/pixel-art/swatch"
)

type AppInit struct {
	PixelCanvas *pxcanvas.PxCanvas
	PixelWindow fyne.Window
	State *apptype.State
	Swatches []*swatch.Swatch
}
