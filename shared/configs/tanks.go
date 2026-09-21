package configs

var GenericTank *Definition = New().
	Label("Generic Tank").
	Type(TypeTank).
	Danger(7).
	MotionType(MotionTypeMotor).
	FacingType(FacingTypeTurnWithTarget).
	Body(base).
	GiveKillMessage(true).
	DrawHealth(true).
	AcceptsScore(true).
	CanGoOutsideRoom(false).
	PersistsAfterDeath(false).
	CanBeOnLeaderboard(true).
	SkillCaps(SkillCaps{SkillCapNormal, SkillCapNormal, SkillCapNormal, SkillCapNormal, SkillCapNormal, SkillCapNormal, SkillCapNormal, SkillCapNormal, SkillCapNormal, SkillCapNormal}).
	DamageClass(DamageClassTanks).
	Build()

// Tier 0 (Danger = 4)

var Basic *Definition = New().
	Label("Basic").
	Parent(GenericTank).
	Danger(4).
	Gun(
		NewGun(gunArrasToGenn(18, 8, 1, 0, 0, 0, 0)).
			Shoots(
				combineStats(gBasic),
				Bullet,
			).
			Build(),
	).
	Build()

// Tier 1 (Danger = 5)

var Twin *Definition = New().
	Label("Twin").
	Parent(GenericTank).
	Danger(5).
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
	Danger(5).
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
	Danger(5).
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
	Danger(5).
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

var Director *Definition = New().
	Parent(GenericTank).
	Label("Director").
	Danger(5).
	StatNames(StatNamesDrone).
	Body(
		NewBody().
			Acceleration(*base.Acceleration * 0.75).
			FOV(*base.FOV * 1.1).
			Build(),
	).
	MaxChildren(5).
	Gun(
		NewGun(gunArrasToGenn(6, 12, 1.2, 8, 0, 0, 0)).
			Shoots(
				combineStats(gDrone),
				Drone,
			).
			Autofire(true).
			SyncSkills(true).
			StatCalculator(GunCalcNameDrone).
			Build(),
	).
	Build()

var Trapper *Definition = New().
	Parent(GenericTank).
	Label("Trapper").
	Danger(5).
	StatNames(StatNamesTrap).
	Body(
		NewBody().
			Speed(*base.Speed*0.9).
			Build(),
	).
	Gun(
		NewGun(gunArrasToGenn(15, 7, 1, 0, 0, 0, 0)).Build(),
		NewGun(gunArrasToGenn(3, 7, 1.6, 15, 0, 0, 0)).
			Shoots(
				combineStats(gTrap),
				Trap,
			).
			StatCalculator(GunCalcNameTrap).
			Build(),
	).
	Build()

var END bool = true

// Tier 2 (Danger = 6)

// Tier 3 (Danger = 7)
