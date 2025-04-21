package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"

	"main/textures"
)

type Game struct{}

type Vec2 struct {
	x, y float64
}

var snake_tex *textures.Texture
var sanke = NewSnake(100, 100, 100, "./art/snake_head.png")

var chicken_tex *textures.Texture

var steve_pos = Vec2{100, 100}
var time = float64(0)

func (g *Game) Update() error {
	sanke.Upate()
	time += 0.01
	mx, my := ebiten.CursorPosition()
	MousePos.x = float64(mx)
	MousePos.y = float64(my)

	steve_pos.x += math.Cos(time)

	return nil
}

func (g *Game) Draw(s *ebiten.Image) {
	screen := textures.NewTexture("", Test_Shader)

	screen.SetUniforms(map[string]any{
		"Time": time,
	})

	s.Fill(color.White)

	sanke.Draw(screen.GetTexture())

	chicken_op := ebiten.DrawImageOptions{}
	chicken_op.GeoM.Translate(-32, 16)
	chicken_op.GeoM.Rotate(time)
	chicken_op.GeoM.Scale(2, 2)
	chicken_op.GeoM.Translate(steve_pos.x, steve_pos.y)
	chicken_tex.Draw(screen.GetTexture(), &chicken_op)

	op := ebiten.DrawImageOptions{}
	screen.Draw(s, &op)
}

func (g *Game) Layout(ow, oh int) (sw, sh int) {
	return 1280, 720
}

func main() {
	ebiten.SetWindowSize(1280, 720)

	snake_tex = textures.NewTexture("./art/snake_head.png", "")
	chicken_tex = textures.NewTexture("./art/lava_chicken.png", "")
	snake_tex.SetUniforms(map[string]any{})

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
