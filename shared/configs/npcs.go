package configs

// Walls

var Obstacle *Definition = New().
	Label("Rock").
	Type(TypeObstacle).
	FacingType(FacingTypeTurnWithSpeed).
	Shape(GenericShape(-9)).
	Body(
		NewBody().
			Health(10000).
			Shield(10000).
			Damage(1).
			Resist(100).
			Pushability(0).
			Build(),
	).
	Size(60).
	Color(16).
	VariesInSize(true).
	Build()

var BabyObstacle *Definition = New().
	Label("Gravel").
	Parent(Obstacle).
	Size(25).
	Shape(GenericShape(-7)).
	Build()

// Crashers

var Crasher *Definition = New().
	Label("Crasher").
	Type(TypeCrasher).
	Color(5).
	Shape(GenericShape(3)).
	Size(5).
	Value(10).
	VariesInSize(true).
	AI(
		NewAI().
			DirectAim(true).
			Build(),
	).
	Body(
		NewBody().
			Health(0.5).
			Damage(5).
			Speed(5).
			Acceleration(0.01).
			Penetration(2).
			Pushability(0.5).
			Density(10).
			Resist(2).
			Build(),
	).
	MotionType(MotionTypeMotor).
	FacingType(FacingTypeSmoothWithMotion).
	HitsOwnType(HitsOwnTypeHard).
	DrawHealth(true).
	Build()

// Sentries

var Sentry *Definition = New().
	Label("Sentry").
	Parent(Crasher).
	Danger(3).
	Size(10).
	Value(1500).
	FacingType(FacingTypeSmoothWithTarget).
	GiveKillMessage(true).
	Body(
		NewBody().
			Health(*base.Health * 1.25).
			Damage(*base.Damage * 2).
			Speed(*base.Speed * 0.5).
			FOV(0.5).
			Build(),
	).
	Build()

var SentryGun *Definition = makeAuto(Sentry, NewMakeAutoOptions().Name("Sentry"))
