package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"github.com/E-phraim/pixls/apptype"
	"github.com/E-phraim/pixls/swatch"
	"github.com/E-phraim/pixls/ui"
)

func main() {
	pixlsApp := app.New()
	pixlsWindow := pixlsApp.NewWindow("pixls")

	state := apptype.State{
		BrushColor: color.NRGBA{255,255,255,255},
		SwatchSelected: 0,
	}

	appInit := ui.AppInit{
		PixlsWindow: pixlsWindow,
		State: &state,
		Swatches: make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	appInit.PixlsWindow.ShowAndRun()
}