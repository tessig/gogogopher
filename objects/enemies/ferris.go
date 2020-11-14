package enemies

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/tessig/gogogopher/mechanics"
	"github.com/tessig/gogogopher/resources"
)

func FerrisProvider() *mechanics.Enemy {
	return ferris(resources.Ferris)
}

func AlienFerrisProvider() *mechanics.Enemy {
	return ferris(resources.AlienFerris)
}

func ferris(img *ebiten.Image) *mechanics.Enemy {
	return mechanics.NewEnemy(
		Ferris,
		img,
		2,
		34,
		14,
		[]int{0, 1, 2, 3, 4},
	)
}
