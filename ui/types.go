package ui

import (
	"fyne.io/fyne/v2"
	"github.com/E-phraim/pixls/swatch"
	"github.com/E-phraim/pixls/apptype"

)

type AppInit struct {
	PixlsWindow fyne.Window
	State *apptype.State
	Swatches []*swatch.Swatch
}