package apptype

import (
	"image/color"
	"fyne.io/fyne/v2"
)

type BrushType = int

type PxCanvasConfig struct {
	DrawingArea fyne.Size
	CanvasOffset fyne.Position
	PxRows, PxCols int
	PxSze int
}

type State struct {
	BrushColor color.Color
	BurshType int
	SwatchSelected int
	FilePath string
}

func (state *State) SetFilePath(path string) {
	state.FilePath = path
}

type Brushable interface {
	SetColor(c color.Color, x, y int)
	MouseToCanvasXY(ev *desktop.MouseEvent) (*int, *int)
}
