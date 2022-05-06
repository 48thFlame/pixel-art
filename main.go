package main

import (
	"syscall/js"
)

func main() {
	document := js.Global().Get("document")
	canvas := document.Call("getElementById", "canvas")

	document.Get("body").Call("appendChild", canvas)
	ctx := canvas.Call("getContext", "2d")
	ctx.Set("fillStyle", "rgba(255, 0, 0, 1)")
	width, height := canvas.Get("width"), canvas.Get("height")

	ctx.Call("fillRect", 0, 0, width, height)

}