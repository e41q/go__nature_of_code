package main

import (
	"e41q/noc_exercises/canvas"
	"syscall/js"
)

func main() {
	canvas := canvas.NewCanvas()

	var renderFrame js.Func
	renderFrame = js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		// Очищаем холст
		canvas.Clear()

		// Рисуем анимированный синий прямоугольник
		x, y := 0.0, 0.0

		canvas.Fill("blue")
		canvas.FillRect(x, y, 60, 40)

		js.Global().Call("requestAnimationFrame", renderFrame)
		return nil
	})

	// Start running
	js.Global().Call("requestAnimationFrame", renderFrame)

	select {}
}
