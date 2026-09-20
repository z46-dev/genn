package configs

func combineStats(parts ...GunStats) (out GunStats) {
	out = GunStats{
		Reload: 1, Recoil: 1, Shudder: 1, Size: 1,
		Health: 1, Damage: 1, Penetration: 1,
		Speed: 1, MaxSpeed: 1, Range: 1,
		Density: 1, Spray: 1, Resist: 1,
	}

	for _, p := range parts {
		out.Reload *= p.Reload
		out.Recoil *= p.Recoil
		out.Shudder *= p.Shudder
		out.Size *= p.Size
		out.Health *= p.Health
		out.Damage *= p.Damage
		out.Penetration *= p.Penetration
		out.Speed *= p.Speed
		out.MaxSpeed *= p.MaxSpeed
		out.Range *= p.Range
		out.Density *= p.Density
		out.Spray *= p.Spray
		out.Resist *= p.Resist
	}

	return out
}

var (
	// Core stats
	gBasic = GunStats{Reload: 18, Recoil: 1.4, Shudder: 0.1, Size: 1, Health: 1, Damage: 0.75, Penetration: 1, Speed: 4.5, MaxSpeed: 1, Range: 1, Density: 1, Spray: 15, Resist: 1}
	gDrone = GunStats{Reload: 50, Recoil: 0.25, Shudder: 0.1, Size: 0.6, Health: 1, Damage: 1, Penetration: 1, Speed: 2, MaxSpeed: 1, Range: 1, Density: 1, Spray: 0.1, Resist: 1}
	gTrap  = GunStats{Reload: 36, Recoil: 1, Shudder: 0.25, Size: 0.6, Health: 1, Damage: 0.75, Penetration: 1, Speed: 5, MaxSpeed: 1, Range: 1, Density: 1, Spray: 15, Resist: 3}

	// Tier 1 Stats
	gTwin   = GunStats{Reload: 1, Recoil: 0.5, Shudder: 0.9, Size: 1, Health: 0.9, Damage: 0.7, Penetration: 1, Speed: 1, MaxSpeed: 1, Range: 1, Density: 1, Spray: 1.2, Resist: 1}
	gSniper = GunStats{Reload: 1.35, Recoil: 1, Shudder: 0.25, Size: 1, Health: 1, Damage: 0.8, Penetration: 1.1, Speed: 1.5, MaxSpeed: 1.5, Range: 1, Density: 1.5, Spray: 0.2, Resist: 1.15}
	gMach   = GunStats{Reload: 0.5, Recoil: 0.8, Shudder: 1.7, Size: 1, Health: 0.7, Damage: 0.7, Penetration: 1, Speed: 1, MaxSpeed: 0.8, Range: 1, Density: 1, Spray: 2.5, Resist: 1}
	gFlank  = GunStats{Reload: 1, Recoil: 1.2, Shudder: 1, Size: 1, Health: 1.02, Damage: 0.81, Penetration: 0.9, Speed: 1, MaxSpeed: 0.85, Range: 1, Density: 1.2, Spray: 1, Resist: 1}

	// Tier 2 Stats

	// Tier 3 Stats

	// NPC Stats

	// Balance Modifiers
	gAuto = GunStats{Reload: 1.15, Recoil: 0.1, Shudder: 1, Size: 1, Health: 0.7, Damage: 0.8, Penetration: 1.25, Speed: 1.05, MaxSpeed: 1.05, Range: 0.98, Density: 1, Spray: 1.1, Resist: 1}
)

var base *BodyStats = NewBody().
	Health(20).
	Damage(3).
	Penetration(1.05).
	Shield(8).
	Regeneration(0.025).
	Heterogeneity(3).
	Density(0.5).
	Pushability(0.9).
	Resist(1).
	Acceleration(1.6).
	Speed(5.25).
	FOV(1).
	Build()
