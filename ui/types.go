package ui

import (
	"fyne.io/fyne/v2"
	"github.com/E-phraim/pixls/apptype"
	"github.com/E-phraim/pixls/pxcanvas"
	"github.com/E-phraim/pixls/swatch"
)

type AppInit struct {
	PixlsCanvas *pxcanvas.PxCanvas
	PixlsWindow fyne.Window
	State *apptype.State
	Swatches []*swatch.Swatch
}