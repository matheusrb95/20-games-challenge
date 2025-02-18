package entities

import "math/rand"

type Ball struct {
	Circle
	Speed  float32
	Dx, Dy float32
}

type Circle struct {
	X      float32
	Y      float32
	Radius float32
}

func NewBall(x, y, r float32) *Ball {
	dx := rand.Intn(2)
	if dx == 0 {
		dx = -1
	}

	dy := rand.Intn(2)
	if dy == 0 {
		dy = -1
	}

	return &Ball{Circle{x, y, r}, 2.0, float32(dx), float32(dy)}
}
