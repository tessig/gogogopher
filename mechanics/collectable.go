package mechanics

import (
	"image"
	"math"

	"github.com/SolarLune/resolv"
	"github.com/hajimehoshi/ebiten/v2"
)

type (
	Collectable struct {
		resolv.Shape
		typ            int
		scale          float64
		animationIndex int
		animation      []int
		sprite         *ebiten.Image
		spriteWidth    int
		spriteHeight   int
	}
)

var (
	_ Object = new(Collectable)
)

func NewCollectable(typ int, sprite *ebiten.Image, scale float64, spriteWidth, spriteHeight int, animation []int) *Collectable {
	return &Collectable{
		Shape: resolv.NewRectangle(
			0, 0,
			int32(float64(spriteWidth)*scale), int32(float64(spriteHeight)*scale),
		),
		typ:            typ,
		scale:          scale,
		animationIndex: 0,
		animation:      animation,
		sprite:         sprite,
		spriteWidth:    spriteWidth,
		spriteHeight:   spriteHeight,
	}
}

func (c *Collectable) Type() int {
	return c.typ
}

func (c *Collectable) Width() float64 {
	return float64(c.spriteWidth) * c.scale
}

func (c *Collectable) Height() float64 {
	return float64(c.spriteHeight) * c.scale
}

func (c *Collectable) Draw(screen *ebiten.Image) {
	posX, posY := c.GetXY()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(c.scale, c.scale)
	op.GeoM.Translate(float64(posX), float64(posY))

	i := (c.animationIndex / 10) % len(c.animation)

	x0 := c.animation[i] * c.spriteWidth

	screen.DrawImage(c.sprite.SubImage(image.Rect(x0, 0, x0+c.spriteWidth, c.spriteHeight)).(*ebiten.Image), op)
}

func (c *Collectable) Update(char *Character, _ *resolv.Space) error {
	c.animationIndex++
	if c.animationIndex > math.MaxInt64-100 {
		c.animationIndex = 0
	}

	if char.IsColliding(c) {
		return ErrDestroy
	}

	return nil
}
