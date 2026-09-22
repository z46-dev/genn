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
func (e *Entity) Define(def *configs.Definition) {
	for _, parent := range def.Parents {
		e.Define(parent)
	}

	if def.Index != nil {
		e.Index = *def.Index
	}

	if def.Label != nil {
		e.Label = *def.Label
	}

	if def.Name != nil {
		e.Name = *def.Name
	}

	if def.Type != nil {
		e.Type = *def.Type
	}

	if def.Shape != nil {
		e.Shape = *def.Shape
	}

	if def.Color != nil {
		e.Color = *def.Color
	}

	if def.Controllers != nil {
		
	}
}

// Update guns/turrets, apply forces, etc.
func (e *Entity) Update() {
	e.Velocity.Add(e.DeltaV)
	e.Game.HSHG.Insert(e)
}

// Think is called after the physics update. It's used for things like AI and controller logic
func (e *Entity) Think() {}

func (e *Entity) Life() {}

func (e *Entity) RefreshBodyAttributes() {}

func (e *Entity) UpdateAABB() {
	var eX, eY float64 = e.Position.X + e.Velocity.X + e.DeltaV.X, e.Position.Y + e.Velocity.Y + e.DeltaV.Y
	e.AABB.X1 = min(e.Position.X, eX) - e.Size
	e.AABB.Y1 = min(e.Position.Y, eY) - e.Size
	e.AABB.X2 = max(e.Position.X, eX) + e.Size
	e.AABB.Y2 = max(e.Position.Y, eY) + e.Size
}

func (e *Entity) Destroy() {
	e.Game.Entities.Remove(e.ID)
}
