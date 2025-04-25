package scenes

import (
	"image/color"
	"main/coffe"
	"main/shader"
	"main/snake"
	"main/textures"
	"main/utils"
	"math"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
)

var GameScene = NewScene(1, GameSceneDraw, GameSceneUpdate, GameSceneInit)

var snake_tex *textures.Texture
var sanke = snake.NewSnake(100, 100, 5, "./art/snake_head.png")

var chicken_tex *textures.Texture

var steve_pos = utils.Vec2{X: 100, Y: 100}
var time = float64(0)

func GameSceneInit() {
	snake_tex = textures.NewTexture("./art/snake_head.png", "")
	chicken_tex = textures.NewTexture("./art/lava_chicken.png", "")
	snake_tex.SetUniforms(map[string]any{})

	coffe.Coffe_List = append(coffe.Coffe_List, coffe.NewCoffe(utils.Vec2{X: rand.Float64() * 1200, Y: rand.Float64() * 660}))
}

func GameSceneDraw(s *ebiten.Image, screen_img *ebiten.Image) {
	screen := textures.NewTexture("./art/empty.png", shader.Test_Shader)
	s.Fill(color.RGBA{100, 100, 100, 255})

	sanke.Draw(screen.GetTexture())

	for ci := 0; ci < len(coffe.Coffe_List); ci++ {
		coffe := &coffe.Coffe_List[ci]
		coffe.Draw(screen.GetTexture())
	}

	op := ebiten.DrawImageOptions{}
	screen.Draw(s, &op)
}

func GameSceneUpdate() {
	sanke.Upate()
	time += 0.01
	mx, my := ebiten.CursorPosition()
	snake.MousePos.X = float64(mx)
	snake.MousePos.Y = float64(my)

	if math.Mod(utils.GameTime, 200) == 0 {
		coffe.Coffe_List = append(coffe.Coffe_List, coffe.NewCoffe(utils.Vec2{X: rand.Float64() * 1200, Y: rand.Float64() * 660}))
	}

	steve_pos.X += math.Cos(time)
}
