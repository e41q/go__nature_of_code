package main

import (
	"e41q/noc_exercises/p5math"
	"math"
	"syscall/js"
)

var (
	width  float64
	height float64
	ctx    js.Value
)

func main() {
	// Init Canvas stuff
	doc := js.Global().Get("document")
	canvasEl := doc.Call("getElementById", "mycanvas")
	width = canvasEl.Get("clientWidth").Float()
	height = canvasEl.Get("clientHeight").Float()
	ctx = canvasEl.Call("getContext", "2d")

	done := make(chan struct{}, 0)

	// Очищаем холст
	ctx.Call("clearRect", 0, 0, width, height)
	ctx.Set("fillStyle", "white")
	ctx.Call("fillRect", 0, 0, width, height)

	ctx.Set("fillStyle", "rgba(0,0,0,0.01)")

	var renderFrame js.Func
	renderFrame = js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		y := height / 2
		x := p5math.RandomGaussian(400, 100)

		// println("x is ", x)

		ctx.Call("beginPath")
		ctx.Call("arc", x, y, 12, 0, math.Pi*2)
		ctx.Call("fill")

		js.Global().Call("requestAnimationFrame", renderFrame)
		return nil
	})
	defer renderFrame.Release()

	// Start running
	js.Global().Call("requestAnimationFrame", renderFrame)

	<-done
}
