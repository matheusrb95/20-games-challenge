package scenes

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/matheusrb95/20-games-challenge/pong/constants"
)

type PauseScene struct {
	loaded bool
	option int
}

func NewPauseScene() *PauseScene {
	return &PauseScene{}
}

func (p *PauseScene) Load() {
	p.loaded = true
}

func (p *PauseScene) Loaded() bool {
	return p.loaded
}

func (p *PauseScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	ebitenutil.DebugPrintAt(screen, "Game Paused", constants.ScreenWidth/2-40, 20)

	ebitenutil.DebugPrintAt(screen, "Resume Game", 20, constants.ScreenHeight/2)
	ebitenutil.DebugPrintAt(screen, "Exit", 20, constants.ScreenHeight/2+20)

	if p.option == 0 {
		vector.StrokeLine(screen, 18, constants.ScreenHeight/2+15, 90, constants.ScreenHeight/2+15, 2, color.White, false)
	}

	if p.option == 1 {
		vector.StrokeLine(screen, 18, constants.ScreenHeight/2+35, 90, constants.ScreenHeight/2+35, 2, color.White, false)
	}
}

func (p *PauseScene) Update() SceneId {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return GameSceneId
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyW) || inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyS) || inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		p.option = 1 - p.option
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		if p.option == 0 {
			return GameSceneId
		}

		if p.option == 1 {
			return ExitSceneId
		}
	}

	return PauseSceneId
}
