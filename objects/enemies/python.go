package enemies

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/tessig/gogogopher/mechanics"
	"github.com/tessig/gogogopher/resources"
)

func PythonProvider() *mechanics.Enemy {
	return python(resources.Python)
}

func AlienPythonProvider() *mechanics.Enemy {
	return python(resources.AlienPython)
}

func python(img *ebiten.Image) *mechanics.Enemy {
	return mechanics.NewEnemy(
		Python,
		img,
		3,
		32,
		14,
		[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	)
}
