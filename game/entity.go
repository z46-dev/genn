package game

import "github.com/z46-dev/gamelib/vector"

const runSpeed float64 = 1.5

func NewEntity(g *Game, pos *vector.Vec2[float64], master *Entity) (e *Entity) {
	e = &Entity{
		Game:     g,
		ID:       g.EntitiesIDAccumulator,
		Position: pos,
		Master:   master,
	}

	e.Source = e
	e.Parent = e
	if master == nil {
		e.Master = e
	}

	g.EntitiesIDAccumulator++
	e.Game.Entities.Add(e)
	return
}

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
