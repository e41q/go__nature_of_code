package main

import (
	"syscall/js"
)

var (
	width  float64
	height float64
	ctx    js.Value
)

func main() {
	doc := js.Global().Get("document")
	canvasEl := doc.Call("getElementById", "mycanvas")
	width = canvasEl.Get("clientWidth").Float()
	height = canvasEl.Get("clientHeight").Float()
	ctx = canvasEl.Call("getContext", "2d")

	done := make(chan struct{}, 0)

	var renderFrame js.Func

	walker := Walker{int(width / 2), int(height / 2)}

	ctx.Call("clearRect", 0, 0, width, height)
	ctx.Set("fillStyle", "white")
	ctx.Call("fillRect", 0, 0, width, height)
	ctx.Set("fillStyle", "black")

	renderFrame = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		walker.Step()
		walker.Show(ctx)

		js.Global().Call("requestAnimationFrame", renderFrame)
		return nil
	})
	defer renderFrame.Release()

	// Start running
	js.Global().Call("requestAnimationFrame", renderFrame)

	<-done
}
