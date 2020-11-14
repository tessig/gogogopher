package enemies

import (
	"github.com/tessig/gogogopher/mechanics"
	"github.com/tessig/gogogopher/resources"
)

var AlienProvider = func() *mechanics.Enemy {
	enemy := mechanics.NewEnemy(
		Alien,
		resources.Alien,
		3,
		34,
		32,
		[]int{
			0, 1, 2, // closed
			3, 4, 5, // opening
			6, 7, 8,
			6, 7, 8,
			6, 7, 8,
			9, 10, 11, 12, 13, 14, // closing
		},
	)
	enemy.Mover.Gravity = 0
	return enemy
}
