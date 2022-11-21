package pxcanvas

import "fyne.io/fyne/v2"

func (pxCanvas *PxCanvas) Pan(prerviousCoord, currentCoord fyne.PointEvent){
	xDiff := currentCoord.Position.X - prerviousCoord.Position.X
	yDiff := currentCoord.Position.Y - prerviousCoord.Position.Y	

	pxCanvas.CanvasOffset.X += xDiff
	pxCanvas.CanvasOffset.Y += yDiff
	pxCanvas.Refresh()
}

func (pxCanvas *PxCanvas) scale(direction int){
	switch{
	case direction > 0:
		pxCanvas.PxSize += 1
	case direction < 0:
		if pxCanvas.PxSize > 2 {
			
		}
	}
}