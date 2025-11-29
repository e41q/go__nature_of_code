package main

import (
	"e41q/noc_exercises/canvas"
	"e41q/noc_exercises/p5math"
	"errors"
	"fmt"
	"strconv"
	"syscall/js"

	"go.mway.dev/x/math/clamp"
)

func main() {
	canvas := canvas.NewCanvas()

	var renderFrame js.Func
	renderFrame = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		spread, err := getInputValue("spread")
		if err != nil {
			spread = 0.25
		}
		size, err := getInputValue("size")
		if err != nil {
			size = 0.25
		}
		alpha, err := getInputValue("alpha")
		if err != nil {
			alpha = 0.8
		}

		x := p5math.RandomGaussian(canvas.Width()/2, spread*(canvas.Height()/2))
		y := p5math.RandomGaussian(canvas.Height()/2, spread*(canvas.Height()/2))
		d := p5math.RandomGaussian(size/canvas.Height(), size)
		if d <= 0 {
			d = 10
		}

		baseHue, err := getInputValue("base-hue")
		if err != nil {
			baseHue = 250
		}
		hueSpread, err := getInputValue("hue-spread")
		if err != nil {
			hueSpread = 15
		}
		hue := p5math.RandomGaussian(baseHue, hueSpread)
		if hue < 0 {
			hue += 360
		} else if hue >= 360 {
			hue -= 360
		}
		sat := p5math.RandomGaussian(80, 20)
		sat = clamp.Clamp(sat, 0, 100)
		bright := p5math.RandomGaussian(80, 20)
		bright = clamp.Clamp(bright, 0, 100)

		canvas.Fill(fmt.Sprintf("hsla(%d, %d%%, %d%%, %.2f)", int(hue), int(sat), int(bright), alpha))
		canvas.FillCircle(x, y, d)

		js.Global().Call("requestAnimationFrame", renderFrame)
		return nil
	})

	js.Global().Call("requestAnimationFrame", renderFrame)

	resetBtn := js.Global().Get("document").Call("getElementById", "reset")
	if resetBtn.IsNull() {
		js.Global().Get("console").Call("error", "Reset button not found")
	} else {
		resetFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			canvas.Clear()
			return nil
		})

		resetBtn.Call("addEventListener", "click", resetFunc)
	}

	select {}
}

func getInputValue(inputId string) (float64, error) {
	document := js.Global().Get("document")
	spreadInput := document.Call("getElementById", inputId)
	valueStr := spreadInput.Get("value").String()
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		msg := "Invalid float value: " + valueStr
		js.Global().Get("console").Call("error", msg, "| Error:", err.Error())
		return 0, errors.New(msg)
	}
	return value, nil
}
