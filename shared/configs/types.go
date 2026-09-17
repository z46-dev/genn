package configs

import "github.com/z46-dev/gamelib/vector"

type (
	Type         uint8
	DefinitionID uint16
	GunCalcName  uint8
	FacingType   uint8
	MotionType   uint8
	UpgradeTier  uint8

	GunStats struct {
		Reload      float64
		Recoil      float64
		Shudder     float64
		Size        float64
		Health      float64
		Damage      float64
		Penetration float64
		Speed       float64
		MaxSpeed    float64
		Range       float64
		Density     float64
		Spray       float64
		Resist      float64
	}

	BodyStats struct {
		Health, Damage, Penetration, Shield, Regeneration *float64
		Heterogeneity, Density, Pushability, Resist       *float64
		Acceleration, Speed                               *float64
		FOV                                               *float64
	}

	AISettings struct {
		DirectAim             bool // Aim directly at the target's current position instead of leading shots.
		IgnoreGunRange        bool // Use the entity's own movement/FOV for target acquisition instead of gun tracking.
		OwnerVisionBound      bool // Only acquire targets that are also near the master/owner.
		TargetAllDangerLevels bool // Allow lower-danger targets instead of filtering only to the highest danger tier.
		FullView              bool // Ignore turret firing arcs while searching for targets.
		IgnoreFood            bool // Do not target food/shapes.
		RandomizeStrafe       bool // Occasionally reverse orbit/strafing direction around the master.
	}

	Definition struct {
		ID         DefinitionID
		Type       *Type
		Name       *string
		Parents    []*Definition
		Guns       []*Gun
		Turrets    []*Turret
		Upgrades   [UpgradeTier_SENTINEL][]*Definition
		Body       *BodyStats
		AI         *AISettings
		MotionType *MotionType
		FacingType *FacingType
	}

	GunPosition struct { // Design of the gun
		Length    float64
		BaseWidth float64
		EndWidth  float64
		Offset    vector.Vec2[float64]
		Angle     float64
		Delay     float64
	}

	GunProperties struct { // Extra properties, includes weapon types/stats, color/skin, behavior
		Shoots         *Definition
		ShootSettings  GunStats
		StatCalculator *GunCalcName
	}

	Gun struct {
		Position   *GunPosition
		Properties *GunProperties
	}

	TurretPosition struct {
		Size   float64
		Offset vector.Vec2[float64]
		Angle  float64
		Arc    float64
		Layer  int
	}

	Turret struct {
		Position    *TurretPosition
		Definitions []*Definition
	}

	DefBuilder struct {
		Definition *Definition
	}

	GunBuilder struct {
		Gun *Gun
	}

	TurretBuilder struct {
		Turret *Turret
	}

	BodyBuilder struct {
		Body *BodyStats
	}
)
