package main

import (
	"image"

	"github.com/AllenDang/giu"
)

var (
	WinWidth, WinHeight int = 400, 350
	isDragging          bool
	dragStartMouse      image.Point
	dragStartPos        image.Point
)

func loop() {
	x, y := giu.Context.Backend().GetWindowPos()
	WinMousePos := giu.GetMousePos()
	WinMousePos.X, WinMousePos.Y = WinMousePos.X-int(x), WinMousePos.Y-int(y)
	if isDragging && !giu.IsMouseDown(giu.MouseButtonLeft) {
		isDragging = false
	}
	if !isDragging && giu.IsMouseClicked(giu.MouseButtonLeft) && WinMousePos.X >= 8 && WinMousePos.Y >= 8 && WinMousePos.X < WinWidth-8 && WinMousePos.Y < 41 {
		dragStartMouse = giu.GetMousePos()
		dragStartPos = image.Pt(int(x), int(y))
		isDragging = true
	}
	if isDragging {
		mousePos := giu.GetMousePos()
		deltaX := mousePos.X - dragStartMouse.X
		deltaY := mousePos.Y - dragStartMouse.Y
		newPos := image.Pt(dragStartPos.X+int(deltaX), dragStartPos.Y+int(deltaY))
		giu.Context.Backend().SetWindowPos(newPos.X, newPos.Y)
	}
	giu.SingleWindow().Layout(
		giu.Child().Size(float32(WinWidth-16), 33).Layout(
			giu.Label("giu snippet"),
		),
		giu.Row(
			giu.Child().Size(50, float32(WinHeight)-20-33).Border(true).Layout(
				giu.Custom(func() {
					giu.Align(giu.AlignCenter).To(
						giu.Button("1").Size(30, 30).OnClick(func() {
							println("Button clicked")
						}),
					).Build()
				}),
			),
			giu.Child().Size(float32(WinWidth-24-50), 100).Layout(
				giu.Button("Debug").OnClick(func() {
				}),
			),
		),
	)
}
func main() {
	wnd := giu.NewMasterWindow("", WinWidth, WinHeight, giu.MasterWindowFlagsNotResizable|giu.MasterWindowFlagsFrameless)
	wnd.Run(loop)
}
