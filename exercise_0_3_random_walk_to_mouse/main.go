package main

import (
	"syscall/js"
)

type Mouse struct {
	x, y float64
}

var (
	width    float64
	height   float64
	ctx      js.Value
	mousePos Mouse
)

func main() {
	// Init Canvas stuff
	doc := js.Global().Get("document")
	canvasEl := doc.Call("getElementById", "mycanvas")
	width = canvasEl.Get("clientWidth").Float()
	height = canvasEl.Get("clientHeight").Float()
	ctx = canvasEl.Call("getContext", "2d")

	done := make(chan struct{}, 0)

	mouseMoveEvt := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := args[0]
		mousePos.x = e.Get("clientX").Float()
		mousePos.y = e.Get("clientY").Float()
		return nil
	})
	defer mouseMoveEvt.Release()
	js.Global().Get("document").Call("addEventListener", "mousemove", mouseMoveEvt)

	var renderFrame js.Func

	walker := Walker{int(width / 2), int(height / 2)}

	ctx.Call("clearRect", 0, 0, width, height)
	ctx.Set("fillStyle", "white")
	ctx.Call("fillRect", 0, 0, width, height)
	ctx.Set("fillStyle", "black")

	renderFrame = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		walker.Step(mousePos)
		walker.Show(ctx)

		js.Global().Call("requestAnimationFrame", renderFrame)
		return nil
	})
	defer renderFrame.Release()

	// Start running
	js.Global().Call("requestAnimationFrame", renderFrame)

	<-done
}
