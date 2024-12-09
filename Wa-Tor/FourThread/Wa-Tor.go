package main

import (
	"encoding/csv"
	"fmt"
	"image/color"
	"math/rand"
	"os"
	"strconv"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Simulation parameters.
const (
	xdim        = 150  // Grid width.
	ydim        = 150  // Grid height.
	WindowXSize = 750  // Window width in pixels.
	WindowYSize = 600  // Window height in pixels.
	NumShark    = 15   // Starting population of sharks.
	NumFish     = 1000 // Starting population of fish.
	fishBreed   = 5    // Steps required for fish to reproduce.
	sharkBreed  = 10   // Steps required for sharks to reproduce.
	sharkStarve = 7    // Steps before a shark starves without eating.
)

var (
	cellXSize = WindowXSize / xdim  // Width of each cell in pixels.
	cellYSize = WindowYSize / ydim  // Height of each cell in pixels.
	recArray  [xdim][ydim]Rectangle // Grid representing the simulation world.
	rectImg   *ebiten.Image         // Shared rectangle image used for drawing cells.

	fishColor  = color.RGBA{255, 255, 0, 255} // Color representing fish (yellow).
	sharkColor = color.RGBA{255, 0, 0, 255}   // Color representing sharks (red).
	waterColor = color.RGBA{0, 41, 58, 255}   // Color representing water (blue).
)

// Rectangle represents a rectangular cell in the simulation grid.
type Rectangle struct {
	x, y   int         // Top-left position of the rectangle.
	w, h   int         // Width and height of the rectangle.
	color  color.Color // Color of the rectangle (fish, shark, or water).
	starve int         // Starvation counter for sharks.
	breed  int         // Breeding counter for both fish and sharks.
}

// Game implements the Ebiten Game interface for the Wa-Tor simulation.
type Game struct {
	frameCount  int
	tpsSum      float64
	csvWriter   *csv.Writer
	threadCount int
}

// Update updates the state of the simulation and logs data to CSV.
func (g *Game) Update() error {
	g.frameCount++
	currentTPS := ebiten.ActualTPS()
	g.tpsSum += currentTPS

	// Check if csvWriter is not nil before writing
	if g.csvWriter != nil {
		err := g.csvWriter.Write([]string{
			strconv.Itoa(g.frameCount),
			fmt.Sprintf("%.2f", currentTPS),
			strconv.Itoa(g.threadCount),
		})
		if err != nil {
			fmt.Println("Error writing to CSV:", err)
		}
	}

	// Use two threads to update the grid concurrently
	var wg sync.WaitGroup
	midY := ydim / 2
	midX := xdim / 2

	// Process the top-left of the grid
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < midX; i++ {
			for k := 0; k < midY; k++ {
				updateCell(i, k)
			}
		}
	}()

	// Process the top-right of the grid
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := midX; i < xdim; i++ {
			for k := 0; k < midY; k++ {
				updateCell(i, k)
			}
		}
	}()

	// Process the bottom-left of the grid
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < midX; i++ {
			for k := midY; k < ydim; k++ {
				updateCell(i, k)
			}
		}
	}()

	// Process the bottom-right of the grid
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := midX; i < xdim; i++ {
			for k := midY; k < ydim; k++ {
				updateCell(i, k)
			}
		}
	}()

	wg.Wait() // Wait for all 4 goroutines to complete

	return nil
}

// updateCell updates the state of a single cell based on its contents.
func updateCell(i, k int) {
	rect := &recArray[i][k]
	if rect.color == fishColor {
		moveFish(i, k)
	} else if rect.color == sharkColor {
		if rect.starve > 0 {
			moveShark(i, k)
		} else {
			rect.color = waterColor // Shark starves and the cell becomes water.
		}
	}
}

// Draw draws the simulation grid to the screen.
func (g *Game) Draw(screen *ebiten.Image) {
	for i := 0; i < xdim; i++ {
		for k := 0; k < ydim; k++ {
			drawRectangle(screen, recArray[i][k])
		}
	}

	// Draw a background rectangle for the TPS display
	tpsBackground := ebiten.NewImage(120, 30)
	tpsBackground.Fill(color.RGBA{0, 0, 0, 180})

	// Draw the background rectangle at a fixed position
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(10, 10)
	screen.DrawImage(tpsBackground, op)

	// Draw the TPS text over the background
	msg := fmt.Sprintf("TPS: %.2f", ebiten.ActualTPS())
	ebitenutil.DebugPrintAt(screen, msg, 20, 20)
}

// Layout defines the layout of the game window.
func (g *Game) Layout(_, _ int) (int, int) {
	return WindowXSize, WindowYSize
}

// drawRectangle draws a single rectangle to the screen.
func drawRectangle(screen *ebiten.Image, rect Rectangle) {
	rectImg.Fill(rect.color)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(rect.x), float64(rect.y))
	screen.DrawImage(rectImg, op)
}

// placeEntities randomly places a specified number of entities (fish or sharks)
// on the grid. Entities are placed only in empty (water) cells.
func placeEntities(num int, entityColor color.Color) {
	count := 0
	for count < num {
		x := rand.Intn(xdim)
		y := rand.Intn(ydim)

		rect := &recArray[x][y]

		if rect.color == waterColor {
			recArray[x][y].color = entityColor
			if entityColor == sharkColor {
				recArray[x][y].starve = sharkBreed // Initialize shark's starvation counter.
			}
			recArray[x][y].breed = 0 // Initialize breeding counter.
			count++
		}
	}
}

// moveEntity moves an entity to an adjacent cell in a random direction.
//
// The function wraps around the edges of the grid (toroidal behavior).
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

// moveFish moves a fish to an adjacent water cell.
//
// If the fish reaches its breeding threshold, it reproduces in its original cell.
func moveFish(x, y int) {
	newX, newY := moveEntity(x, y)
	if recArray[newX][newY].color == waterColor {
		recArray[newX][newY].color = fishColor
		recArray[x][y].color = waterColor
		recArray[newX][newY].breed = recArray[x][y].breed + 1
		recArray[x][y].breed = 0
	}
	if recArray[newX][newY].breed == fishBreed {
		recArray[x][y].color = fishColor
		recArray[x][y].breed = 0
		recArray[newX][newY].breed = 0
	}
}

// moveShark moves a shark to an adjacent cell.
//
// If a shark eats a fish, its starvation counter is reset.
// Sharks reproduce after reaching their breeding threshold.
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
	recArray[newX][newY].breed = recArray[x][y].breed + 1
	recArray[x][y].breed = 0
	if recArray[newX][newY].breed == sharkBreed {
		recArray[x][y].color = sharkColor
		recArray[x][y].breed = 0
		recArray[x][y].starve = sharkStarve
		recArray[newX][newY].breed = 0
	}
}

// checkAdjacent checks for fish in the adjacent cells.
//
// If a fish is found, it returns the coordinates of the fish. Otherwise,
// it returns the current cell's coordinates.
func checkAdjacent(x, y int) (newx, newy int) {
	newx, newy = x, y

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

// eatFish allows a shark to eat a fish at a specified cell.
//
// The shark's starvation counter is reset after eating.
func eatFish(x, y int) {
	if recArray[x][y].color == fishColor {
		recArray[x][y].color = sharkColor
		recArray[x][y].starve = sharkStarve
	}
}

// main initializes the simulation and runs the game loop.
func main() {
	rectImg = ebiten.NewImage(cellXSize, cellYSize) // Initialize shared rectangle image.

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

	placeEntities(NumFish, fishColor)   // Place initial fish.
	placeEntities(NumShark, sharkColor) // Place initial sharks.

	// Create and open the CSV file
	file, err := os.Create("tps_data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()

	// Write the header row to the CSV file
	csvWriter.Write([]string{"Frame", "TPS", "ThreadCount"})

	// Create the Game instance with csvWriter
	game := &Game{csvWriter: csvWriter, threadCount: 4}
	ebiten.SetWindowSize(WindowXSize, WindowYSize)
	ebiten.SetWindowTitle("Go Wa-Tor World")

	// Run the game loop
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
