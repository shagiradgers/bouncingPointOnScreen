package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
	"simpleGameOnGo/pkg/player"
)

type Game struct {
	player       *player.Player
	windowWidth  int
	windowHeight int
}

type backgroundImage struct {
	image   *ebiten.Image
	options *ebiten.DrawImageOptions
}

var (
	bgImage *backgroundImage
)

func (g *Game) Update() error {
	fmt.Println(g.player)

	if g.player.PosX+g.player.R >= float64(g.windowWidth) {
		player.MoveSpeedX *= -1
	} else if g.player.PosX-g.player.R <= 0 {
		player.MoveSpeedX = math.Abs(player.MoveSpeedX)
	}

	if g.player.PosY+g.player.R >= float64(g.windowHeight) {
		player.MoveSpeedY *= -1
	} else if g.player.PosY-g.player.R <= 0 {
		player.MoveSpeedY = math.Abs(player.MoveSpeedY)
	}

	g.player.PosX += player.MoveSpeedX
	g.player.PosY += player.MoveSpeedY
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(bgImage.image, bgImage.options)

	// рисуем гг
	ebitenutil.DrawCircle(screen, g.player.PosX, g.player.PosY, g.player.R,
		color.RGBA{255, 0, 0, 255})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	game := &Game{
		player:       player.NewPlayer(),
		windowWidth:  1280,
		windowHeight: 720,
	}
	ebiten.SetWindowSize(game.windowWidth, game.windowHeight)
	ebiten.SetWindowTitle("chto-to")

	f, err := os.Open("bgImage.png")
	if err != nil {
		log.Fatal(err)
	}
	img, err := png.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	bgImage = &backgroundImage{
		image:   ebiten.NewImageFromImage(img),
		options: new(ebiten.DrawImageOptions),
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
