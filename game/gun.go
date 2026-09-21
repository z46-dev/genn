package game

import (
	"math"

	"github.com/z46-dev/genn/shared/configs"
)

func NewGun() (g *Gun) {
	return
}
func (g *Gun) Update() {
	if !g.CanShoot {
		return
	}

	// Simulate visual recoil
	if g.AnimMotion != 0 || g.AnimPos != 0 {
		g.AnimMotion -= .25 * g.AnimPos
		g.AnimPos += g.AnimMotion

		// Bouncing off the back
		if g.AnimPos < 0 {
			g.AnimPos = 0
			g.AnimMotion *= -1
		}

		if g.AnimMotion > 0 {
			g.AnimMotion *= .75
		}
	}

	// Recoil force
	if g.AnimMotion > 0 {
		var force float64 = -g.AnimPos * g.Settings.Recoil * 0.045
		g.Body.DeltaV.X += math.Cos(g.Body.Facing+g.Angle) * force
		g.Body.DeltaV.Y += math.Sin(g.Body.Facing+g.Angle) * force
	}

	var (
		sk         *Skill = g.Body.Skill
		permission bool
	)

	// Decide what to do based on child-counting settings

	if g.MaxChildren > 0 || g.Body.MaxChildren > 0 {
		var (
			children float64 = float64(len(g.Body.Children))
			childCap int     = g.MaxChildren
		)

		if g.Calculator == configs.GunCalcNameNecro {
			children *= sk.Reload
		}

		if childCap == 0 {
			childCap = g.Body.MaxChildren
		}

		permission = childCap > int(math.Ceil(children))
	} else {
		permission = true
	}

	// Invuln people are not able to shoot
	if g.Body.Master.Invulnerable {
		permission = false
	}

	// Cycle up if we should
	if (permission || !g.WaitToCycle) && g.Cycle < 1 {
		var rldSk float64 = 1
		if g.Calculator != configs.GunCalcNameFixedReload && g.Calculator != configs.GunCalcNameNecro {
			rldSk = sk.Reload
		}

		g.Cycle += 1 / g.Settings.Reload / rldSk
	}

	// Determine if the player wants to shoot
	var wantsToShoot bool
	if g.AltFire {
		wantsToShoot = g.Body.Control.Alt
	} else {
		wantsToShoot = g.Body.Control.Main
	}

	// Firing Routines
	if permission && (g.Autofire || wantsToShoot) {
		if g.Cycle >= 1 {
			// var gx, gy float64 = 0, 0
			// g.Fire(gx, gy)
			g.Cycle--
		}
	} else {
		// If we're not shooting, only cycle up to where we'll have the proper firing delay
		var cycleDelay float64
		if g.WaitToCycle {
			cycleDelay = -g.Delay
		} else {
			cycleDelay = 1 - g.Delay
		}

		g.Cycle = min(g.Cycle, cycleDelay)
	}
}
