package main

import (
	"e41q/noc_exercises/canvas"
	"math/rand"
	"syscall/js"
)

var frameCount int

func main() {
	canvas := canvas.NewCanvas()
	values := make([]float64, int(canvas.Width()))

	var renderFrame js.Func
	renderFrame = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		canvas.Clear()
		canvas.Stroke("black")
		canvas.NoFill()
		canvas.StrokeWeight(2)

		y := rand.Float64() * canvas.Height()
		copy(values, values[1:])
		values[len(values)-1] = y

		canvas.BeginShape()
		for i := 0; i < int(canvas.Width()); i++ {
			canvas.Vertex(float64(i), values[i])
		}
		canvas.EndShape()

		frameCount++

		js.Global().Call("requestAnimationFrame", renderFrame)
		return nil
	})

	js.Global().Call("requestAnimationFrame", renderFrame)

	select {}
}
