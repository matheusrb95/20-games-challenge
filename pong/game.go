package main

import (
	"github.com/matheusrb95/20-games-challenge/pong/scenes"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	sceneMap      map[scenes.SceneId]scenes.Scene
	activeSceneId scenes.SceneId
}

func NewGame() *Game {
	sceneMap := map[scenes.SceneId]scenes.Scene{
		scenes.MenuSceneId:  scenes.NewMenuScene(),
		scenes.GameSceneId:  scenes.NewGameScene(),
		scenes.PauseSceneId: scenes.NewPauseScene(),
	}

	return &Game{
		sceneMap:      sceneMap,
		activeSceneId: scenes.MenuSceneId,
	}
}

func (g *Game) Update() error {
	nextSceneId := g.sceneMap[g.activeSceneId].Update()

	if nextSceneId == scenes.ExitSceneId {
		return ebiten.Termination
	}

	if nextSceneId != g.activeSceneId {
		nextScene := g.sceneMap[nextSceneId]
		if !nextScene.Loaded() {
			nextScene.Load()
		}
	}

	g.activeSceneId = nextSceneId

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneMap[g.activeSceneId].Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
