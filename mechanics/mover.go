package mechanics

import (
	"github.com/SolarLune/resolv"
)

const (
	MoverDirectionRight = iota
	MoverDirectionLeft
)

type (
	MoverDirection int

	Mover struct {
		resolv.Shape
		Vx        float64
		Vy        float64
		Gravity   float64
		Direction MoverDirection
		onRamp    bool
	}
)

func (m *Mover) CollideMove(collisions *resolv.Space) (int32, int32) {
	// gravity
	if m.Vy < 20 {
		m.Vy += m.Gravity
	}

	solids := collisions.FilterByTags(CollisionGroupSolid)
	ramps := collisions.FilterByTags(CollisionGroupRamp)

	dx, dy := int32(m.Vx), int32(m.Vy)
	// Check horizontal collision
	if res := solids.Resolve(m, dx, 0); res.Colliding() {
		dx = res.ResolveX
	}
	m.Move(dx, 0)

	// Check vertical collision
	m.onRamp = true
	// first ramps with additional velocity to stick to them
	res := ramps.Resolve(m, 0, dy+6)
	// if no ramps, normal collision against solids
	if !res.Colliding() {
		m.onRamp = false
		res = solids.Resolve(m, 0, dy)
	}
	if res.Colliding() {
		dy = res.ResolveY
		m.Vy = 0
	}
	m.Move(0, dy)
	m.determineDirection()

	return dx, dy
}

func (m *Mover) determineDirection() {
	switch {
	case m.Vx > 0:
		m.Direction = MoverDirectionRight
	case m.Vx < 0:
		m.Direction = MoverDirectionLeft
	}
}
