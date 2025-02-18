package scenes

import "github.com/hajimehoshi/ebiten/v2"

type SceneId uint

const (
	MenuSceneId SceneId = iota
	GameSceneId
	PauseSceneId
	ExitSceneId
)

type Scene interface {
	Update() SceneId
	Draw(screen *ebiten.Image)
	Load()
	Loaded() bool
}
