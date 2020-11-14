package enemies

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/tessig/gogogopher/mechanics"
	"github.com/tessig/gogogopher/resources"
)

func ElephpantProvider() *mechanics.Enemy {
	return elephpant(resources.Elephpant)
}

func AlienElephpantProvider() *mechanics.Enemy {
	return elephpant(resources.AlienElephpant)
}

func elephpant(img *ebiten.Image) *mechanics.Enemy {
	enemy := mechanics.NewEnemy(
		Elephpant,
		img,
		6,
		16,
		14,
		[]int{0, 1, 2, 3, 4, 5, 6, 7},
	)
	enemy.Vx = -1
	return enemy
}
