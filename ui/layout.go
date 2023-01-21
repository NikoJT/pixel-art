package ui

import "fyne.io/fyne/v2/container"
import "NikoJT/pixel-art/ui"

func Setup(app *AppInit) {
	SetupMenus(app)
	swatchesContainer := BuildSwatches(app)
	colorPicker := SetupColorPicker(app)

	appLayout := container.NewBorder(nil, swatchesContainer, nil, colorPicker, app.PixelCanvas)

	app.PixelWindow.SetContent(appLayout)
}

