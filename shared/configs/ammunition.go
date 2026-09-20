package configs

var Bullet *Definition = New().
	Label("Bullet").
	Type(TypeBullet).
	Body(
		NewBody().
			Health(0.33 * WeaponHealthFactor).
			Damage(4 * WeaponDamageFactor).
			Penetration(1).
			Speed(3.75).
			Range(90).
			Density(1.25).
			Pushability(0.3).
			Build(),
	).
	FacingType(FacingTypeSmoothWithMotion).
	CanGoOutsideRoom(true).
	HitsOwnType(HitsOwnTypeNever).
	DieAtRange(true).
	Build()

var Drone *Definition = New().
	Label("Drone").
	Type(TypeDrone).
	Shape(GenericShape(3)).
	MotionType(MotionTypeChase).
	FacingType(FacingTypeSmoothWithTarget).
	Controllers(ControllerNearestDifferentMaster, ControllerCanRepel, ControllerMapTargetToGoal, ControllerHangOutNearMaster).
	AI(NewAI().OwnerVisionBound(true).Build()).
	Body(
		NewBody().
			Penetration(1.2).
			Pushability(0.6).
			Acceleration(0.05).
			Health(0.6 * WeaponHealthFactor).
			Damage(1.25 * WeaponDamageFactor).
			Speed(3.8).
			Range(200).
			Density(0.03).
			Resist(1.5).
			FOV(0.8).
			Build(),
	).
	HitsOwnType(HitsOwnTypeHard).
	ClearOnMasterUpgrade(true).
	Build()

var Trap *Definition = New().
	Label("Thrown Trap").
	Type(TypeTrap).
	Shape(GenericShape(-3)).
	MotionType(MotionTypeGlide).
	FacingType(FacingTypeTurnWithSpeed).
	HitsOwnType(HitsOwnTypePush).
	DieAtRange(true).
	Body(
		NewBody().
			Health(1 * WeaponHealthFactor).
			Damage(2 * WeaponDamageFactor).
			Range(450).
			Density(2.5).
			Resist(2.5).
			Speed(0).
			Build(),
	).
	Build()
