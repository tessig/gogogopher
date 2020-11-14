package enemies

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/tessig/gogogopher/mechanics"
	"github.com/tessig/gogogopher/resources"
)

func DukeProvider() *mechanics.Enemy {
	return duke(resources.Duke)
}

func AlienDukeProvider() *mechanics.Enemy {
	return duke(resources.AlienDuke)
}

func duke(img *ebiten.Image) *mechanics.Enemy {
	enemy := mechanics.NewEnemy(
		Duke,
		img,
		5,
		12,
		14,
		[]int{0, 1, 2, 3, 4, 5, 6, 5, 4, 3, 2, 1},
	)
	enemy.Vx = -1
	return enemy
}
