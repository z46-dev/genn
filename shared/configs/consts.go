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
	HitsOwnTypeNormal HitsOwnType = iota
	HitsOwnTypeNever
	HitsOwnTypeHard
	HitsOwnTypeHardWithBuffer
	HitsOwnTypeRepel
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
