package main

import (
	"math/rand"
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

	var (
		renderFrame  js.Func
		randomCounts [40]int
	)

	total := len(randomCounts)

	ctx.Call("clearRect", 0, 0, width, height)
	ctx.Set("fillStyle", "white")
	ctx.Call("fillRect", 0, 0, width, height)

	renderFrame = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		index := rand.Intn(total)
		randomCounts[index]++

		ctx.Set("fillStyle", "grey")
		w := width / float64(total)
		println("w is ", w)
		for i, v := range randomCounts {
			if v > 0 {
				ctx.Call("fillRect", float64(i)*w, height-float64(v), w-1, v)
			}
		}

		js.Global().Call("requestAnimationFrame", renderFrame)
		return nil
	})
	defer renderFrame.Release()

	// Start running
	js.Global().Call("requestAnimationFrame", renderFrame)

	<-done
}
