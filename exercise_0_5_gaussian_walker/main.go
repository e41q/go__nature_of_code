package main

import (
	"e41q/noc_exercises/canvas"
	"e41q/noc_exercises/p5math"
	"math"
	"syscall/js"
)

type Walker struct {
	x, y int
}

func (w *Walker) Show(c *canvas.Canvas) {
    cw, ch := int(c.Width()), int(c.Height())
    if w.x >= cw {
        w.x -= cw
    } else if w.x < 0 {
        w.x += cw
    }
    if w.y >= ch {
        w.y -= ch
    } else if w.y < 0 {
        w.y += ch
    }
	c.FillRect(float64(w.x), float64(w.y), 1, 1)
}

func (w *Walker) Step() {
	xstep := int(math.Round(p5math.RandomGaussian(0, 3)))
	ystep := int(math.Round(p5math.RandomGaussian(0, 3)))
	w.x += xstep
	w.y += ystep
}

func main() {
	canvas := canvas.NewCanvas()
	canvas.Fill("black")

	walker := Walker{int(canvas.Width() / 2), int(canvas.Height() / 2)}

	var renderFrame js.Func
	renderFrame = js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		walker.Step()
		walker.Show(canvas)

		js.Global().Call("requestAnimationFrame", renderFrame)
		return nil
	})

	// Start running
	js.Global().Call("requestAnimationFrame", renderFrame)

	select {}
}
