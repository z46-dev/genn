package configs

var (
	GenericTank *Definition = New("Generic Tank").
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
		CanBeOnLeaderboard(false).
		SetValue(0).
		SetMaxChildren(0).
		Build()
)
