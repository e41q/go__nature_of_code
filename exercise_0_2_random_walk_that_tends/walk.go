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

func (w *Walker) Step() {
	prob := []int{-1, 0, 1, 1}
	xstep := prob[rand.Intn(len(prob))]
	ystep := prob[rand.Intn(len(prob))]
	w.x += xstep
	w.y += ystep
}
