package main

import (
	"e41q/noc_exercises/canvas"
	"math/rand"
	"syscall/js"
)

type Walker struct {
	x, y int
}

func (w *Walker) Show(c *canvas.Canvas) {
	c.FillRect(float64(w.x), float64(w.y), 1, 1)
}

func (w *Walker) Step() {
	xstep := rand.Intn(3) - 1
	ystep := rand.Intn(3) - 1
	w.x += xstep
	w.y += ystep
}

func main() {
	canvas := canvas.NewCanvas()

	var renderFrame js.Func

	walker := Walker{int(canvas.Width() / 2), int(canvas.Height() / 2)}

	canvas.Clear()
	canvas.Fill("white")
	canvas.FillRect(0, 0, canvas.Width(), canvas.Height())
	canvas.Fill("black")

	renderFrame = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		walker.Step()
		walker.Show(canvas)

		js.Global().Call("requestAnimationFrame", renderFrame)
		return nil
	})
	defer renderFrame.Release()

	js.Global().Call("requestAnimationFrame", renderFrame)

	select {}
}
