package game

import "github.com/z46-dev/gamelib/physics"

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
