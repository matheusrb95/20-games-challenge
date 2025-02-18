package scenes

import (
	"fmt"
	"image/color"
	"math"

	"github.com/matheusrb95/20-games-challenge/pong/constants"
	"github.com/matheusrb95/20-games-challenge/pong/entities"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type GameScene struct {
	loaded bool

	player1 *entities.Player
	player2 *entities.Player
	ball    *entities.Ball
}

func NewGameScene() *GameScene {
	return &GameScene{
		player1: entities.NewPlayer(10.0, 95.0, 10.0, 50.0),
		player2: entities.NewPlayer(300.0, 95.0, 10.0, 50.0),
		ball:    entities.NewBall(constants.ScreenWidth/2, constants.ScreenHeight/2, 3.0),
	}
}

func (g *GameScene) Load() {
	g.loaded = true
}

func (g *GameScene) Loaded() bool {
	return g.loaded
}

func (g *GameScene) Update() SceneId {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return PauseSceneId
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.player1.Y -= 1 * constants.PlayerSpeed
	} else if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.player1.Y += 1 * constants.PlayerSpeed
	}

	if g.player1.Y+g.player1.Height >= constants.ScreenHeight {
		g.player1.Y = constants.ScreenHeight - g.player1.Height
	} else if g.player1.Y <= 0.0 {
		g.player1.Y = 0.0
	}

	if g.ball.Dx > 0 {
		if g.ball.Y > g.player2.Y {
			g.player2.Y += 1 * constants.PlayerSpeed
		} else if g.ball.Y < g.player2.Y+g.player2.Height {
			g.player2.Y -= 1 * constants.PlayerSpeed
		}
	}

	if g.player2.Y+g.player2.Height >= constants.ScreenHeight {
		g.player2.Y = constants.ScreenHeight - g.player2.Height
	} else if g.player2.Y <= 0.0 {
		g.player2.Y = 0.0
	}

	g.ball.X += 1 * g.ball.Dx * g.ball.Speed
	g.ball.Y += 1 * g.ball.Dy * g.ball.Speed

	if g.ball.Y+g.ball.Radius >= constants.ScreenHeight || g.ball.Y-g.ball.Radius <= 0 {
		g.ball.Dy *= -1
		g.ball.Speed += 0.2
	}

	if checkCollision(g.ball.Circle, g.player1.Rect) {
		g.ball.Dx *= -1
		g.ball.Dy = calculateReboundAngle(g.ball.Y, g.player1.Y, g.player1.Height)
		g.ball.Speed += 0.2
	}

	if checkCollision(g.ball.Circle, g.player2.Rect) {
		g.ball.Dx *= -1
		g.ball.Dy = calculateReboundAngle(g.ball.Y, g.player2.Y, g.player2.Height)
		g.ball.Speed += 0.2
	}

	if g.ball.X > constants.ScreenWidth {
		g.player1.Score++
		g.RecenterBall()
	} else if g.ball.X < 0 {
		g.player2.Score++
		g.RecenterBall()
	}

	return GameSceneId
}

func (g *GameScene) Draw(screen *ebiten.Image) {
	vector.StrokeRect(screen, 0, 0, 320, 240, 3, color.RGBA{255, 255, 255, 255}, false)

	x := constants.ScreenWidth / 2
	y1 := 0.0
	y2 := float64(constants.ScreenHeight)
	distance := 0.0
	for distance < y2-y1 {
		startY := y1 + distance
		endY := startY + constants.DashLength

		if endY > y2 {
			endY = y2
		}

		vector.StrokeLine(screen, float32(x), float32(startY), float32(x), float32(endY), 1.0, color.White, false)

		distance += constants.DashLength + constants.GapLength
	}

	vector.DrawFilledRect(screen, g.player1.X, g.player1.Y, g.player1.Width, g.player1.Height, color.White, false)
	vector.DrawFilledRect(screen, g.player2.X, g.player2.Y, g.player2.Width, g.player2.Height, color.White, false)
	vector.DrawFilledCircle(screen, g.ball.X, g.ball.Y, g.ball.Radius, color.White, false)

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d", g.player1.Score), constants.ScreenWidth/4, 10)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d", g.player2.Score), constants.ScreenWidth*3/4, 10)
}

func (g *GameScene) Layout(outsideWidth, outsideHeight int) (int, int) {
	return constants.ScreenWidth, constants.ScreenHeight
}

func (g *GameScene) RecenterBall() {
	g.ball = entities.NewBall(constants.ScreenWidth/2, constants.ScreenHeight/2, 3.0)
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

func calculateReboundAngle(BallY, playerY, playerHeight float32) float32 {
	relativeIntersect := (playerY + (playerHeight / 2)) - BallY
	normalizedRelativeIntersect := relativeIntersect / (playerHeight / 2)
	maxAngle := float32(math.Pi / 4)

	return normalizedRelativeIntersect * maxAngle
}
