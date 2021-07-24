package mechanics

import (
	"image"
	"math"

	"github.com/SolarLune/resolv"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

const (
	CharStateStand CharState = iota
	CharStateWalk
	CharStateJump
	CharStateAir
	CharStateFall
	CharStateRun
	CharStateCollide
	CharStateDead
)

type (
	CharState uint8

	Character struct {
		*Mover

		Scale float64

		Image        *ebiten.Image
		SpriteWidth  int
		SpriteHeight int

		JumpPlayer *audio.Player
		HurtPlayer *audio.Player

		State CharState

		animationIndex int
		animations     map[CharState][]int
		onGround       bool
		onWall         bool
		IsDead         bool
	}
)

func NewCharacter(image *ebiten.Image, jumpPlayer, hurtPlayer *audio.Player, spriteWidth, spriteHeight int, scale float64, animations map[CharState][]int) *Character {
	return &Character{
		Scale:        scale,
		SpriteWidth:  spriteWidth,
		SpriteHeight: spriteHeight,
		Image:        image,
		JumpPlayer:   jumpPlayer,
		HurtPlayer:   hurtPlayer,
		State:        CharStateStand,
		Mover: &Mover{
			Direction: MoverDirectionRight,
			Shape:     resolv.NewRectangle(0, 0, int32(spriteWidth*int(scale)), int32(spriteHeight*int(scale))-1),
			Gravity:   .7,
		},
		animations: animations,
	}
}

func (c *Character) Jump() {
	if c.IsDead || !c.onGround {
		return
	}
	c.onGround = false
	_ = c.JumpPlayer.Rewind()
	c.JumpPlayer.Play()
	c.Vy = -16
}

func (c *Character) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(c.Scale, c.Scale)
	if c.Direction == MoverDirectionRight {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(c.Scale*float64(c.SpriteWidth), 0)
	}
	posX, posY := c.GetXY()
	op.GeoM.Translate(float64(posX), float64(posY))

	animation := c.animations[c.State]
	i := (c.animationIndex / 10) % len(animation)

	x0 := animation[i] * c.SpriteWidth
	screen.DrawImage(c.Image.SubImage(image.Rect(x0, 0, x0+c.SpriteWidth, c.SpriteHeight)).(*ebiten.Image), op)
}

func (c *Character) Update(collisions *resolv.Space) (int32, int32) {
	c.animationIndex++
	if c.animationIndex > math.MaxInt64-100 {
		c.animationIndex = 0
	}

	dx, dy := c.move(collisions)
	c.checkDeath(collisions.FilterByTags(CollisionGroupDie))
	c.determineState()
	// hit the breaks
	if c.Vx > 0 {
		c.Vx -= 4
	} else if c.Vx < 0 {
		c.Vx += 4
	}

	return dx, dy
}

func (c *Character) move(collisions *resolv.Space) (int32, int32) {
	solids := collisions.FilterByTags(CollisionGroupSolid)

	// check if we stand next to a wall by checking for wall collision on both sides
	right := solids.Resolve(c, 1, 0)
	left := solids.Resolve(c, -1, 0)
	c.onWall = right.Colliding() || left.Colliding()

	dx, dy := c.CollideMove(collisions)

	// Check for a collision downwards by just attempting a resolution downwards and seeing if it collides with something.
	down := collisions.FilterOutByTags(CollisionGroupStopper).Resolve(c, 0, 2)
	c.onGround = down.Colliding()

	return dx, dy
}

func (c *Character) checkDeath(deadSpace *resolv.Space) {
	if deadSpace.IsColliding(c) {
		c.Die()
	}
}

func (c *Character) Die() {
	if !c.IsDead {
		_ = c.HurtPlayer.Rewind()
		c.HurtPlayer.Play()
	}
	c.IsDead = true
	c.Vy = 0
	c.Vx = 0
}

func (c *Character) determineState() {
	vxAbs := c.Vx
	if vxAbs < 0 {
		vxAbs = -vxAbs
	}

	old := c.State
	switch {
	case c.IsDead:
		c.State = CharStateDead
	case c.Vy > 0 && !c.onGround:
		c.State = CharStateFall
	case c.Vy == 0 && !c.onGround:
		c.State = CharStateAir
	case c.Vy < 0 && !c.onGround:
		c.State = CharStateJump
	// collide state is flaky when going up a ramp
	case vxAbs > 0 && c.onWall && !c.onRamp:
		c.State = CharStateCollide
		if old != CharStateCollide {
			c.animationIndex = 0
		}
	case vxAbs > 4:
		c.State = CharStateRun
	case vxAbs > 0:
		c.State = CharStateWalk
	default:
		c.State = CharStateStand
	}
}
