package configs

var AutoTurret *Definition = New().
	Label("Turret").
	Parent(GenericTank).
	Gun(
		NewGun(gunArrasToGenn(22, 10, 1, 0, 0, 0, 0)).
			Shoots(
				combineStats(gBasic, gAuto),
				Bullet,
			).
			Build(),
	).
	Build()
