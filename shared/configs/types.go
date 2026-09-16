package configs

import "github.com/z46-dev/gamelib/vector"

type (
	Type         uint8
	DefinitionID uint16
	GunStatKey   uint8
	GunCalcName  uint8
	FacingType   uint8
	MotionType   uint8
	GunStats     [GunStatKey_SENTINEL]float64

	Definition struct {
		ID      DefinitionID  `json:"id"`
		Type    Type          `json:"type,omitempty"`
		Name    string        `json:"name,omitempty"`
		Parents []*Definition `json:"parents,omitempty"`
		Guns    []*Gun        `json:"guns,omitempty"`
		Turrets []*Turret     `json:"turrets,omitempty"`
	}

	GunPosition struct { // Design of the gun
		Length    float64               `json:"length,omitempty"`
		BaseWidth float64               `json:"baseWidth,omitempty"`
		EndWidth  float64               `json:"endWidth,omitempty"`
		Offset    *vector.Vec2[float64] `json:"offset,omitempty"`
		Angle     float64               `json:"angle,omitempty"`
		Delay     float64               `json:"delay,omitempty"`
	}

	GunProperties struct { // Extra properties, includes weapon types/stats, color/skin, behavior
		Shoots        *Definition `json:"shoots,omitempty"`
		ShootSettings GunStats    `json:"shootSettings,omitempty"`
	}

	Gun struct {
		Position   *GunPosition   `json:"position,omitempty"`
		Properties *GunProperties `json:"properties,omitempty"`
	}

	Turret struct {
		Definitions []*Definition `json:"definitions,omitempty"`
	}
)
