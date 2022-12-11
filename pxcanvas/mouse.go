package pxcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/E-phraim/pixls/pxcanvas/brush"
)

func (pxCanvas *PxCanvas) Scrolled(v *fyne.ScrollEvent){
	pxCanvas.scale(int(v.Scrolled.DY))
	pxCanvas.Refresh()
}

func (pxCanvas *PxCanvas) MouseMoved(v *desktop.MouseEvent){
	if x, y := pxCanvas.MouseToCanvasXY(v); x != nil && y != nil{
		brush.TryBrush(pxCanvas.appState, pxCanvas, v) 
		cursor := brush.Cursor(pxCanvas.PxCanvasConfig, pxCanvas.appState.BrushType, v, *x, *y)
		pxCanvas.renderer.SetCursor(cursor)
	}else{
		pxCanvas.renderer.SetCursor(make([]fyne.CanvasObject, 0))
	}
	pxCanvas.TryPan(pxCanvas.mouseState.previousCoord, v)
	pxCanvas.Refresh()
	pxCanvas.mouseState.previousCoord = &v.PointEvent
}

func (pxCanvas *PxCanvas) MouseIn(ev *desktop.MouseEvent){

}

func (pxCanvas *PxCanvas) MouseOut(){
	
}

func (pxCanvas *PxCanvas) MouseDown(v *desktop.MouseEvent){
	brush.TryBrush(pxCanvas.appState, pxCanvas, v)
}

func (PxCanvas *PxCanvas) MouseUp(v *desktop.MouseEvent){

}