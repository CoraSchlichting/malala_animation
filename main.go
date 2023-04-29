package main

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	background *ebiten.Image
	character  *ebiten.Image

	charX  = 0.0
	charY  = screenHeight / 2
	speedX = 2
)

type Game struct {
	count int
}

func (g *Game) Update() error {
	g.count++

	charX += float64(speedX)

	if charX >= screenWidth {
		charX = 0
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(charX, float64(charY))
	screen.DrawImage(character, op)

	backgroundWidth, backgroundHeight := background.Size()
	backgroundScaleY := float64(screenHeight) / float64(backgroundHeight)

	for i := -1; i < int(math.Ceil(float64(screenWidth)/float64(backgroundWidth))); i++ {
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Scale(1, backgroundScaleY)
		op.GeoM.Translate(float64(i*backgroundWidth)+charX, 0)
		screen.DrawImage(background, op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	var err error

	background, _, err = ebitenutil.NewImageFromFile("background.png")
	if err != nil {
		log.Fatal(err)
	}

	character, _, err = ebitenutil.NewImageFromFile("character.png")
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("2D Side Scrolling Game")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
