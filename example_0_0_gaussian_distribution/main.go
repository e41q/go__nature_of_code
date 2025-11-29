package main

import (
	"e41q/noc_exercises/canvas"
	"e41q/noc_exercises/p5math"
	"syscall/js"

	"go.mway.dev/x/math/clamp"
)

func main() {
	// Создаём canvas через наш удобный пакет
	c := canvas.NewCanvas()

	var (
		renderFrame  js.Func
		randomCounts [40]int
	)

	total := len(randomCounts)

	renderFrame = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		gaussVal := p5math.RandomGaussian(float64(total)/2, float64(total)*0.1)
		index := clamp.Clamp(int(gaussVal), 0, total)
		randomCounts[index]++

		c.Clear()

		c.Fill("grey")
		barWidth := c.Width() / float64(total)
		for i, count := range randomCounts {
			x := float64(i) * barWidth
			y := c.Height() - float64(count)
			h := float64(count)
			c.FillRect(x, y, barWidth-1, h)
		}

		js.Global().Call("requestAnimationFrame", renderFrame)
		return nil
	})
	defer renderFrame.Release()

	js.Global().Call("requestAnimationFrame", renderFrame)

	select {}
}
