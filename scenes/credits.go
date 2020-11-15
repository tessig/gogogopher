package scenes

import (
	"image"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/tessig/gogogopher/mechanics"
	"github.com/tessig/gogogopher/music"
	"github.com/tessig/gogogopher/resources"
)

type (
	Credits struct {
		width, height int
		cameraY       int
		maxY          int
		faceSize      int
		wait          bool
	}
)

var (
	_ mechanics.Scene = new(Credits)
)

func NewCreditsScene(width, height int) *Credits {
	q := &Credits{
		width:    width,
		height:   height,
		cameraY:  -20,
		faceSize: 96,
		wait:     true,
	}

	return q
}

func (c *Credits) Init() {
	music.SetTrack(resources.BGPlayer[resources.MusicPlatform])
	music.Play()
	c.cameraY = -20
	c.wait = true
}

func (c *Credits) Layout(int, int) (int, int) {
	if c.wait {
		go func() {
			time.Sleep(5 * time.Second)
			c.wait = false
		}()
	}
	return c.width, c.height
}

func (c *Credits) Update() (mechanics.SceneType, error) {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		c.wait = true
		return mechanics.SceneTypeTitle, nil
	}
	if c.cameraY > c.maxY {
		c.wait = true
	}

	if !c.wait {
		c.cameraY++
	}
	return mechanics.SceneTypeCredits, nil
}

func (c *Credits) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	y := c.drawFace(screen, 32, -c.cameraY, 1)

	title := "GO GO GOPHER!"
	y = c.headline(screen, title, y)

	msg := []string{
		"A game written in Go by Thorsten Essig",
		"github.com/tessig/gogogopher",
		"",
		"Based on the Ebiten game engine",
		"by Hajime Hoshi",
		"github.com/hajimehoshi/ebiten",
		"",
		"Go Gopher by Renee French is",
		"licenced under CC BY 3.0",
	}
	c.centeredText(screen, msg, y)

	y = c.height - c.cameraY

	y = c.drawFace(screen, 34, y, .5)
	y = c.headline(screen, "Graphics", y)
	table := [][2]string{
		{"Gopher Sprite", "Egon Elbre"},
		{"Gopher Emojis", "Egon Elbre"},
		{"Elephpant", "Sebastian Niesen"},
		{"Python", "Sebastian Niesen"},
		{"The Duke", "Sebastian Niesen"},
		{"Ferris", "Sebastian Niesen"},
		{"Alien", "Sebastian Niesen"},
		{"Key", "Sebastian Niesen"},
		{"Coin", "Puddin"},
		{"Gopher Plains", "aekiro"},
		{"The Forest", "Tio Aimar"},
		{"", ""},
		{"Graphics tidying", "Sebastian Niesen"},
	}
	y = c.tableText(screen, table, y)

	msg = []string{
		"The PHP Elephpant by Vincent Pontier",
		"is licenced under GNU GPL",
		"",
		"The Java Duke was open sourced",
		"by Sun under the New BSD",
		"",
		"The Rust Ferris by Karen Rustad Tölva",
		"is dedicated to public domain under CC0",
		"",
		"The Lisp Alien by Conrad Barski",
		"is free to any usage",
	}
	y = c.centeredText(screen, msg, y+20)

	y = c.drawFace(screen, 2, y, .5)
	y = c.headline(screen, "Music", y)
	msg = []string{
		"in alphabetical order",
		"",
		`"A Little Journey"`, "by shiru8bit",
		"", "",
		`"Birthday Cake"`, "composed, performed,", "mixed and mastered", "by Viktor Kraus",
		"", "",
		`"Funny Chase"`, "by wyver9",
		"", "",
		`"Green Hills"`, "by Igor Gundarev",
		"", "",
		`"Platform"`, "by Roald Strauss", "IndieGameMusic.com",
		"", "",
		`"Proper Summer"`, "by shiru8bit",
		"", "",
		`"Retro No Hope"`, "Music by Cleyton Kauffman", "soundcloud.com/cleytonkauffman",
		"", "",
		`"Spring Thing"`, "by shiru8bit",
		"", "",
		`"Theme Song"`, "by nene",
		"", "",
		`"Under the sun"`, "by shiru8bit",
		"", "",
		`"Victory"`, "by celestialghost8",
		"", "",
	}
	y = c.centeredText(screen, msg, y)

	y = c.drawFace(screen, 27, y, .5)
	y = c.headline(screen, "Sounds", y)
	table = [][2]string{
		{"Coin", "Luke.RUSTLTD"},
		{"Jump", "Jesús Lastra"},
		{"Life Pickup", "Jesús Lastra"},
		{"Die", "Baŝto"},
	}
	y = c.tableText(screen, table, y)

	y = c.drawFace(screen, 1, y, .5)
	y = c.headline(screen, "Fonts", y)
	msg = []string{
		`"Press Start 2P"`, "by Cody \"CodeMan38\" Boisclair",
	}
	y = c.centeredText(screen, msg, y)

	y = c.drawFace(screen, 26, y, .5)
	y = c.headline(screen, "Level Design", y)
	msg = []string{
		"Thorsten Essig",
	}
	y = c.centeredText(screen, msg, y)

	y = c.drawFace(screen, 14, y+c.height/2, 1)
	y = c.headline(screen, "Thanks for playing!", y)

	y += c.height
	y = c.drawFace(screen, 33, y, .5)
	y0 := y
	y += 50
	msg = []string{
		"For complete asset licence information",
		"please refer to the readme file in",
		"github.com/tessig/gogogopher",
	}
	y = c.centeredText(screen, msg, y)

	c.maxY = y + c.cameraY - c.height + y - y0
}

func (c *Credits) headline(screen *ebiten.Image, title string, y int) int {
	height := resources.ArcadeFont.Metrics().Height.Floor()
	x := (c.width - len(title)*height) / 2
	y += 2 * height
	text.Draw(screen, title, resources.ArcadeFont, x, y, color.White)
	y += height
	return y
}

func (c *Credits) drawFace(screen *ebiten.Image, pose int, y int, scale float64) int {
	x0 := (pose % 7) * c.faceSize
	y0 := (pose / 7) * c.faceSize
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(float64(c.width/2-int(float64(c.faceSize)*scale)/2), float64(y))
	screen.DrawImage(resources.GopherEmojis.SubImage(image.Rect(x0, y0, x0+c.faceSize, y0+c.faceSize)).(*ebiten.Image), op)
	y += int(float64(c.faceSize) * scale)
	return y
}

func (c *Credits) centeredText(screen *ebiten.Image, msg []string, y int) int {
	height := resources.SmallArcadeFont.Metrics().Height.Floor()
	for _, line := range msg {
		y += height + 4
		x := (c.width - len(line)*height) / 2
		text.Draw(screen, line, resources.SmallArcadeFont, x, y, color.White)
	}
	return y + height + 20
}

func (c *Credits) tableText(screen *ebiten.Image, table [][2]string, y int) int {
	height := resources.SmallArcadeFont.Metrics().Height.Floor()
	maxLength := 0
	for _, entry := range table {
		if maxLength < len(entry[0]) {
			maxLength = len(entry[0])
		}
	}
	x0 := c.width/2 - maxLength*height - 10
	for _, line := range table {
		y += height + 20
		text.Draw(screen, line[0], resources.SmallArcadeFont, x0, y, color.White)
		text.Draw(screen, line[1], resources.SmallArcadeFont, x0+maxLength*height+20, y, color.White)
	}

	return y + height + 20
}
