package main

import (
	"github.com/hajimehoshi/ebiten/v2"

	"main/scenes"
	"main/utils"
)

type Game struct{}

func (g *Game) Update() error {
	scenes.List_Of_Scenes[scenes.Current_Scene].Update()
	utils.GameTime++

	return nil
}

func (g *Game) Draw(s *ebiten.Image) {
	scenes.List_Of_Scenes[scenes.Current_Scene].Draw(s, s)
}

func (g *Game) Layout(ow, oh int) (sw, sh int) {
	return 1280, 720
}

func main() {
	ebiten.SetWindowSize(1280, 720)

	scenes.List_Of_Scenes[0].Setup()

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
