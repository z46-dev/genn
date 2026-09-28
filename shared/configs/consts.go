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
	GunCalcNameFixedReload
	GunCalcNameThruster
	GunCalcNameSustained
	GunCalcNameNecro
	GunCalcNameTrap
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
	FacingTypeLocksFacing
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

const (
	RoomCellTypeNorm     RoomCellType = iota // Normal/Default type, empty from features
	RoomCellTypeNest                         // Valuable nest/spawn for food/crashers/protectors
	RoomCellTypeRock                         // Larger rocks, sparse
	RoomCellTypeRoid                         // Smaller rocks, dense
	RoomCellTypeWall                         // Wall occupies entire cell
	RoomCellTypeBas0                         // Team NPC barrier/base
	RoomCellTypeBas1                         // Team 1 base, spawnpoint
	RoomCellTypeBap1                         // Team 1 base, spawnpoint, arras base protector inhabits
	RoomCellTypeBad1                         // Team 1 base, spawnpoint, diep base protector inhabits
	RoomCellTypeBas2                         // Team 2 base, spawnpoint
	RoomCellTypeBap2                         // Team 2 base, spawnpoint, arras base protector inhabits
	RoomCellTypeBad2                         // Team 2 base, spawnpoint, diep base protector inhabits
	RoomCellTypeBas3                         // Team 3 base, spawnpoint
	RoomCellTypeBap3                         // Team 3 base, spawnpoint, arras base protector inhabits
	RoomCellTypeBad3                         // Team 3 base, spawnpoint, diep base protector inhabits
	RoomCellTypeBas4                         // Team 4 base, spawnpoint
	RoomCellTypeBap4                         // Team 4 base, spawnpoint, arras base protector inhabits
	RoomCellTypeBad4                         // Team 4 base, spawnpoint, diep base protector inhabits
	RoomCellTypeDom0                         // Dominator spawnpoint, NPC/contested team
	RoomCellTypeDom1                         // Dominator spawnpoint, Team 1
	RoomCellTypeDom2                         // Dominator spawnpoint, Team 2
	RoomCellTypeDom3                         // Dominator spawnpoint, Team 3
	RoomCellTypeDom4                         // Dominator spawnpoint, Team 4
	RoomCellTypePtl0                         // Portal zone, linked with other ptl0, team agnostic
	RoomCellTypePtl1                         // Portal zone, linked with other ptl1, team 1
	RoomCellTypePtl2                         // Portal zone, linked with other ptl2, team 2
	RoomCellTypePtl3                         // Portal zone, linked with other ptl3, team 3
	RoomCellTypePtl4                         // Portal zone, linked with other ptl4, team 4
	RoomCellTypeBarr                         // Barrier, used to separate sections of the map. Pushes people out of it as if it were the edge of the map. Stacks with itself to povide dense borders
	RoomCellTypeBoss                         // Boss spawnpoint, gamemode specific
	RoomCellTypeSENTINEL                     // NOT a real room type, invalid. Used as a sentinel for loops
)

var (
	// Please follow the naming convention inferred from here
	// See (package game) > (struct Room) > (func IndexType) for
	// any string maniupulations that may be done on the names
	RoomCellTypeNames map[RoomCellType]string = map[RoomCellType]string{
		RoomCellTypeNorm: "norm",
		RoomCellTypeNest: "nest",
		RoomCellTypeRock: "rock",
		RoomCellTypeRoid: "roid",
		RoomCellTypeWall: "wall",
		RoomCellTypeBas0: "bas0",
		RoomCellTypeBas1: "bas1",
		RoomCellTypeBap1: "bap1",
		RoomCellTypeBad1: "bad1",
		RoomCellTypeBas2: "bas2",
		RoomCellTypeBap2: "bap2",
		RoomCellTypeBad2: "bad2",
		RoomCellTypeBas3: "bas3",
		RoomCellTypeBap3: "bap3",
		RoomCellTypeBad3: "bad3",
		RoomCellTypeBas4: "bas4",
		RoomCellTypeBap4: "bap4",
		RoomCellTypeBad4: "bad4",
		RoomCellTypeDom0: "dom0",
		RoomCellTypeDom1: "dom1",
		RoomCellTypeDom2: "dom2",
		RoomCellTypeDom3: "dom3",
		RoomCellTypeDom4: "dom4",
		RoomCellTypePtl0: "ptl0",
		RoomCellTypePtl1: "ptl1",
		RoomCellTypePtl2: "ptl2",
		RoomCellTypePtl3: "ptl3",
		RoomCellTypePtl4: "ptl4",
		RoomCellTypeBarr: "barr",
		RoomCellTypeBoss: "boss",
	}

	RoomCellTypeIDs map[string]RoomCellType = map[string]RoomCellType{}
)

func init() {
	for id := range RoomCellTypeSENTINEL {
		if name, ok := RoomCellTypeNames[id]; ok {
			RoomCellTypeIDs[name] = id
		} else {
			panic("room cell type " + string(id) + " has no name")
		}
	}

	// Ensure that RoomCellTypeNames is unique and has no key outside of the range of (0, RoomCellTypeSENTINEL)
	for id := range RoomCellTypeNames {
		if id >= RoomCellTypeSENTINEL {
			panic("room cell type " + string(id) + " is outside of the range of (0, RoomCellTypeSENTINEL)")
		}
	}

	if len(RoomCellTypeNames) != len(RoomCellTypeIDs) {
		panic("room cell type names are not unique")
	}
}
