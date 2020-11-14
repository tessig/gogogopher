package scenes

import (
	"image"
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/exp/shiny/materialdesign/colornames"

	"github.com/tessig/gogogopher/mechanics"
	"github.com/tessig/gogogopher/music"
	"github.com/tessig/gogogopher/resources"
)

type (
	Title struct {
		width, height int

		faceAnimation   []int
		faceSize        int
		cursorAnimation []int
		runnerAnimation []int
		cursorPos       mechanics.SceneType
		gopherWidth     int
		gopherHeight    int
		tileSize        int
		cloudY          int
		animationIndex  int
		menu            []menuEntry
		grace           bool
		showTitle       bool
		showMenu        bool
	}

	menuEntry struct {
		sceneType mechanics.SceneType
		label     string
	}
)

var (
	_ mechanics.Scene = new(Title)
)

func NewTitleScene(w, h int) *Title {
	t := &Title{
		width:  w,
		height: h,
		faceAnimation: []int{
			0, 0, 12,
			0, 0, 6,
			0, 0, 27,
			0, 0, 2,
			12, 13, 13,
			0, 0, 4,
			0, 0, 7,
			0, 0, 26,
			0, 0, 30,
			0, 0, 32,
		},
		faceSize:        96,
		cursorAnimation: []int{8, 9, 10, 11, 12, 13, 14, 15},
		runnerAnimation: []int{16, 17, 18, 20, 21, 22},
		gopherWidth:     12,
		gopherHeight:    14,
		tileSize:        64,
		cloudY:          170,
		menu: []menuEntry{
			{
				sceneType: mechanics.SceneTypeLevel,
				label:     "Start Game",
			},
			{
				sceneType: mechanics.SceneTypeCredits,
				label:     "Credits",
			},
			{
				sceneType: mechanics.SceneTypeQuit,
				label:     "Quit",
			},
		},
	}
	return t
}

func (t *Title) Init() {
	music.SetTrack(resources.BGPlayer[resources.MusicThemeSongFull])
	music.Play()
	t.graceTime()
}

func (t *Title) Layout(int, int) (int, int) {
	return t.width, t.height
}

func (t *Title) Update() (mechanics.SceneType, error) {
	t.animationIndex++
	if t.animationIndex > math.MaxInt64-100 {
		t.animationIndex = 0
	}
	if t.animationIndex > 100 {
		t.showTitle = true
	}
	if t.animationIndex > 220 {
		t.showMenu = true
	}

	// we have to wait a bit after cursor movement so that single ups and downs can be steered by the user
	if t.showMenu && !t.grace {
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			if t.cursorPos > 0 {
				t.cursorPos--
				t.graceTime()
			}
		}
		if ebiten.IsKeyPressed(ebiten.KeyDown) {
			if t.cursorPos < mechanics.SceneType(len(t.menu)-1) {
				t.cursorPos++
				t.graceTime()
			}
		}
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			return t.menu[t.cursorPos].sceneType, nil
		}
	}

	return mechanics.SceneTypeTitle, nil
}

func (t *Title) graceTime() {
	t.grace = true
	go func() {
		time.Sleep(200 * time.Millisecond)
		t.grace = false
	}()
}

func (t *Title) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 0x80, G: 0xa0, B: 0xc0, A: 0xff})

	t.drawScene(screen)

	title := "GO GO GOPHER!"
	height := resources.ArcadeFont.Metrics().Height.Floor()
	y := t.faceSize + 2*height

	if t.showTitle {
		scale := 1.0
		if !t.showMenu {
			scale = 8 - math.Min(7*float64(t.animationIndex-100)/100, 7)
		}
		boundString := text.BoundString(resources.ArcadeFont, title)
		titleImg := ebiten.NewImage(boundString.Dx(), height)
		text.Draw(titleImg, title, resources.ArcadeFont, 0, height, color.White)
		x := (t.width - int(scale*float64(len(title)*height))) / 2
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(scale, scale)
		op.GeoM.Translate(float64(x), float64(y-height))
		colorScale := 1 - (scale-1)/7
		op.ColorM.Scale(colorScale, 1, 1, colorScale)
		screen.DrawImage(titleImg, op)
	}
	if t.showMenu {
		t.drawFace(screen)

		t.drawMenu(screen, y+2*height)

		msg := []string{
			"by Thorsten Essig",
			"github.com/tessig/gogogopher",
			"",
			"Go Gopher by Renee French is",
			"licenced under CC BY 3.0",
		}
		height = resources.SmallArcadeFont.Metrics().Height.Floor()
		for i, line := range msg {
			x := (t.width - len(line)*height) / 2
			text.Draw(screen, line, resources.SmallArcadeFont, x, t.height-4-(len(msg)-i-1)*(height+4), colornames.Grey300)
		}
	}
}

func (t *Title) drawScene(screen *ebiten.Image) {
	// Gopher
	i := (t.animationIndex / 10) % len(t.runnerAnimation)
	pose := t.runnerAnimation[i]
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(-4, 4)
	op.GeoM.Translate(float64(100+4*t.gopherWidth), float64(t.height-2*t.tileSize-4*t.gopherHeight))
	screen.DrawImage(resources.GopherSprite.SubImage(image.Rect(pose*t.gopherWidth, 0, (pose+1)*t.gopherWidth, t.gopherHeight)).(*ebiten.Image), op)

	// Cloud
	op.GeoM.Reset()
	diff := 2 * t.animationIndex % (t.width + 6*t.tileSize)
	x0 := t.width + 3*t.tileSize - diff
	if diff == 0 {
		t.cloudY += rand.Intn(51) - 25
		if t.cloudY > t.height-3*t.tileSize-4*t.gopherHeight || t.cloudY < 0 {
			t.cloudY = 170
		}
	}
	op.GeoM.Translate(float64(x0), float64(t.cloudY))
	op.ColorM.Scale(1, 1, 1, .7)
	screen.DrawImage(resources.ForestTiles.SubImage(image.Rect(2*t.tileSize, 4*t.tileSize, 3*t.tileSize, 5*t.tileSize)).(*ebiten.Image), op)
	op.GeoM.Translate(float64(t.tileSize), 0)
	screen.DrawImage(resources.ForestTiles.SubImage(image.Rect(3*t.tileSize, 4*t.tileSize, 4*t.tileSize, 5*t.tileSize)).(*ebiten.Image), op)
	op.ColorM.Reset()

	// Ground
	for x := -t.tileSize; x < t.width+t.tileSize; x = x + t.tileSize {
		op.GeoM.Reset()
		x0 := x - 8*t.animationIndex%t.width
		if x0 < -t.tileSize {
			x0 += t.width + t.tileSize
		}
		op.GeoM.Translate(float64(x0), float64(t.height-2*t.tileSize))
		screen.DrawImage(resources.PlainsTiles.SubImage(image.Rect(3*t.tileSize, 0, 4*t.tileSize, t.tileSize)).(*ebiten.Image), op)
		op.GeoM.Translate(0, float64(t.tileSize))
		screen.DrawImage(resources.PlainsTiles.SubImage(image.Rect(3*t.tileSize, t.tileSize, 4*t.tileSize, 2*t.tileSize)).(*ebiten.Image), op)
	}
}

func (t *Title) drawMenu(screen *ebiten.Image, y int) {
	height := resources.SmallArcadeFont.Metrics().Height.Floor()

	maxLength := 0
	for _, entry := range t.menu {
		if maxLength < len(entry.label) {
			maxLength = len(entry.label)
		}
	}

	x0 := (t.width - maxLength*height) / 2
	y0 := y
	for i, line := range t.menu {
		text.Draw(screen, line.label, resources.SmallArcadeFont, x0, y0+i*(height+20), color.White)
	}

	i := (t.animationIndex / 10) % len(t.cursorAnimation)
	pose := t.cursorAnimation[i]

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(-1, 1)
	op.GeoM.Translate(float64(x0-10), float64(y0-height+int(t.cursorPos)*(height+20)))
	screen.DrawImage(resources.GopherSprite.SubImage(image.Rect(pose*t.gopherWidth, 0, (pose+1)*t.gopherWidth, t.gopherHeight)).(*ebiten.Image), op)
}

func (t *Title) drawFace(screen *ebiten.Image) {
	i := (t.animationIndex / 100) % len(t.faceAnimation)
	pose := t.faceAnimation[i]

	x0 := (pose % 7) * t.faceSize
	y0 := (pose / 7) * t.faceSize
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(t.width/2-t.faceSize/2), 0)
	screen.DrawImage(resources.GopherEmojis.SubImage(image.Rect(x0, y0, x0+t.faceSize, y0+t.faceSize)).(*ebiten.Image), op)
}
