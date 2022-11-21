package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/E-phraim/pixls/apptype"
	"github.com/E-phraim/pixls/pxcanvas"
	"github.com/E-phraim/pixls/swatch"
	"github.com/E-phraim/pixls/ui"
)

func main() {
	pixlsApp := app.New()
	pixlsWindow := pixlsApp.NewWindow("Pixls")

	state := apptype.State{
		BrushColor: color.NRGBA{255,255,255,255},
		SwatchSelected: 0,
	}

	pixlsCanvasConfig := apptype.PxCanvasConfig{
		DrawingArea: fyne.NewSize(600,600),
		CanvasOffset: fyne.NewPos(0,0),
		PxRows: 10,
		PxCols: 10,
		PxSize: 30,
	}

	pixlsCanvas := pxcanvas.NewPxCanvas(&state, pixlsCanvasConfig)

	appInit := ui.AppInit{
		PixlsCanvas: pixlsCanvas,
		PixlsWindow: pixlsWindow,
		State: &state,
		Swatches: make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	appInit.PixlsWindow.ShowAndRun()
}