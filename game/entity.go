package game

import (
	"github.com/z46-dev/gamelib/vector"
	"github.com/z46-dev/genn/shared/configs"
)

const runSpeed float64 = 1.5

// Take the pointer of a value and return it.
func ptr[T any](v T) (pv *T) {
	pv = &v
	return
}

// Lazy ternary implementation. Returns a if cond is true, otherwise returns b.
func ternary[T any](cond bool, a, b T) (out T) {
	if cond {
		out = a
	} else {
		out = b
	}

	return
}

func eParent(e *Entity) (p *Parent) {
	p = &Parent{
		IsGun:  false,
		Entity: e,
		Gun:    nil,
	}

	return
}

func gParent(g *Gun) (p *Parent) {
	p = &Parent{
		IsGun:  true,
		Entity: nil,
		Gun:    g,
	}

	return
}

func NewEntity(g *Game, pos *vector.Vec2[float64], master *Entity) (e *Entity) {
	e = &Entity{
		Game:     g,
		ID:       g.EntitiesIDAccumulator,
		Position: pos,
		Master:   master,
	}

	e.Source = e
	e.Parent = eParent(e)
	if master == nil {
		e.Master = e
	}

	g.EntitiesIDAccumulator++
	e.Game.Entities.Add(e)
	return
}

// Define sets the entity's properties based on a definition
func (e *Entity) Define(def *configs.Definition) {}

// Update guns/turrets, apply forces, etc.
func (e *Entity) Update() {
	e.Velocity.Add(e.DeltaV)
	e.Game.HSHG.Insert(e)
}

// Think is called after the physics update. It's used for things like AI and controller logic
func (e *Entity) Think() {}

func (e *Entity) Destroy() {
	e.Game.Entities.Remove(e.ID)
}
