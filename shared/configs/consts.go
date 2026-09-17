package configs

const (
	TypeTank     Type = iota // Players, bosses, etc
	TypeBullet               // Anything spawned from a gun
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
	UpgradeTier1 UpgradeTier = iota
	UpgradeTier2
	UpgradeTier3
	UpgradeTier4
	UpgradeTier_SENTINEL
)

var idCounter DefinitionID = 0
