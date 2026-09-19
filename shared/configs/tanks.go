package configs

var GenericTank *Definition = New().
	Label("Generic Tank").
	Type(TypeTank).
	Danger(5).
	MotionType(MotionTypeMotor).
	FacingType(FacingTypeTurnWithTarget).
	Body(base).
	GiveKillMessage(true).
	DrawHealth(true).
	AcceptsScore(true).
	CanGoOutsideRoom(false).
	PersistsAfterDeath(false).
	CanBeOnLeaderboard(true).
	Build()

// Tier 0

var Basic *Definition = New().
	Label("Basic").
	Parent(GenericTank).
	Gun(
		NewGun(gunArrasToGenn(18, 8, 1, 0, 0, 0, 0)).
			Shoots(
				combineStats(gBasic),
				Bullet,
			).
			Build(),
	).
	Build()

// Tier 1

var Twin *Definition = New().
	Label("Twin").
	Parent(GenericTank).
	Gun(
		NewGun(gunArrasToGenn(20, 8, 1, 0, 5.5, 0, 0)).
			Shoots(
				combineStats(gBasic, gTwin),
				Bullet,
			).
			Build(),
		NewGun(gunArrasToGenn(20, 8, 1, 0, -5.5, 0, 0.5)).
			Shoots(
				combineStats(gBasic, gTwin),
				Bullet,
			).
			Build(),
	).
	Build()

var Sniper *Definition = New().
	Label("Sniper").
	Parent(GenericTank).
	Body(NewBody().Acceleration(*base.Acceleration * 0.7).FOV(*base.FOV * 1.2).Build()).
	Gun(
		NewGun(gunArrasToGenn(24, 8.5, 1, 0, 0, 0, 0)).
			Shoots(
				combineStats(gBasic, gSniper),
				Bullet,
			).
			Build(),
	).
	Build()

var MachineGun *Definition = New().
	Label("Machine Gun").
	Parent(GenericTank).
	Gun(
		NewGun(gunArrasToGenn(12, 10, 1.4, 8, 0, 0, 0)).
			Shoots(
				combineStats(gBasic, gMach),
				Bullet,
			).
			Build(),
	).
	Build()

var FlankGuard *Definition = New().
	Label("Flank Guard").
	Parent(GenericTank).
	Body(NewBody().FOV(*base.Speed * 1.1).Build()).
	Gun((func() (out []*Gun) {
		for i := range 3 {
			out = append(
				out,
				NewGun(gunArrasToGenn(18, 8, 1, 0, 0, 120*float64(i), 0)).
					Shoots(
						combineStats(gBasic, gFlank),
						Bullet,
					).
					Build(),
			)
		}

		return
	})()...).
	Build()

var END bool = true

// Tier 2

// Tier 3
