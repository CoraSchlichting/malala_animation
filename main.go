package main

import (
	"image/color"
	"log"
	"strings"

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
	charY  = screenHeight - 150
	speedX = 1.25
)

type Game struct {
	count int
}

func (g *Game) Update() error {
	g.count++

	charX += float64(speedX)

	// Reset the charX position once image is done passing
	if charX >= float64(background.Bounds().Dx()) {
		charX = 0
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	bgImgPoint := background.Bounds().Size()
	backgroundScaleY := float64(screenHeight) / float64(bgImgPoint.Y)
	backgroundScaleX := .5
	var xMovement float64
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(backgroundScaleX, backgroundScaleY)

	xMovement = -charX
	op.GeoM.Translate(xMovement, 0)
	screen.DrawImage(background, op)

	drawMalala(screen)
	printStory(screen, int(xMovement))
}

func printStory(screen *ebiten.Image, screenLocation int) {
	if screenLocation < -100 && screenLocation > -500 {
		drawTextBox(screen, "Malala was the youngest Nobel peace price winner!! What would happen if this was a really long string with lots and lots of text, I guess we'd just have to see, how far can we go!? CAN WE GO FARTHER?!? I THINK WE CAN!!!!!! OOOOOOOOOOHHHHWWWWEEEEEEEE RICK, I THINK, OH GEE RICK I THINK WE'RE GOING TO BREAK THIS STRING!!!!!! *BURPRPRPRP* You- you know what your problem is, Morty? You don't know a char pointer when you see one *burp* that's right, there are no strings, Morty! It's arrays! arrays of arrays morty! *burp* They're everywhere! And you can get unsafe morty, you can start doing pointer arithmatic if that's your thing, but you gotta be careful morty, you gotta be careful! You can't just go around dereferencing pointers all willy nilly, you gotta be careful morty!")
	}
}

func drawTextBox(screen *ebiten.Image, text string) {
	// Draw a semi-transparent black rectangle over the entire screen
	rect := ebiten.NewImage(screenWidth, screenHeight)
	screen.DrawImage(rect, nil)

	// Draw a border around the text box
	borderWidth := 2.0
	borderColor := color.NRGBA{0, 0, 0, 0}
	rect = ebiten.NewImage(int(screenWidth-(borderWidth*2)), 50)
	rect.Fill(borderColor)
	rect.Fill(color.NRGBA{0, 0, 0, 128})

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(borderWidth, float64(screenHeight-50-borderWidth))

	defaultFontPixelWidth := 12
	maxCharsPerLine := (screenWidth / defaultFontPixelWidth) * 2
	currentY := 0
	for i := 0; i < len(text); i += maxCharsPerLine {
		var stringVal string
		if i+maxCharsPerLine > len(text) {
			stringVal = text[i:]
		} else {
			stringVal = text[i : i+maxCharsPerLine]
		}

		ebitenutil.DebugPrintAt(rect, strings.TrimLeft(stringVal, " "), 0, currentY)

		currentY += defaultFontPixelWidth
	}

	screen.DrawImage(rect, op)

}

func drawMalala(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.05, 0.05)
	op.GeoM.Translate(10, float64(charY))
	screen.DrawImage(character, op)
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
	ebiten.SetWindowTitle("Malala's Story")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
