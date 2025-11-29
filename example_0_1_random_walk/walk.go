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
	xstep := rand.Intn(3) - 1
	ystep := rand.Intn(3) - 1
	w.x += xstep
	w.y += ystep
}
