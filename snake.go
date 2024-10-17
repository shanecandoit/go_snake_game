package main

import (
	"math/rand"
)

type Snake struct {
	Xs    []int
	Ys    []int
	Dir   int
	Score int
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

// func (snake *Snake) Draw(screen *ebiten.Image) {
// 	for i := range snake.Xs {
// 		op := &ebiten.DrawImageOptions{}
// 		op.GeoM.Translate(float64(snake.Xs[i]), float64(snake.Ys[i]))
// 		// screen.DrawImage(snakeImage, op)

// 		// vector.DrawFilledRect(screen, float32(snake.Xs[j]), float32(g.Snakes[i].Ys[j]), gridSize, gridSize, color.White, true)
// 	}
// }
