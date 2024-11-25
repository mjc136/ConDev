package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
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
	rectImg   *ebiten.Image
)

var (
	fishColor  = color.RGBA{255, 255, 0, 255} // YELLOW
	sharkColor = color.RGBA{255, 0, 0, 255}   // RED
	waterColor = color.RGBA{0, 41, 58, 255}   // BLUE
)

// Rectangle struct to represent each cell
type Rectangle struct {
	x, y  int
	w, h  int
	color color.Color
}

// Game implements the Ebiten Game interface
type Game struct{}

func (g *Game) Update() error {
	start := time.Now() // Measure update time
	for i := 0; i < xdim; i++ {
		for k := 0; k < ydim; k++ {
			rect := &recArray[i][k]
			if rect.color == fishColor {
				moveEntity(i, k, fishColor)
			} else if rect.color == sharkColor {
				moveEntity(i, k, sharkColor)
			}
		}
	}
	log.Printf("Update took: %s", time.Since(start))
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	start := time.Now() // Measure draw time
	for i := 0; i < xdim; i++ {
		for k := 0; k < ydim; k++ {
			drawRectangle(screen, recArray[i][k])
		}
	}
	log.Printf("Draw took: %s", time.Since(start))
}

func (g *Game) Layout(_, _ int) (int, int) {
	return WindowXSize, WindowYSize
}

func drawRectangle(screen *ebiten.Image, rect Rectangle) {
	rectImg.Fill(rect.color)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(rect.x), float64(rect.y))
	screen.DrawImage(rectImg, op)
}

func moveEntity(x, y int, entityColor color.Color) {
	dir := rand.Intn(4)
	newX, newY := x, y

	switch dir {
	case 0:
		newY = (y - 1 + ydim) % ydim
	case 1:
		newX = (x + 1) % xdim
	case 2:
		newY = (y + 1) % ydim
	case 3:
		newX = (x - 1 + xdim) % xdim
	}

	if recArray[newX][newY].color == waterColor {
		log.Printf("Entity at (%d, %d) moved to (%d, %d)", x, y, newX, newY)
		recArray[newX][newY].color = entityColor
		recArray[x][y].color = waterColor
	} else {
		log.Printf("Entity at (%d, %d) could not move", x, y)
	}
}

func main() {
	// Initialize rectangle image for reuse
	rectImg = ebiten.NewImage(cellXSize, cellYSize)

	// Initialize the grid
	for i := 0; i < xdim; i++ {
		for k := 0; k < ydim; k++ {
			var rectColor color.Color
			num := rand.Intn(100)
			if num < 90 {
				rectColor = waterColor
			} else if num < 97 {
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

	game := &Game{}
	ebiten.SetWindowSize(WindowXSize, WindowYSize)
	ebiten.SetWindowTitle("Go Wa-Tor World")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
