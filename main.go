package main

import (
	// "image"
	"fmt"
	"image/color"
	"math/rand"
	"syscall/js"
	"time"
)

const (
	pixelsWide = 12
	pixelsHigh = 8
)

func newBlankCanvas(htmlCanvasID string) *blankCanvas {
	bc := &blankCanvas{}

	bc.canvas = js.Global().Get("document").Call("getElementById", htmlCanvasID)
	bc.ctx = bc.canvas.Call("getContext", "2d")
	bc.width = bc.canvas.Get("width").Int()
	bc.height = bc.canvas.Get("height").Int()
	bc.pixelWidth = bc.width / pixelsWide
	bc.pixelHeight = bc.height / pixelsHigh
	bc.pixels = make([][]color.RGBA, pixelsHigh)

	for i := 0; i < pixelsHigh; i++ {
		bc.pixels[i] = make([]color.RGBA, pixelsWide)
	}
	return bc
}

type blankCanvas struct {
	canvas                  js.Value
	ctx                     js.Value
	width, height           int
	pixelWidth, pixelHeight int // the rect dims for a single pixel to draw to the html canvas
	pixels                  [][]color.RGBA
}

func (bc *blankCanvas) draw() {
	for rowI, row := range bc.pixels {
		for colI, col := range row {
			col.R = uint8(rand.Intn(255))
			col.G = uint8(rand.Intn(255))
			col.B = uint8(rand.Intn(255))
			col.A = uint8(rand.Intn(255))

			bc.ctx.Set("fillStyle", fmt.Sprintf("rgba(%v, %v, %v, %v)", col.R, col.G, col.B, alphaValueConverter(col.A)))
			bc.ctx.Call("fillRect", colI*bc.pixelWidth, rowI*bc.pixelHeight, bc.pixelWidth, bc.pixelHeight)
		}
	}
}

// a func to convert a number between 0-255 to a float between 0-1
func alphaValueConverter(a uint8) float64 {
	return float64(a) / 255
}

// a function to make the canvas be blank meaning gray and white pixels reprenting blankness
// func (bc *blankCanvas) makeEmpty() {
// 	for rowI, row := range bc.pixels {
// 		for colI, col := range row {

// 		}
// 	}
// }

func main() {
	rand.Seed(time.Now().UnixNano())
	canvas := newBlankCanvas("canvas")

	for {
		canvas.draw()
		time.Sleep(time.Millisecond * 100)
	}
	// blankCanvas = make([][]color.RGBA, pixelsHigh)
	// for _, row := range blankCanvas {

	// }

	// canvasWidth, canvasHeight = width.Int(), height.Int()
	// pixelWidth, pixelHeight = canvasWidth/pixelsWidth, canvasHeight/pixelsHeight

	// log.Println("canvasWidth:", canvasWidth, "canvasHeight:", canvasHeight, "pixelWidth:", pixelWidth, "pixelHeight:", pixelHeight)

	// ctx := canvas.Call("getContext", "2d")
	// ctx.Set("fillStyle", "rgba(255, 0, 0, 1)")
	// ctx.Call("fillRect", 0, 0, width, height)
	// img := image.NewRGBA(image.Rect(0, 0, pixelsWide, pixelsHigh))
}

// package main

// import (
// 	"image"
// 	"image/color"
// 	"image/png"
// 	"log"
// 	"os"
// )

// func main() {
// 	const width, height = 256, 256

// 	// Create a colored image of the given width and height.
// 	img := image.NewNRGBA(image.Rect(0, 0, width, height))

// 	for y := 0; y < height; y++ {
// 		for x := 0; x < width; x++ {
// 			img.Set(x, y, color.NRGBA{
// 				R: uint8((x + y) & 255),
// 				G: uint8((x + y) << 1 & 255),
// 				B: uint8((x + y) << 2 & 255),
// 				A: 255,
// 			})
// 		}
// 	}

// 	f, err := os.Create("image.png")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if err := png.Encode(f, img); err != nil {
// 		f.Close()
// 		log.Fatal(err)
// 	}

// 	if err := f.Close(); err != nil {
// 		log.Fatal(err)
// 	}

// }
