package main

import (
	"image"
	_ "image/png"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/tessig/gogogopher/mechanics"
	"github.com/tessig/gogogopher/objects"
	"github.com/tessig/gogogopher/resources"
	"github.com/tessig/gogogopher/scenes"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	screenWidth  = 640
	screenHeight = 480
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Go Go Gopher")

	game := mechanics.NewGame(
		scenes.NewTitleScene(screenWidth, screenHeight),
		scenes.NewLevelScene(
			screenWidth,
			screenHeight,
			mechanics.NewCharacter(
				resources.GopherSprite,
				resources.JumpPlayer,
				resources.HurtPlayer,
				12,
				14,
				4,
				// Front: 0 to 0
				// FrontBlink: 1 to 1
				// LookUp: 2 to 2
				// LeftStand: 3 to 7
				// LeftRight: 4 to 6
				// LeftBlink: 7 to 7
				// Walk: 8 to 15
				// Run: 16 to 23
				// Jump: 24 to 26
				// Dead: 27 to 27
				map[mechanics.CharState][]int{
					mechanics.CharStateStand:   {0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 2, 2, 2, 2, 0, 0, 0, 0},
					mechanics.CharStateWalk:    {8, 9, 10, 11, 12, 13, 14, 15},
					mechanics.CharStateRun:     {16, 17, 18, 20, 21, 22},
					mechanics.CharStateCollide: {8, 9, 10, 3, 3, 3, 7, 3, 3, 7, 3, 5, 4, 4},
					mechanics.CharStateJump:    {24},
					mechanics.CharStateAir:     {25},
					mechanics.CharStateFall:    {26},
					mechanics.CharStateDead:    {27},
				},
			),
			&scenes.Symbol{
				Img:   resources.GopherEmojis.SubImage(image.Rect(0, 0, 96, 96)).(*ebiten.Image),
				Scale: .25,
			},
			&scenes.Symbol{
				Img:   resources.Coin.SubImage(image.Rect(0, 0, 16, 16)).(*ebiten.Image),
				Scale: 1.5,
			},
			resources.CoinPlayer,
			resources.LifePlayer,
			resources.BGPlayer[resources.MusicVictory],
			&objects.Provider{},
		),
		scenes.NewGameOverScene(screenWidth, screenHeight,
			func() *scenes.Symbol {
				x0 := 29 % 7 * 96
				y0 := 29 / 7 * 96
				return &scenes.Symbol{
					Img:   resources.GopherEmojis.SubImage(image.Rect(x0, y0, x0+96, y0+96)).(*ebiten.Image),
					Scale: 1,
				}
			}(),
			func() *scenes.Symbol {
				x0 := 23 % 7 * 96
				y0 := 23 / 7 * 96
				return &scenes.Symbol{
					Img:   resources.GopherEmojis.SubImage(image.Rect(x0, y0, x0+96, y0+96)).(*ebiten.Image),
					Scale: .5,
				}
			}(),
		),
		scenes.NewCreditsScene(screenWidth, screenHeight),
		scenes.NewQuitScene(),
	)
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
