package main

import (
	"math/rand"
	"syscall/js"
)

type Walker struct {
	x, y int
}

func (w *Walker) Show(ctx js.Value) {
	ctx.Call("fillRect", w.x, w.y, 1, 1)
}

func (w *Walker) Step(mousePos Mouse) {
	var xstep, ystep int
	xstep = rand.Intn(3) - 1
	ystep = rand.Intn(3) - 1
	if rand.Float32() < 0.5 {
		if rand.Float32() < 0.5 {
			if mousePos.x > float64(w.x) {
				xstep = 1
			} else if mousePos.x < float64(w.x) {
				xstep = -1
			} else {
				xstep = 0
			}
		} else {
			if mousePos.y > float64(w.y) {
				ystep = 1
			} else if mousePos.y < float64(w.y) {
				ystep = -1
			} else {
				ystep = 0
			}
		}
	}
	w.x += xstep
	w.y += ystep
}
