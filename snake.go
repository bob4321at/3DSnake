package main

import (
	"main/textures"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func GetAngle(pos1, pos2 Vec2) float64 {
	offx := pos1.x - pos2.x
	offy := pos2.y - pos1.y
	return math.Atan2(offx, offy)
}

func GetDist(pos1, pos2 Vec2) float64 {
	return math.Sqrt((pos1.x-pos2.x)*(pos1.x-pos2.x) + (pos1.y-pos2.y)*(pos1.y-pos2.y))
}

var MousePos Vec2

type Snake struct {
	Pos   Vec2
	Rot   float64
	Tex   *textures.Texture
	Parts []Snake_Part
}

func NewSnake(x, y float64, length int, tex string) Snake {
	snake_parts := []Snake_Part{}

	for i := 0; i < length; i++ {
		snake_parts = append(snake_parts, Snake_Part{Vec2{x + float64(i*-10), y}, textures.NewTexture("./art/snake_part.png", "")})
	}

	snake := Snake{Vec2{x, y}, 0, textures.NewTexture(tex, ""), snake_parts}
	return snake
}

func (snake *Snake) Upate() {
	snake.Rot = GetAngle(snake.Pos, MousePos)

	snake.Pos.x -= math.Sin(snake.Rot) * 4
	snake.Pos.y += math.Cos(snake.Rot) * 4

	if ebiten.IsKeyPressed(ebiten.KeyR) {
		snake.Rot = 0
		snake.Pos = Vec2{100, 100}
	}

	if ebiten.IsKeyPressed(ebiten.KeyE) {
		snake.Parts = append(snake.Parts, Snake_Part{Vec2{snake.Pos.x, snake.Pos.y}, textures.NewTexture("./art/snake_part.png", "")})
	}

	for pi := 1; pi < len(snake.Parts); pi++ {
		part := &snake.Parts[pi]
		if GetDist(part.Pos, snake.Parts[pi-1].Pos) > 16 {
			part.Pos.x -= math.Sin(GetAngle(part.Pos, snake.Parts[pi-1].Pos)) * 4
			part.Pos.y += math.Cos(GetAngle(part.Pos, snake.Parts[pi-1].Pos)) * 4
		}
	}

	first_part := &snake.Parts[0]
	first_part.Pos = Vec2{snake.Pos.x - 32 - math.Sin(snake.Rot)*-32, snake.Pos.y - 32 - -math.Cos(snake.Rot)*-32}
}

func (snake *Snake) Draw(screen *ebiten.Image) {
	for spi := 0; spi < len(snake.Parts); spi++ {
		part := &snake.Parts[spi]
		part.Draw(screen)
	}

	snake_op := ebiten.DrawImageOptions{}
	snake_op.GeoM.Scale(2, 2)
	snake_op.GeoM.Translate(-float64(snake_tex.GetTexture().Bounds().Max.X), -float64(snake_tex.GetTexture().Bounds().Max.Y))
	snake_op.GeoM.Rotate(((GetAngle(snake.Pos, MousePos)) - math.Pi/2))
	snake_op.GeoM.Translate(snake.Pos.x, snake.Pos.y)
	snake.Tex.Draw(screen, &snake_op)

}

type Snake_Part struct {
	Pos Vec2
	Tex *textures.Texture
}

func (part *Snake_Part) Draw(screen *ebiten.Image) {
	snake_op := ebiten.DrawImageOptions{}
	snake_op.GeoM.Scale(2, 2)
	snake_op.GeoM.Translate(part.Pos.x, part.Pos.y)
	part.Tex.Draw(screen, &snake_op)
}
