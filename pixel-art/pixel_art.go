package main

import (
	"NikoJT/pixel-art/swatch"
	"NikoJT/pixel-art/ui"	
	"NikoJT/pixel-art/apptype"
	"image/color"
	"fyne.io/fyne/v2/app"
)

func main() {
	pixelApp := app.New()
	pixelWindow := pixelApp.NewWindow("pixel")

	state := apptype.State {
		BrushColor: color.NRGBA{255, 255, 255, 255},
		SwatchSelected: 0,
	}

	appInit := ui.AppInit{
		PixelWindow: pixelWindow,
		State: &state,
		Swatches: make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	appInit.PixelWindow.ShowAndRun()
}
