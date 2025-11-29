package canvas

import (
	"math"
	"syscall/js"
)

type Canvas struct {
	ctx    js.Value
	width  float64
	height float64
}

func NewCanvas() *Canvas {
	doc := js.Global().Get("document")
	canvas := doc.Call("getElementById", "mycanvas")
	ctx := canvas.Call("getContext", "2d")
	width := canvas.Get("clientWidth").Float()
	height := canvas.Get("clientHeight").Float()
	return &Canvas{ctx, width, height}
}

func (c *Canvas) Width() float64  { return c.width }
func (c *Canvas) Height() float64 { return c.height }

func (c *Canvas) Clear() {
	c.ctx.Call("clearRect", 0, 0, c.width, c.height)
}

func (c *Canvas) Fill(color string) {
	c.ctx.Set("fillStyle", color)
}

func (c *Canvas) FillRect(x, y, w, h float64) {
	c.ctx.Call(
		"fillRect",
		x,
		y,
		w,
		h,
	)
}

func (c *Canvas) FillCircle(x, y, d float64) {
	c.ctx.Call("beginPath")
	c.ctx.Call("arc", x, y, d, 0, math.Pi*2)
	c.ctx.Call("fill")
}
