package main

import (
	"e41q/noc_exercises/canvas"
	"math/rand"
	"syscall/js"
)

func main() {
	c := canvas.NewCanvas()

	var (
		renderFrame  js.Func
		randomCounts [40]int
	)

	total := float64(len(randomCounts))

	renderFrame = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		arVal := int(acceptReject() * total)
		randomCounts[arVal]++

		c.Clear()

		c.Fill("grey")
		barWidth := c.Width() / total
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

func acceptReject() float64 {
	for {
		r1 := rand.Float64()
		probability := r1
		r2 := rand.Float64()

		if r2 < probability {
			return r1
		}
	}
}
