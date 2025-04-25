package coffe

import (
	"main/textures"
	"main/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

var coffe_tex = textures.NewTexture("./art/coffe.png", "")

type Coffe struct {
	Pos utils.Vec2
}

func NewCoffe(pos utils.Vec2) Coffe {
	c := Coffe{pos}

	return c
}

func (c *Coffe) Draw(s *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(c.Pos.X, c.Pos.Y)
	coffe_tex.Draw(s, &op)
}

var Coffe_List = []Coffe{}
