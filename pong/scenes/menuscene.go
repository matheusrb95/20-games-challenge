package scenes

import (
	"bytes"
	_ "embed"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/matheusrb95/20-games-challenge/pong/constants"
)

//go:embed bgm_aware.mp3
var bgmBytes []byte

type MenuScene struct {
	loaded bool
	option int

	audioContext *audio.Context
	bgmPlayer    *audio.Player
}

func NewMenuScene() *MenuScene {
	return &MenuScene{}
}

func (m *MenuScene) Load() {
	if bgmBytes == nil {
		m.loaded = true
		return
	}

	m.audioContext = audio.NewContext(constants.SampleRate)

	decoded, err := mp3.DecodeWithSampleRate(constants.SampleRate, bytes.NewReader(bgmBytes))
	if err != nil {
		log.Fatal(err)
	}

	m.bgmPlayer, err = m.audioContext.NewPlayer(decoded)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Sound loaded")

	m.bgmPlayer.SetVolume(0.5)
	m.bgmPlayer.Play()

	m.loaded = true
}

func (m *MenuScene) Loaded() bool {
	return m.loaded
}

func (m *MenuScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	ebitenutil.DebugPrintAt(screen, "Start Game", 20, constants.ScreenHeight/2)
	ebitenutil.DebugPrintAt(screen, "Exit", 20, constants.ScreenHeight/2+20)

	if m.option == 0 {
		vector.StrokeLine(screen, 18, constants.ScreenHeight/2+15, 90, constants.ScreenHeight/2+15, 2, color.White, false)
	}

	if m.option == 1 {
		vector.StrokeLine(screen, 18, constants.ScreenHeight/2+35, 90, constants.ScreenHeight/2+35, 2, color.White, false)
	}
}

func (m *MenuScene) Update() SceneId {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return ExitSceneId
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyW) || inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyS) || inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		m.option = 1 - m.option
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		if m.option == 0 {
			return GameSceneId
		}

		if m.option == 1 {
			return ExitSceneId
		}
	}

	return MenuSceneId
}
