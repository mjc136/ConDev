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
	NumShark    = 10  // Starting population of sharks.
	NumFish     = 50  // Starting population of fish.
	GridSize    = 100 // Dimensions of the world (GridSize x GridSize).
	WindowXSize = 800 // Window width in pixels.
	WindowYSize = 600 // Window height in pixels.
)

var (
	cellXSize = WindowXSize / GridSize      // Width of each cell in pixels.
	cellYSize = WindowYSize / GridSize      // Height of each cell in pixels.
	recArray  [GridSize][GridSize]Rectangle // 2D array representing the grid of rectangles.
	rectImg   *ebiten.Image                 // Image used to draw each rectangle.
)

var (
	fishColor  = color.RGBA{255, 255, 0, 255} // Color representing fish (yellow).
	sharkColor = color.RGBA{255, 0, 0, 255}   // Color representing sharks (red).
	waterColor = color.RGBA{0, 41, 58, 255}   // Color representing water (blue).
)

// Rectangle represents a rectangular cell in the simulation grid.
type Rectangle struct {
	X, Y  int         // Position of the rectangle in the grid (top-left corner).
	W, H  int         // Width and height of the rectangle.
	Color color.Color // Color of the rectangle (fish, shark, or water).
}

// Game implements the Ebiten Game interface for the Wa-Tor simulation.
type Game struct{}

// Update updates the state of the simulation.
//
// This function is called once per frame. It updates the positions
// and states of all entities in the grid (fish and sharks).
func (g *Game) Update() error {
	for i := 0; i < GridSize; i++ {
		for k := 0; k < GridSize; k++ {
			rect := &recArray[i][k]
			if rect.Color == fishColor {
				moveFish(i, k)
			} else if rect.Color == sharkColor {
				moveShark(i, k)
			}
		}
	}
	return nil
}

// Draw draws the simulation grid to the screen.
//
// Each cell in the grid is drawn using its current color (fish, shark, or water).
func (g *Game) Draw(screen *ebiten.Image) {
	for i := 0; i < GridSize; i++ {
		for k := 0; k < GridSize; k++ {
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

// moveFish moves a fish in the grid.
//
// Fish move randomly to adjacent empty cells.
func moveFish(x, y int) {
	direction := rand.Intn(4) // Randomly pick a direction
	newX, newY := x, y

	switch direction {
	case 0: // Move north
		newY = (y - 1 + GridSize) % GridSize
	case 1: // Move east
		newX = (x + 1) % GridSize
	case 2: // Move south
		newY = (y + 1) % GridSize
	case 3: // Move west
		newX = (x - 1 + GridSize) % GridSize
	}

	// Check if the target cell is empty (water) before moving
	if recArray[newX][newY].Color == waterColor {
		// Move the fish
		recArray[newX][newY] = recArray[x][y]
		recArray[x][y].Color = waterColor // Reset the original cell to water
	}
}

// moveShark moves a shark in the grid.
//
// Sharks move randomly to adjacent empty cells.
func moveShark(x, y int) {
	direction := rand.Intn(4) // Randomly pick a direction
	newX, newY := x, y

	switch direction {
	case 0: // Move north
		newY = (y - 1 + GridSize) % GridSize
	case 1: // Move east
		newX = (x + 1) % GridSize
	case 2: // Move south
		newY = (y + 1) % GridSize
	case 3: // Move west
		newX = (x - 1 + GridSize) % GridSize
	}

	// Check if the target cell is empty (water) before moving
	if recArray[newX][newY].Color == waterColor {
		// Move the shark
		recArray[newX][newY] = recArray[x][y]
		recArray[x][y].Color = waterColor // Reset the original cell to water
	}
}

// drawRectangle draws a single rectangle to the screen.
//
// The rectangle's color represents its state (water, fish, or shark).
func drawRectangle(screen *ebiten.Image, rect Rectangle) {
	rectImg.Fill(rect.Color)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(rect.X), float64(rect.Y))
	screen.DrawImage(rectImg, op)
}

// main initializes and runs the Wa-Tor simulation.
//
// This function sets up the grid, populates it with initial entities,
// and starts the Ebiten game loop.
func main() {
	// Initialize the rectangle image for reuse
	rectImg = ebiten.NewImage(cellXSize, cellYSize)

	// Initialize the grid
	for i := 0; i < GridSize; i++ {
		for k := 0; k < GridSize; k++ {
			recArray[i][k] = Rectangle{
				X:     i * cellXSize,
				Y:     k * cellYSize,
				W:     cellXSize,
				H:     cellYSize,
				Color: waterColor, // Default to water
			}
		}
	}

	// Populate initial fish and sharks
	placeEntities(NumFish, fishColor)
	placeEntities(NumShark, sharkColor)

	game := &Game{}
	ebiten.SetWindowSize(WindowXSize, WindowYSize)
	ebiten.SetWindowTitle("Wa-Tor")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}

// placeEntities places entities (fish or sharks) randomly on the grid.
//
// It ensures that entities are only placed in empty cells.
func placeEntities(num int, entityColor color.Color) {
	count := 0
	for count < num {
		x := rand.Intn(GridSize)
		y := rand.Intn(GridSize)

		if recArray[x][y].Color == waterColor {
			recArray[x][y].Color = entityColor
			count++
		}
	}
}
