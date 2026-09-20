package configs

func setUpgrades(target *Definition, tier UpgradeTier, upgrades ...*Definition) {
	if target.Upgrades[tier] == nil {
		target.Upgrades[tier] = []*Definition{}
	}

	target.Upgrades[tier] = append(target.Upgrades[tier], upgrades...)
}

func init() {
	// Food Upgrade Paths

	setUpgrades(Food, UpgradeTier1, Egg, GreenEgg, Gem)

	//// Standard Food Upgrade Path

	setUpgrades(Egg, UpgradeTier1, Square)
	setUpgrades(Square, UpgradeTier1, Triangle)
	setUpgrades(Triangle, UpgradeTier1, Pentagon)
	setUpgrades(Pentagon, UpgradeTier1, BetaPentagon)
	setUpgrades(BetaPentagon, UpgradeTier1, AlphaPentagon)

	//// Rare Food Upgrade Path

	setUpgrades(GreenEgg, UpgradeTier1, GreenSquare)
	setUpgrades(GreenSquare, UpgradeTier1, GreenTriangle)
	setUpgrades(GreenTriangle, UpgradeTier1, GreenPentagon)
	setUpgrades(GreenPentagon, UpgradeTier1, GreenBetaPentagon)
	setUpgrades(GreenBetaPentagon, UpgradeTier1, GreenAlphaPentagon)

	setUpgrades(Gem, UpgradeTier1, BetaGem)
	setUpgrades(BetaGem, UpgradeTier1, AlphaGem)

	// Tank Upgrade Paths

	//// Tier 0
	setUpgrades(Basic, UpgradeTier1, Twin, Sniper, MachineGun, FlankGuard, Director, Trapper)
}
