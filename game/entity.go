package game

func NewEntity(g *Game) (e *Entity, err error) {
	e = &Entity{
		Game: g,
	}

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
