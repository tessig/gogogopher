package mechanics

import (
	"image"
	"math"

	"github.com/SolarLune/resolv"
	"github.com/hajimehoshi/ebiten/v2"
)

type (
	Enemy struct {
		*Mover
		typ                       int
		animation                 []int
		scale                     float64
		sprite                    *ebiten.Image
		spriteWidth, spriteHeight int
		animationIndex            int
	}
)

var (
	_ Object = new(Enemy)
)

func NewEnemy(typ int, sprite *ebiten.Image, scale float64, spriteWidth, spriteHeight int, animation []int) *Enemy {
	return &Enemy{
		Mover: &Mover{
			Shape: resolv.NewRectangle(
				0, 0,
				int32(spriteWidth)*int32(scale), int32(spriteHeight)*int32(scale),
			),
			Gravity: .7,
			Vx:      -2,
		},
		typ:          typ,
		scale:        scale,
		animation:    animation,
		sprite:       sprite,
		spriteWidth:  spriteWidth,
		spriteHeight: spriteHeight,
	}
}

func (e *Enemy) Type() int {
	return e.typ
}

func (e *Enemy) Width() float64 {
	return float64(e.spriteWidth) * e.scale
}

func (e *Enemy) Height() float64 {
	return float64(e.spriteHeight) * e.scale
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	posX, posY := e.GetXY()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(e.scale, e.scale)
	if e.Direction == MoverDirectionRight {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(e.spriteWidth)*e.scale, 0)
	}
	op.GeoM.Translate(float64(posX), float64(posY))

	i := (e.animationIndex / 10) % len(e.animation)

	x0 := e.animation[i] * e.spriteWidth

	screen.DrawImage(e.sprite.SubImage(image.Rect(x0, 0, x0+e.spriteWidth, e.spriteHeight)).(*ebiten.Image), op)
}

func (e *Enemy) Update(char *Character, collisions *resolv.Space) error {
	e.animationIndex++
	if e.animationIndex > math.MaxInt64-100 {
		e.animationIndex = 0
	}

	if !char.IsDead {
		if char.WouldBeColliding(e, 1, 0) ||
			char.WouldBeColliding(e, -1, 0) ||
			char.WouldBeColliding(e, 0, -1) {
			char.Die()
		} else if char.WouldBeColliding(e, 0, 1) {
			char.Jump()
			return ErrDestroy
		}
	}

	oldVx, oldVy := e.Vx, e.Vy
	stopper := collisions.FilterByTags(CollisionGroupStopper)
	if res := stopper.Resolve(e, int32(e.Vx), 0); res.Colliding() {
		e.Vx = float64(res.ResolveX)
	}

	dx, dy := e.CollideMove(collisions)
	e.Vx = oldVx

	if oldVx != 0 && dx == 0 {
		e.Vx = -oldVx
	}

	if oldVy != 0 && dy == 0 {
		e.Vy = -oldVy
	}

	return nil
}
