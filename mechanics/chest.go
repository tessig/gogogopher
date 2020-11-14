package mechanics

import (
	"image"

	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten/v2"
)

type (
	Chest struct {
		resolv.Shape
		scale        float64
		sprite       *ebiten.Image
		spriteWidth  int
		spriteHeight int
		open         bool
	}
)

var (
	_ Object = new(Chest)
)

func NewChest(sprite *ebiten.Image, scale float64, spriteWidth, spriteHeight int) *Chest {
	return &Chest{
		Shape: resolv.NewRectangle(
			0, 0,
			int32(float64(spriteWidth)*scale), int32(float64(spriteHeight)*scale),
		),
		scale:        scale,
		sprite:       sprite,
		spriteWidth:  spriteWidth,
		spriteHeight: spriteHeight,
	}
}

func (c *Chest) Width() float64 {
	return float64(c.spriteWidth) * c.scale
}

func (c *Chest) Height() float64 {
	return float64(c.spriteHeight) * c.scale
}

func (c *Chest) Draw(screen *ebiten.Image) {
	posX, posY := c.GetXY()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(c.scale, c.scale)
	op.GeoM.Translate(float64(posX), float64(posY))

	x0 := 0
	if c.open {
		x0 = int(float64(c.spriteWidth) * c.scale)
	}

	screen.DrawImage(c.sprite.SubImage(image.Rect(x0, 0, x0+c.spriteWidth, c.spriteHeight)).(*ebiten.Image), op)
}

func (c *Chest) Update(char *Character, collisions *resolv.Space) error {
	wins := collisions.FilterByTags(CollisionGroupWin)
	if char.IsColliding(wins) {
		c.open = true
	}

	return nil
}
