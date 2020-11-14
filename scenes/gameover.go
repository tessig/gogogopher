package scenes

import (
	"image"
	"image/color"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/exp/shiny/materialdesign/colornames"

	"github.com/tessig/gogogopher/mechanics"
	"github.com/tessig/gogogopher/music"
	"github.com/tessig/gogogopher/resources"
)

type (
	GameOver struct {
		width, height int
		koGopher      *Symbol
		sadGopher     *Symbol

		cursorAnimation []int
		cursorPos       mechanics.SceneType
		cursorWidth     int
		cursorHeight    int
		animationIndex  int
		menu            []menuEntry
		grace           bool
	}
)

var (
	_ mechanics.Scene = new(GameOver)
)

func NewGameOverScene(w, h int, koGopher, sadGopher *Symbol) *GameOver {
	g := &GameOver{
		width:           w,
		height:          h,
		koGopher:        koGopher,
		sadGopher:       sadGopher,
		cursorAnimation: []int{8, 9, 10, 11, 12, 13, 14, 15},
		cursorWidth:     12,
		cursorHeight:    14,
		menu: []menuEntry{
			{
				sceneType: mechanics.SceneTypeTitle,
				label:     "Back to Title",
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
	return g
}

func (g *GameOver) Init() {
	music.SetTrack(resources.BGPlayer[resources.MusicRetroNoHope])
	music.Play()
	g.graceTime()
}

func (g *GameOver) Layout(int, int) (int, int) {
	return g.width, g.height
}

func (g *GameOver) Update() (mechanics.SceneType, error) {
	g.animationIndex++
	if g.animationIndex > math.MaxInt64-100 {
		g.animationIndex = 0
	}

	// we have to wait a bit after cursor movement so that single ups and downs can be steered by the user
	if !g.grace {
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			if g.cursorPos > 0 {
				g.cursorPos--
				g.graceTime()
			}
		}
		if ebiten.IsKeyPressed(ebiten.KeyDown) {
			if g.cursorPos < mechanics.SceneType(len(g.menu)-1) {
				g.cursorPos++
				g.graceTime()
			}
		}
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			return g.menu[g.cursorPos].sceneType, nil
		}
	}

	return mechanics.SceneTypeGameOver, nil
}

func (g *GameOver) graceTime() {
	g.grace = true
	go func() {
		time.Sleep(200 * time.Millisecond)
		g.grace = false
	}()
}

func (g *GameOver) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	title := "Game Over!"
	height := resources.ArcadeFont.Metrics().Height.Floor()
	x := (g.width - len(title)*height) / 2
	y := 2 * height
	text.Draw(screen, title, resources.ArcadeFont, x, y, color.White)

	y = g.drawKoGopher(screen, y+20)
	y += 2 * height

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(g.sadGopher.Scale, g.sadGopher.Scale)
	op.GeoM.Translate(80, float64(y-height))
	screen.DrawImage(g.sadGopher.Img, op)
	w, _ := g.sadGopher.Img.Size()
	op.GeoM.Translate(float64(g.width-160)-float64(w)*g.sadGopher.Scale, 0)
	screen.DrawImage(g.sadGopher.Img, op)

	height = resources.SmallArcadeFont.Metrics().Height.Floor()
	text.Draw(screen, "Try harder next time", resources.SmallArcadeFont, x, y, colornames.TealA100)
	y += height

	g.drawMenu(screen, y+50)
}

func (g *GameOver) drawMenu(screen *ebiten.Image, y int) {
	height := resources.SmallArcadeFont.Metrics().Height.Floor()

	maxLength := 0
	for _, entry := range g.menu {
		if maxLength < len(entry.label) {
			maxLength = len(entry.label)
		}
	}

	x0 := (g.width - maxLength*height) / 2
	y0 := y
	for i, line := range g.menu {
		text.Draw(screen, line.label, resources.SmallArcadeFont, x0, y0+i*(height+20), color.White)
	}

	i := (g.animationIndex / 10) % len(g.cursorAnimation)
	pose := g.cursorAnimation[i]

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(-1, 1)
	op.GeoM.Translate(float64(x0-10), float64(y0-height+int(g.cursorPos)*(height+20)))
	screen.DrawImage(resources.GopherSprite.SubImage(image.Rect(pose*g.cursorWidth, 0, (pose+1)*g.cursorWidth, g.cursorHeight)).(*ebiten.Image), op)
}

func (g *GameOver) drawKoGopher(screen *ebiten.Image, y int) int {
	op := &ebiten.DrawImageOptions{}
	w, h := g.koGopher.Img.Size()
	w = int(float64(w) * g.koGopher.Scale)
	h = int(float64(h) * g.koGopher.Scale)

	op.GeoM.Scale(g.koGopher.Scale, g.koGopher.Scale)
	op.GeoM.Translate(float64(g.width/2-w/2), float64(y))
	screen.DrawImage(g.koGopher.Img, op)

	return y + h
}
