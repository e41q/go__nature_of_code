package main

import (
	"e41q/noc_exercises/canvas"
	"e41q/noc_exercises/p5math"
	"syscall/js"
)

func main() {
	canvas := canvas.NewCanvas()
	var t float64 = 0.0

	var renderFrame js.Func
	renderFrame = js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		canvas.Clear()

		xoff := t
		canvas.Stroke("black")
		canvas.NoFill()
		canvas.StrokeWeight(2)
		canvas.BeginShape()

		for i := 0; i < int(canvas.Width()); i++ {
			y := p5math.Noise(xoff) * canvas.Height()
			xoff += 0.01
			canvas.Vertex(float64(i), y)
		}

		canvas.EndShape()
		t += 0.01

		js.Global().Call("requestAnimationFrame", renderFrame)
		return nil
	})

	js.Global().Call("requestAnimationFrame", renderFrame)

	select {}
}
