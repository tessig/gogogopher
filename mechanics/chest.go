package mechanics

import (
	"image"
	"math"

	"github.com/SolarLune/resolv"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	stateClosed = iota
	stateOpening
	stateOpen
)

type (
	Chest struct {
		resolv.Shape
		scale          float64
		sprite         *ebiten.Image
		spriteWidth    int
		spriteHeight   int
		state          uint8
		animations     [3][]int
		animationIndex int
		waitFor        int
	}
)

var (
	_ Object = new(Chest)
)

func NewChest(sprite *ebiten.Image, scale float64, spriteWidth, spriteHeight int, animations [3][]int) *Chest {
	return &Chest{
		Shape: resolv.NewRectangle(
			0, 0,
			int32(float64(spriteWidth)*scale), int32(float64(spriteHeight)*scale),
		),
		scale:        scale,
		sprite:       sprite,
		spriteWidth:  spriteWidth,
		spriteHeight: spriteHeight,
		animations:   animations,
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

	animations := c.animations[c.state]
	pose := animations[(c.animationIndex/10)%len(animations)]

	x0 := int(float64(pose*c.spriteWidth) * c.scale)

	screen.DrawImage(c.sprite.SubImage(image.Rect(x0, 0, x0+c.spriteWidth, c.spriteHeight)).(*ebiten.Image), op)
}

func (c *Chest) Update(char *Character, collisions *resolv.Space) error {
	c.animationIndex++
	if c.animationIndex > math.MaxInt64-100 {
		c.animationIndex = 0
	}
	if c.waitFor > 0 {
		c.waitFor--
	}

	wins := collisions.FilterByTags(CollisionGroupWin)
	if char.IsColliding(wins) {
		switch c.state {
		case stateClosed:
			c.state = stateOpening
			c.waitFor = len(c.animations[stateOpening])
		case stateOpening:
			if c.waitFor == 0 {
				c.state = stateOpen
			}
		}
	}

	return nil
}
