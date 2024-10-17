package main

import (
	"image/color"
	"log"
	"math/rand"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 640
	screenHeight = 480
	gridSize     = 20

	frame_delay = 20

	// Direction
	DirUp    = 0
	DirRight = 1
	DirDown  = 2
	DirLeft  = 3
)

type Game struct {
	Snakes []Snake
	Foods  []Food
	done   bool

	framePause int
}

type Food struct {
	X int
	Y int
}

func (g *Game) Update() error {

	// maybe quit
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		g.done = true
	}
	if g.done {
		return ebiten.Termination
	}

	// pause for a few frames
	if g.framePause > 0 {
		g.framePause--
		return nil
	} else {
		g.framePause = frame_delay
	}

	// for each snake
	for i := range g.Snakes {

		g.Snakes[i].Update()

		// check for collision with food
		for j, food := range g.Foods {
			headX := g.Snakes[i].Xs[0]
			headY := g.Snakes[i].Ys[0]
			if headX == food.X && headY == food.Y {
				// eat food
				g.Snakes[i].Score++
				g.Foods = append(g.Foods[:j], g.Foods[j+1:]...)
				// add new food
				newFood := Food{
					X: rand.Intn(screenWidth/gridSize) * gridSize,
					Y: rand.Intn(screenHeight/gridSize) * gridSize,
				}
				g.Foods = append(g.Foods, newFood)
				// grow snake
				g.Snakes[i].Xs = append([]int{headX}, g.Snakes[i].Xs...)
				g.Snakes[i].Ys = append([]int{headY}, g.Snakes[i].Ys...)
			}
		}
	}

	// Check for collision with walls or itself
	for i := range g.Snakes {
		headX := g.Snakes[i].Xs[0]
		headY := g.Snakes[i].Ys[0]
		if headX < 0 || headX >= screenWidth || headY < 0 || headY >= screenHeight {

			// reset this snake
			g.Snakes[i].New()
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	// draw snakes
	for i := range g.Snakes {
		for j := range g.Snakes[i].Xs {
			vector.DrawFilledRect(screen, float32(g.Snakes[i].Xs[j]), float32(g.Snakes[i].Ys[j]), gridSize, gridSize, color.White, true)
		}
	}

	// draw food
	for _, food := range g.Foods {
		red := color.RGBA{0xff, 0x00, 0x00, 0xff}
		vector.DrawFilledRect(screen, float32(food.X), float32(food.Y), gridSize, gridSize, red, true)
	}

	// draw score
	scoreStr := strconv.Itoa(g.Snakes[0].Score)
	ebitenutil.DebugPrint(screen, "Score: "+scoreStr)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// main initializes the game, creating a new game instance,
// adding a snake to the game, generating 10 random food items,
// setting the window size and title, and running the game loop.
func main() {
	// rand.Seed(time.Now().UnixNano())
	rand.Seed(0)
	game := &Game{}

	// hold each frame for 5 ticks
	game.framePause = 5

	// add snake
	newSnake := &Snake{}
	newSnake.New()
	game.Snakes = append(game.Snakes, *newSnake)
	newSnake2 := &Snake{}
	newSnake2.New()
	game.Snakes = append(game.Snakes, *newSnake2)

	// add 10 random food
	foodCount := 50
	for i := 0; i < foodCount; i++ {
		newFood := Food{
			X: rand.Intn(screenWidth/gridSize) * gridSize,
			Y: rand.Intn(screenHeight/gridSize) * gridSize,
		}
		game.Foods = append(game.Foods, newFood)
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Snake Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
