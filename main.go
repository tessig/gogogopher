package main

import (
	_ "image/png"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tessig/gogogopher/setup"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	screenWidth  = 640
	screenHeight = 480
)

func main() {
	game := setup.NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
