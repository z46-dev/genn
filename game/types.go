package game

import "github.com/z46-dev/gamelib/physics"

type (
	Configuration struct{}

	Game struct {
		world *physics.World2[float64]
	}

	Gun struct{}

	Entity struct {
		game *Game
		body *physics.Body2[float64]
	}

	Client struct{}
)
