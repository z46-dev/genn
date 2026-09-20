package game

import "github.com/z46-dev/gamelib"

func NewGame(cfg Configuration) (g *Game) {
	g = &Game{
		EntitiesIDAccumulator: 0,
		Entities:              gamelib.NewCollection[*Entity](),
	}

	return
}

func (g *Game) Start() {}

func (g *Game) Stop() {}

func (g *Game) Update() {
}
