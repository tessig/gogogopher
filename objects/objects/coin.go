package objects

import (
	"github.com/tessig/gogogopher/mechanics"
	"github.com/tessig/gogogopher/resources"
)

func CoinProvider() *mechanics.Collectable {
	return mechanics.NewCollectable(
		Coin,
		resources.Coin,
		2,
		16,
		16,
		[]int{0, 1, 2, 3, 4, 5, 6, 7},
	)
}
