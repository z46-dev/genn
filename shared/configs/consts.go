package configs

var idCounter DefinitionID = 0

const (
	TypeUnknown  Type = iota // Unknown, used for generic entities
	TypeTank                 // Players
	TypeBoss                 // Bosses
	TypeBullet               // Bullets
	TypeDrone                // Drones
	TypeSwam                 // Swarms
	TypeMinion               // Minions
	TypeTrap                 // Traps
	TypeFood                 // Polygons
	TypeCrasher              // Crashers, sentries, other guardian-like entities
	TypeObstacle             // Rocks or Maze Walls
)

const (
	GunCalcNameDefault GunCalcName = iota // Default, bullets
	GunCalcNameDrone
	GunCalcNameSwarm
	GunCalcNameTrap
	GunCalcNameFixedReload
	GunCalcNameThruster
	GunCalcNameSustained
)

const (
	StatNamesDefault StatNames = iota
	StatNamesSmasher
	StatNamesDrone
	StatNamesNecro
	StatNamesSwarm
	StatNamesTrap
	StatNamesGeneric
)

const (
	FacingTypeAutospin FacingType = iota
	FacingTypeTurnWithSpeed
	FacingTypeTurnWithMotion
	FacingTypeTurnWithTarget
	FacingTypeSmoothWithMotion
	FacingTypeSmoothWithTarget
	FacingTypeBound
)

const (
	MotionTypeGlide MotionType = iota
	MotionTypeMotor
	MotionTypeSwarm
	MotionTypeChase
	MotionTypeDrift
	MotionTypeBound
)

const (
	HitsOwnTypeNever HitsOwnType = iota
	HitsOwnTypeHard
	HitsOwnTypeHardWithBuffer
	HitsOwnTypeRepel
	HitsOwnTypePush
)

const (
	UpgradeTier1 UpgradeTier = iota
	UpgradeTier2
	UpgradeTier3
	UpgradeTier4
	UpgradeTier_SENTINEL
)

const (
	WeaponHealthFactor float64 = 0.5
	WeaponDamageFactor float64 = 1.5
	BaseFoodHealth     float64 = 2
	BaseFoodDamage     float64 = 1
)

const (
	ControllerDoNothing              Controller = iota
	ControllerMoveInCircles                     // Food stuff
	ControllerNearestDifferentMaster            // Targeting
	ControllerCanRepel                          // Alt repels away from target
	ControllerMapTargetToGoal                   // Map target to goal, then move towards that point
	ControllerHangOutNearMaster                 // When idle (no target) move around master
	ControllerGoToMasterTarget                  // Always map target to master target at time of fire
	ControllerGoToMasterTargetAngle             // Map target to master target at time of fire, rotate around master with gun rotation relative to the shooting entity
	ControllerBoomerang                         // Shoot then come back when at 1/2 range
	ControllerAlwaysFire                        // Autofire
	ControllerTargetSelf                        // May target self if no other target
	ControllerOnlyAcceptInArc                   // Only accept targets in the firing arc of the gun
	ControllerMapAltToFire                      // Alt = Fire
	ControllerSpin                              // Spin
	ControllerFastSpin                          // Fast Spin
	ControllerReverseSpin                       // Reverse Spin
	ControllerDontTurn                          // Don't turn
	ControllerMinion                            // Orbit around target
	ControllerFleeAtLowHealth                   // Flee when low health
)

var StatNamesConfig map[StatNames]SkillNames = map[StatNames]SkillNames{
	StatNamesDefault: {
		BodyDamage:         "Body Damage",
		MaxHealth:          "Max Health",
		BulletSpeed:        "Bullet Speed",
		BulletHealth:       "Bullet Health",
		BulletPenetration:  "Bullet Penetration",
		BulletDamage:       "Bullet Damage",
		Reload:             "Reload",
		Speed:              "Movement Speed",
		ShieldRegeneration: "Shield Regeneration",
		ShieldCapacity:     "Shield Capacity",
	},
	StatNamesSmasher: {
		BodyDamage:         "Body Damage",
		MaxHealth:          "Max Health",
		BulletSpeed:        "Size",
		BulletHealth:       "Shell Density",
		BulletPenetration:  "Shell Penetration",
		BulletDamage:       "Shell Damage",
		Reload:             "Engine Acceleration",
		Speed:              "Movement Speed",
		ShieldRegeneration: "Shield Regeneration",
		ShieldCapacity:     "Shield Capacity",
	},
	StatNamesDrone: {
		BodyDamage:         "Body Damage",
		MaxHealth:          "Max Health",
		BulletSpeed:        "Drone Speed",
		BulletHealth:       "Drone Health",
		BulletPenetration:  "Drone Penetration",
		BulletDamage:       "Drone Damage",
		Reload:             "Respawn Rate",
		Speed:              "Movement Speed",
		ShieldRegeneration: "Shield Regeneration",
		ShieldCapacity:     "Shield Capacity",
	},
	StatNamesNecro: {
		BodyDamage:         "Body Damage",
		MaxHealth:          "Max Health",
		BulletSpeed:        "Drone Speed",
		BulletHealth:       "Drone Health",
		BulletPenetration:  "Drone Penetration",
		BulletDamage:       "Drone Damage",
		Reload:             "Max Drone Count",
		Speed:              "Movement Speed",
		ShieldRegeneration: "Shield Regeneration",
		ShieldCapacity:     "Shield Capacity",
	},
	StatNamesSwarm: {
		BodyDamage:         "Body Damage",
		MaxHealth:          "Max Health",
		BulletSpeed:        "Swarm Speed",
		BulletHealth:       "Swarm Health",
		BulletPenetration:  "Swarm Penetration",
		BulletDamage:       "Swarm Damage",
		Reload:             "Reload",
		Speed:              "Movement Speed",
		ShieldRegeneration: "Shield Regeneration",
		ShieldCapacity:     "Shield Capacity",
	},
	StatNamesTrap: {
		BodyDamage:         "Body Damage",
		MaxHealth:          "Max Health",
		BulletSpeed:        "Placement Speed",
		BulletHealth:       "Trap Health",
		BulletPenetration:  "Trap Penetration",
		BulletDamage:       "Trap Damage",
		Reload:             "Reload",
		Speed:              "Movement Speed",
		ShieldRegeneration: "Shield Regeneration",
		ShieldCapacity:     "Shield Capacity",
	},
	StatNamesGeneric: {
		BodyDamage:         "Body Damage",
		MaxHealth:          "Max Health",
		BulletSpeed:        "Weapon Speed",
		BulletHealth:       "Weapon Health",
		BulletPenetration:  "Weapon Penetration",
		BulletDamage:       "Weapon Damage",
		Reload:             "Reload",
		Speed:              "Movement Speed",
		ShieldRegeneration: "Shield Regeneration",
		ShieldCapacity:     "Shield Capacity",
	},
}

const (
	SkillCapNormal  SkillCap = 9
	SkillCapSmasher SkillCap = 12
)

const (
	DamageClassDefault DamageClass = iota
	DamageClassFood
	DamageClassTanks
	DamageClassObstacles
)
