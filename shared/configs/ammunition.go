package configs

var (
	Bullet *Definition = New().
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
)
