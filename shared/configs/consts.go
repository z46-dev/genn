package configs

const (
	TypeTank     Type = iota // Players, bosses, etc
	TypeBullet               // Anything spawned from a gun
	TypeFood                 // Polygons
	TypeCrasher              // Crashers, sentries, other guardian-like entities
	TypeObstacle             // Rocks or Maze Walls
)

const (
	GunStatKeyReload      GunStatKey = iota // Reload
	GunStatKeyRecoil                        // Recoil
	GunStatKeyShudder                       // Shudder
	GunStatKeySize                          // Size
	GunStatKeyHealth                        // Health
	GunStatKeyDamage                        // Damage
	GunStatKeyPenetration                   // Penetration
	GunStatKeySpeed                         // Speed
	GunStatKeyMaxSpeed                      // Max Speed
	GuNStatKeyRange                         // Range
	GunStatKeyDensity                       // Density
	GunStatKeySpray                         // Angular randomness
	GunStatKeyResist                        // Incoming damage resistance
	GunStatKey_SENTINEL                     // Sentinel, for representing the "number" of gun stat keys
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
	FacingTypeDefault FacingType = iota // Default, no behavior
	FacingTypeAutospin
	FacingTypeTurnWithSpeed
	FacingTypeTurnWithMotion
	FacingTypeTurnWithTarget
	FacingTypeSmoothWithMotion
	FacingTypeSmoothWithTarget
	FacingTypeBound
)

const (
	MotionTypeDefault MotionType = iota // Default, no behavior
	MotionTypeGlide
	MotionTypeMotor
	MotionTypeSwarm
	MotionTypeChase
	MotionTypeDrift
	MotionTypeBound
)
