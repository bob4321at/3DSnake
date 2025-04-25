package snake

import (
	"main/coffe"
	"main/textures"
	"main/utils"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func GetAngle(pos1, pos2 utils.Vec2) float64 {
	offx := pos1.X - pos2.X
	offy := pos2.Y - pos1.Y
	return math.Atan2(offx, offy)
}

func GetDist(pos1, pos2 utils.Vec2) float64 {
	return math.Sqrt((pos1.X-pos2.X)*(pos1.X-pos2.X) + (pos1.Y-pos2.Y)*(pos1.Y-pos2.Y))
}

var MousePos utils.Vec2

type Snake struct {
	Pos   utils.Vec2
	Rot   float64
	Tex   *textures.Texture
	Parts []Snake_Part
}

func NewSnake(x, y float64, length int, tex string) Snake {
	snake_parts := []Snake_Part{}

	for i := 0; i < length; i++ {
		snake_parts = append(snake_parts, Snake_Part{utils.Vec2{X: x + float64(i*-10), Y: y}, textures.NewTexture("./art/snake_part.png", "")})
	}

	snake := Snake{utils.Vec2{X: x, Y: y}, 0, textures.NewTexture(tex, ""), snake_parts}
	return snake
}

func (snake *Snake) Upate() {
	snake.Rot = GetAngle(snake.Pos, MousePos)

	snake.Pos.X -= math.Sin(snake.Rot) * 4
	snake.Pos.Y += math.Cos(snake.Rot) * 4

	if ebiten.IsKeyPressed(ebiten.KeyR) {
		snake.Rot = 0
		snake.Pos = utils.Vec2{X: 100, Y: 100}
	}

	if ebiten.IsKeyPressed(ebiten.KeyE) {
		snake.Parts = append(snake.Parts, Snake_Part{utils.Vec2{X: snake.Pos.X, Y: snake.Pos.Y}, textures.NewTexture("./art/snake_part.png", "")})
	}

	for pi := 1; pi < len(snake.Parts); pi++ {
		part := &snake.Parts[pi]
		if GetDist(part.Pos, snake.Parts[pi-1].Pos) > 16 {
			part.Pos.X -= math.Sin(GetAngle(part.Pos, snake.Parts[pi-1].Pos)) * 4
			part.Pos.Y += math.Cos(GetAngle(part.Pos, snake.Parts[pi-1].Pos)) * 4
		}
	}

	for ci := 0; ci < len(coffe.Coffe_List); ci++ {
		coffee := &coffe.Coffe_List[ci]
		if snake.Pos.X+32 > coffee.Pos.X && snake.Pos.X < coffee.Pos.X+96 {
			if snake.Pos.Y+32 > coffee.Pos.Y && snake.Pos.Y < coffee.Pos.Y+96 {
				snake.Parts = append(snake.Parts, Snake_Part{utils.Vec2{X: snake.Parts[len(snake.Parts)-1].Pos.X, Y: snake.Parts[len(snake.Parts)-1].Pos.Y}, textures.NewTexture("./art/snake_part.png", "")})
				utils.RemoveArrayElement(ci, &coffe.Coffe_List)
			}
		}
	}

	first_part := &snake.Parts[0]
	first_part.Pos = utils.Vec2{X: snake.Pos.X - 32 - math.Sin(snake.Rot)*-32, Y: snake.Pos.Y - 32 - -math.Cos(snake.Rot)*-32}
}

func (snake *Snake) Draw(screen *ebiten.Image) {
	for spi := 0; spi < len(snake.Parts); spi++ {
		part := &snake.Parts[spi]
		part.Draw(screen)
	}

	snake_op := ebiten.DrawImageOptions{}
	snake_op.GeoM.Scale(2, 2)
	snake_op.GeoM.Translate(-float64(snake.Tex.GetTexture().Bounds().Max.X), -float64(snake.Tex.GetTexture().Bounds().Max.Y))
	snake_op.GeoM.Rotate(((GetAngle(snake.Pos, MousePos)) - math.Pi/2))
	snake_op.GeoM.Translate(snake.Pos.X, snake.Pos.Y)
	snake.Tex.Draw(screen, &snake_op)

}

type Snake_Part struct {
	Pos utils.Vec2
	Tex *textures.Texture
}

func (part *Snake_Part) Draw(screen *ebiten.Image) {
	snake_op := ebiten.DrawImageOptions{}
	snake_op.GeoM.Scale(2, 2)
	snake_op.GeoM.Translate(part.Pos.X, part.Pos.Y)
	part.Tex.Draw(screen, &snake_op)
}
