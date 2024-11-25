package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"math/rand"
)

const (
	xdim        = 100
	ydim        = 100
	WindowXSize = 800
	WindowYSize = 600
)

var (
	cellXSize = WindowXSize / xdim
	cellYSize = WindowYSize / ydim
	recArray  [xdim][ydim]Rectangle
)

var (
	fishColor  = color.RGBA{255, 255, 0, 255} // YELLOW
	sharkColor = color.RGBA{255, 0, 0, 255}   // RED
	waterColor = color.RGBA{0, 41, 58, 255}   // Blue
)

// Rectangle struct to represent each cell
type Rectangle struct {
	x, y  int
	w, h  int
	color color.Color
}

// Game implements the Ebiten Game interface
type Game struct{}

// Update updates the game logic (called every frame)
func (g *Game) Update() error {
	// No game logic for now
	return nil
}

// Draw renders the graphics (called every frame)
func (g *Game) Draw(screen *ebiten.Image) {
	// Loop through each rectangle and draw it
	for i := 0; i < xdim; i++ {
		for k := 0; k < ydim; k++ {
			drawRectangle(screen, recArray[i][k])
		}
	}
}

// Layout specifies the screen dimensions
func (g *Game) Layout(_, _ int) (int, int) {
	return WindowXSize, WindowYSize
}

// drawRectangle draws a single rectangle on the screen
func drawRectangle(screen *ebiten.Image, rect Rectangle) {
	// Create a new rectangle image
	img := ebiten.NewImage(rect.w, rect.h)
	img.Fill(rect.color)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(rect.x), float64(rect.y))
	screen.DrawImage(img, op)
}

func drawFish(screen *ebiten.Image) {

}

func drawShark(screen *ebiten.Image) {

}

func drawWater(screen *ebiten.Image) {

}

func main() {
	// Initialize the rectangles
	for i := 0; i < xdim; i++ {
		for k := 0; k < ydim; k++ {
			var rectColor color.Color
			var num = rand.Intn(100)
			if num < 70 {
				rectColor = waterColor
			} else if num < 90 {
				rectColor = fishColor
			} else {
				rectColor = sharkColor
			}

			recArray[i][k] = Rectangle{
				x:     i * cellXSize,
				y:     k * cellYSize,
				w:     cellXSize,
				h:     cellYSize,
				color: rectColor,
			}
		}
	}

	// Create a new game instance
	game := &Game{}

	// Run the game
	ebiten.SetWindowSize(WindowXSize, WindowYSize)
	ebiten.SetWindowTitle("Go Wa-Tor World")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
