package mechanics

import (
	"errors"

	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten/v2"
)

type (
	ObjectProvider interface {
		CreateEnemy(typ int) *Enemy
		CreateObject(typ int) Object
	}

	Object interface {
		resolv.Shape
		Draw(screen *ebiten.Image)
		Update(char *Character, collisions *resolv.Space) error
		Width() float64
		Height() float64
	}
)

var (
	ErrDestroy = errors.New("object is destroyed")
)

func PositionInTiles(x, y int, tileSize, width, height float64) (int32, int32) {
	var dx, dy float64 = 0, 0
	if width < tileSize {
		dx = (tileSize - width) / 2
	}
	if height < tileSize {
		dy = (tileSize - height) / 2
	}
	x0 := float64(x)*tileSize + dx
	y0 := float64(y)*tileSize + dy

	return int32(x0), int32(y0)
}
