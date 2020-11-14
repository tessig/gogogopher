package scenes

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/tessig/gogogopher/mechanics"
)

type (
	Quit struct {
	}
)

var (
	_ mechanics.Scene = new(Quit)
)

func NewQuitScene() *Quit {
	q := &Quit{}

	return q
}

func (q *Quit) Init() {
}

func (q *Quit) Layout(int, int) (int, int) {
	return 1, 1
}

func (q *Quit) Update() (mechanics.SceneType, error) {
	os.Exit(0)
	return mechanics.SceneTypeQuit, nil
}

func (q *Quit) Draw(*ebiten.Image) {

}
