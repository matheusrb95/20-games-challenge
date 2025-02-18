package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/matheusrb95/20-games-challenge/pong/entities"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 320
	screenHeight = 240
	dashLength   = 10.0
	gapLength    = 5.3
	playerSpeed  = 5.0
)

type Game struct {
	Player1 *entities.Player
	Player2 *entities.Player
	Ball    *entities.Ball
}

func NewGame() *Game {
	return &Game{
		Player1: entities.NewPlayer(10.0, 95.0, 10.0, 50.0),
		Player2: entities.NewPlayer(300.0, 95.0, 10.0, 50.0),
		Ball:    entities.NewBall(screenWidth/2, screenHeight/2, 3.0),
	}
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.Player1.Y -= 1 * playerSpeed
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.Player1.Y += 1 * playerSpeed
	}

	if g.Player1.Y+g.Player1.Height >= screenHeight {
		g.Player1.Y = screenHeight - g.Player1.Height
	} else if g.Player1.Y <= 0.0 {
		g.Player1.Y = 0.0
	}

	if g.Ball.Dx > 0 {
		if g.Ball.Y > g.Player2.Y {
			g.Player2.Y += 1 * playerSpeed
		} else if g.Ball.Y < g.Player2.Y+g.Player2.Height {
			g.Player2.Y -= 1 * playerSpeed
		}
	}

	if g.Player2.Y+g.Player2.Height >= screenHeight {
		g.Player2.Y = screenHeight - g.Player2.Height
	} else if g.Player2.Y <= 0.0 {
		g.Player2.Y = 0.0
	}

	g.Ball.X += 1 * g.Ball.Dx * g.Ball.Speed
	g.Ball.Y += 1 * g.Ball.Dy * g.Ball.Speed

	if g.Ball.Y+g.Ball.Radius >= screenHeight || g.Ball.Y-g.Ball.Radius <= 0 {
		g.Ball.Dy *= -1
		g.Ball.Speed += 0.2
	}

	if checkCollision(g.Ball.Circle, g.Player1.Rect) {
		g.Ball.Dx *= -1
		g.Ball.Dy = calculateReboundAngle(g.Ball.Y, g.Player1.Y, g.Player1.Height)
		g.Ball.Speed += 0.2
	}

	if checkCollision(g.Ball.Circle, g.Player2.Rect) {
		g.Ball.Dx *= -1
		g.Ball.Dy = calculateReboundAngle(g.Ball.Y, g.Player2.Y, g.Player2.Height)
		g.Ball.Speed += 0.2
	}

	if g.Ball.X > screenWidth {
		g.Player1.Score++
		g.RecenterBall()
	} else if g.Ball.X < 0 {
		g.Player2.Score++
		g.RecenterBall()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	vector.StrokeRect(screen, 0, 0, 320, 240, 3, color.RGBA{255, 255, 255, 255}, false)

	x := screenWidth / 2
	y1 := 0.0
	y2 := float64(screenHeight)
	distance := 0.0
	for distance < y2-y1 {
		startY := y1 + distance
		endY := startY + dashLength

		if endY > y2 {
			endY = y2
		}

		vector.StrokeLine(screen, float32(x), float32(startY), float32(x), float32(endY), 1.0, color.White, false)

		distance += dashLength + gapLength
	}

	vector.DrawFilledRect(screen, g.Player1.X, g.Player1.Y, g.Player1.Width, g.Player1.Height, color.White, false)
	vector.DrawFilledRect(screen, g.Player2.X, g.Player2.Y, g.Player2.Width, g.Player2.Height, color.White, false)
	vector.DrawFilledCircle(screen, g.Ball.X, g.Ball.Y, g.Ball.Radius, color.White, false)

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d", g.Player1.Score), screenWidth/4, 10)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d", g.Player2.Score), screenWidth*3/4, 10)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) RecenterBall() {
	g.Ball = entities.NewBall(screenWidth/2, screenHeight/2, 3.0)
}

func checkCollision(c entities.Circle, r entities.Rect) bool {
	closestX := clamp(c.X, r.X, r.X+r.Width)
	closestY := clamp(c.Y, r.Y, r.Y+r.Height)

	distanceX := c.X - closestX
	distanceY := c.Y - closestY
	distance := math.Sqrt(float64(distanceX*distanceX + distanceY*distanceY))

	return distance <= float64(c.Radius)
}

func clamp(value, min, max float32) float32 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func calculateReboundAngle(ballY, playerY, playerHeight float32) float32 {
	relativeIntersect := (playerY + (playerHeight / 2)) - ballY
	normalizedRelativeIntersect := relativeIntersect / (playerHeight / 2)
	maxAngle := float32(math.Pi / 4)

	return normalizedRelativeIntersect * maxAngle
}
