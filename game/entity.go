package game

import (
	"github.com/z46-dev/gamelib/physics"
	"github.com/z46-dev/gamelib/vector"
)

func NewEntity(g *Game) (e *Entity, err error) {
	e = &Entity{
		game: g,
		body: nil,
	}

	if e.body, err = g.world.AddBody(physics.Body2Config[float64]{
		Type:  physics.DynamicBody,
		Shape: physics.NewCircle2[float64](1),
	}); err != nil {
		return
	}

	return
}

// Update guns/turrets, apply forces, etc.
func (e *Entity) Update() {
	e.body.Force.Add(vector.NewVec2[float64](1, 0))
}

// Think is called after the physics update. It's used for things like AI and controller logic
func (e *Entity) Think() {}

func (e *Entity) Destroy() {
	e.game.world.RemoveBody(e.body.ID)
}