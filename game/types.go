package game

import (
	"github.com/z46-dev/gamelib"
	"github.com/z46-dev/gamelib/hshg"
	"github.com/z46-dev/gamelib/vector"
	"github.com/z46-dev/genn/shared/configs"
)

type (
	Configuration struct{}

	Game struct {
		EntitiesIDAccumulator uint64
		Entities              *gamelib.Collection[*Entity]
		HSHG                  *hshg.SpatialHash2[*Entity, float64]
	}

	Gun struct {
		Body *Entity
	}

	Entity struct {
		ID                         uint64
		Game                       *Game
		Index                      configs.DefinitionID
		Master, Source             *Entity
		Position, Velocity, DeltaV *vector.Vec2[float64]
		AABB                       *hshg.AABB2[float64]
		Guns                       []*Gun
	}

	Client struct{}
)
