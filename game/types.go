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
		Time                  float64
		EntitiesIDAccumulator uint64
		Entities              *gamelib.Collection[*Entity]
		HSHG                  *hshg.SpatialHash2[*Entity, float64]
	}

	HealthType struct {
		Max, Amount, Resist, Regen float64
		Dynamic                    bool
	}

	Skill struct {
		Raw, Caps                    [SKILLS_LENGTH]int
		MaxHealth                    float64
		BodyDamage                   float64
		BulletSpeed                  float64
		BulletHealth                 float64
		BulletPenetration            float64
		BulletDamage                 float64
		Reload                       float64
		MovementSpeed                float64
		ShieldRegeneration           float64
		ShieldCapacity               float64
		Ghost                        float64
		Level, Points                int
		LevelScore, Score, Deduction float64
		CanUpgrade                   bool
	}

	Gun struct {
		Body, Master                             *Entity
		Children                                 []*Entity
		CanShoot, Autofire, WaitToCycle, AltFire bool
		Length, BaseWidth, EndWidth              float64
		Offset                                   *vector.Vec2[float64]
		Angle, Delay                             float64
		AnimPos, AnimMotion                      float64
		Cycle, TrueRecoil                        float64
		BulletTypes                              []*configs.Definition
		Calculator                               configs.GunCalcName
		Settings                                 configs.GunStats
		MaxChildren                              int
		LastShot                                 struct {
			Time, Power float64
		}
	}

	Control struct {
		Target, Goal    *vector.Vec2[float64]
		Main, Alt, Fire bool
		Power           float64
	}

	Entity struct {
		ID                         uint64
		Game                       *Game
		Index                      configs.DefinitionID
		Master, Source, Parent     *Entity
		Position, Velocity, DeltaV *vector.Vec2[float64]
		AABB                       *hshg.AABB2[float64]
		Size, Facing               float64
		Guns                       []*Gun
		ActivationTime             float64
		Health, Shield             *HealthType
		Skill                      *Skill
		Invulnerable               bool
		Children                   []*Entity
		MaxChildren                int
		Control                    *Control
	}

	Client struct{}
)
