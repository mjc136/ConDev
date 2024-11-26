// Wa-Tor.go: A Wa-Tor simulation using the Ebiten game engine.
//
// This program simulates a Wa-Tor world, where fish and sharks interact
// in a 2D grid. Sharks and fish move randomly, interacting with their environment.
//
// author Michael Cullen
// student_number C00261635
// date 2024-11-25

package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

// Simulation parameters.
const (
	xdim        = 100
	ydim        = 100
	WindowXSize = 800 // Window width in pixels.
	WindowYSize = 600 // Window height in pixels.
	NumShark    = 10  // Starting population of sharks.
	NumFish     = 50  // Starting population of fish.
	fishBreed   = 3   // Steps for fish to reproduce
	sharkBreed  = 6   // Steps for sharks to reproduce
	sharkStarve = 100 // Steps before a shark starves
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
	x, y   int
	w, h   int
	color  color.Color
	starve int
	//breed  int
}

// Game implements the Ebiten Game interface
type Game struct{}

// Update updates the state of the simulation.
//
// This function is called once per frame. It updates the positions
// and states of all entities in the grid (fish and sharks).
func (g *Game) Update() error {
	for i := 0; i < xdim; i++ {
		for k := 0; k < ydim; k++ {
			rect := &recArray[i][k]
			if rect.color == fishColor {
				moveFish(i, k)
			} else if rect.color == sharkColor {
				if rect.starve > 0 {
					moveShark(i, k)
				} else {
					rect.color = waterColor
				}
			}
		}
	}
	return nil
}

// Draw draws the simulation grid to the screen.
//
// Each cell in the grid is drawn using its current color (fish, shark, or water).
func (g *Game) Draw(screen *ebiten.Image) {
	for i := 0; i < xdim; i++ {
		for k := 0; k < ydim; k++ {
			drawRectangle(screen, recArray[i][k])
		}
	}
}

// Layout defines the layout of the game window.
//
// It returns the width and height of the game window in pixels.
func (g *Game) Layout(_, _ int) (int, int) {
	return WindowXSize, WindowYSize
}

func drawRectangle(screen *ebiten.Image, rect Rectangle) {
	rectImg.Fill(rect.color)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(rect.x), float64(rect.y))
	screen.DrawImage(rectImg, op)
}

func placeEntities(num int, entityColor color.Color) {
	count := 0
	for count < num {
		x := rand.Intn(xdim)
		y := rand.Intn(ydim)

		rect := &recArray[x][y]

		if rect.color == waterColor {
			recArray[x][y].color = entityColor
			if entityColor == sharkColor {
				recArray[x][y].starve = sharkStarve // Initialise shark's starvation counter
			}
			count++
		}
	}
}

func moveEntity(x, y int) (newX, newY int) {
	dir := rand.Intn(4)
	newX, newY = x, y

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

	return newX, newY
}

func moveFish(x, y int) {
	newX, newY := moveEntity(x, y)
	if recArray[newX][newY].color == waterColor {
		recArray[newX][newY].color = fishColor
		recArray[x][y].color = waterColor
	}
}

func moveShark(x, y int) {
	newX, newY := checkAdjacent(x, y)
	if newX == x && newY == y {
		newX, newY = moveEntity(x, y)
		if recArray[newX][newY].color == waterColor {
			recArray[newX][newY].color = sharkColor
			recArray[newX][newY].starve = recArray[x][y].starve - 1
			recArray[x][y].color = waterColor
			recArray[x][y].starve = 0
		}
	} else {
		eatFish(newX, newY)
		recArray[x][y].color = waterColor
	}

}

func checkAdjacent(x, y int) (newx, newy int) {
	// Default to the current position (no fish found)
	newx, newy = x, y

	// Check the adjacent cells for a fish
	if recArray[(x+1+xdim)%xdim][y].color == fishColor {
		newx, newy = (x+1+xdim)%xdim, y // East
	} else if recArray[(x-1+xdim)%xdim][y].color == fishColor {
		newx, newy = (x-1+xdim)%xdim, y // West
	} else if recArray[x][(y+1+ydim)%ydim].color == fishColor {
		newx, newy = x, (y+1+ydim)%ydim // South
	} else if recArray[x][(y-1+ydim)%ydim].color == fishColor {
		newx, newy = x, (y-1+ydim)%ydim // North
	}

	return newx, newy
}

func eatFish(x, y int) {
	if recArray[x][y].color == fishColor {
		recArray[x][y].color = sharkColor
		recArray[x][y].starve = sharkStarve
	}
}

func main() {
	// Initialize rectangle image for reuse
	rectImg = ebiten.NewImage(cellXSize, cellYSize)

	// Initialize the grid
	for i := 0; i < xdim; i++ {
		for k := 0; k < ydim; k++ {

			recArray[i][k] = Rectangle{
				x:     i * cellXSize,
				y:     k * cellYSize,
				w:     cellXSize,
				h:     cellYSize,
				color: waterColor,
			}
		}
	}

	// Populate initial fish and sharks
	placeEntities(NumFish, fishColor)
	placeEntities(NumShark, sharkColor)

	game := &Game{}
	ebiten.SetWindowSize(WindowXSize, WindowYSize)
	ebiten.SetWindowTitle("Go Wa-Tor World")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
