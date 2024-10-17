package main

import (
	"image/color"
	"log"
	"math/rand"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
	gridSize     = 20

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
}

type Food struct {
	X int
	Y int
}

type Snake struct {
	Xs    []int
	Ys    []int
	Dir   int
	Score int
}

func (g *Game) Update() error {

	// maybe quit
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		g.done = true
	}
	if g.done {
		return ebiten.Termination
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
	// for _, segment := range g.snake {
	// 	ebitenutil.DrawRect(screen, float64(segment.X), float64(segment.Y), gridSize, gridSize, color.White)
	// }
	// ebitenutil.DrawRect(screen, float64(g.food.X), float64(g.food.Y), gridSize, gridSize, color.RGBA{0xff, 0x00, 0x00, 0xff})

	// draw snakes
	for i := range g.Snakes {
		for j := range g.Snakes[i].Xs {
			ebitenutil.DrawRect(screen, float64(g.Snakes[i].Xs[j]), float64(g.Snakes[i].Ys[j]), gridSize, gridSize, color.White)
		}
	}

	// draw score
	scoreStr := strconv.Itoa(g.Snakes[0].Score)
	ebitenutil.DebugPrint(screen, "Score: "+scoreStr)

	// draw food
	for _, food := range g.Foods {
		ebitenutil.DrawRect(screen, float64(food.X), float64(food.Y), gridSize, gridSize, color.RGBA{0xff, 0x00, 0x00, 0xff})
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (snake *Snake) New() {
	// snake.Xs = []int{screenWidth / 2}
	randX := rand.Intn(screenWidth/gridSize) * gridSize
	snake.Xs = []int{randX}
	// snake.Ys = []int{screenHeight / 2}
	randY := rand.Intn(screenHeight/gridSize) * gridSize
	snake.Ys = []int{randY}
	snake.Dir = rand.Intn(4)
	snake.Score = 0
}

func (snake *Snake) Update() {
	// update snake position
	headX := snake.Xs[0]
	headY := snake.Ys[0]
	// newHeadX := headX + gridSize
	// newHeadY := headY + gridSize
	newHeadX := headX
	newHeadY := headY
	switch snake.Dir {
	case DirUp:
		newHeadY = headY - gridSize
	case DirRight:
		newHeadX = headX + gridSize
	case DirDown:
		newHeadY = headY + gridSize
	case DirLeft:
		newHeadX = headX - gridSize
	}
	snake.Xs = append([]int{newHeadX}, snake.Xs[:len(snake.Xs)-1]...)
	snake.Ys = append([]int{newHeadY}, snake.Ys[:len(snake.Ys)-1]...)
}

// main initializes the game, creating a new game instance,
// adding a snake to the game, generating 10 random food items,
// setting the window size and title, and running the game loop.
func main() {
	// rand.Seed(time.Now().UnixNano())
	rand.Seed(0)
	game := &Game{
		// snake:     []ebiten.Point{{X: screenWidth / 2, Y: screenHeight / 2}},
		// direction: ebiten.Point{X: gridSize, Y: 0},
		// food:      ebiten.Point{X: rand.Intn(screenWidth/gridSize) * gridSize, Y: rand.Intn(screenHeight/gridSize) * gridSize},
	}

	// add snake
	newSnake := &Snake{}
	newSnake.New()
	game.Snakes = append(game.Snakes, *newSnake)
	newSnake2 := &Snake{}
	newSnake2.New()
	game.Snakes = append(game.Snakes, *newSnake2)

	// add 10 random food
	for i := 0; i < 10; i++ {
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
