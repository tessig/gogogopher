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
	)
}
