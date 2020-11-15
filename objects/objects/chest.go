package objects

import (
	"github.com/tessig/gogogopher/mechanics"
	"github.com/tessig/gogogopher/resources"
)

func ChestProvider() *mechanics.Chest {
	return mechanics.NewChest(
		resources.Chest,
		1,
		64,
		64,
		[3][]int{
			{0},
			{1, 2, 3},
			{4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		},
	)
}
