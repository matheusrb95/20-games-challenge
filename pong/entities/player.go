package entities

type Player struct {
	Rect
	Score int
}

type Rect struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}

func NewPlayer(x, y, w, h float32) *Player {
	return &Player{Rect{x, y, w, h}, 0}
}
