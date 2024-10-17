package main

import (
	"math/rand"
)

type Brain struct {
	// input layer
	// these should all be normalized to [0, 1]
	PosX float32
	PosY float32
	Dir  int

	// wall distance
	WallDistUp    float32
	WallDistRight float32
	WallDistDown  float32
	WallDistLeft  float32

	// food distance
	FoodDistUp    float32
	FoodDistRight float32
	FoodDistDown  float32
	FoodDistLeft  float32

	// TODO hidden layers

	// output layer
	// turn probabilities
	ProbLeft  float32
	ProbUp    float32
	ProbRight float32
	ProbDown  float32
}

type Snake struct {
	Xs    []int
	Ys    []int
	Dir   int
	Score int
	Turns int

	// distance to wall in each dir
	WallDistUp    int
	WallDistRight int
	WallDistDown  int
	WallDistLeft  int

	// distance to food in each dir
	FoodDistUp    int
	FoodDistRight int
	FoodDistDown  int
	FoodDistLeft  int

	// maybe dont worry about this
	// distance to tail or other snake in each dir
	// TailDistUp    float32
	// TailDistRight float32
	// TailDistDown  float32
	// TailDistLeft  float32
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

	snake.FoodDistRight = screenWidth
	snake.FoodDistLeft = screenWidth
	snake.FoodDistDown = screenHeight
	snake.FoodDistUp = screenHeight

	snake.WallDistUp = screenHeight
	snake.WallDistRight = screenWidth
	snake.WallDistDown = screenHeight
	snake.WallDistLeft = screenWidth
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

	// inc turns, lifetime
	snake.Turns++
}

// func (snake *Snake) Draw(screen *ebiten.Image) {
// 	for i := range snake.Xs {
// 		op := &ebiten.DrawImageOptions{}
// 		op.GeoM.Translate(float64(snake.Xs[i]), float64(snake.Ys[i]))
// 		// screen.DrawImage(snakeImage, op)

// 		// vector.DrawFilledRect(screen, float32(snake.Xs[j]), float32(g.Snakes[i].Ys[j]), gridSize, gridSize, color.White, true)
// 	}
// }
