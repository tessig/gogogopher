package objects

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/tessig/gogogopher/mechanics"
	"github.com/tessig/gogogopher/resources"
)

func LifeProvider() *mechanics.Collectable {
	return mechanics.NewCollectable(
		Life,
		resources.GopherEmojis.SubImage(image.Rect(0, 0, 96, 96)).(*ebiten.Image),
		.4,
		96,
		96,
		[]int{0},
	)
}
