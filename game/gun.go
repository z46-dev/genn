package game

import (
	"math"
	"math/rand/v2"

	"github.com/z46-dev/gamelib/vector"
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
			g.Fire()
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

func (g *Gun) Fire() {
	if !g.CanShoot {
		return
	}
}

func (g *Gun) SyncChildren() {}

func (g *Gun) Interpret() {
	var (
		sk                         *Skill                = g.Body.Skill
		shudder, spray             float64               = (rand.Float64() + rand.Float64() - 1) * g.Settings.Shudder * 2, (rand.Float64() + rand.Float64() - 1) * g.Settings.Spray / 2 * math.Pi / 180
		realSpeedScalar, realAngle float64               = runSpeed * g.Settings.Speed * sk.BulletSpeed * (1 + shudder), g.Body.Facing + g.Angle + spray
		speed                      *vector.Vec2[float64] = vector.NewVec2(math.Cos(realAngle)*realSpeedScalar, math.Sin(realAngle)*realSpeedScalar)
	)

	g.LastShot.Time = g.Body.Game.Time
	g.LastShot.Power = 3 * math.Log(math.Sqrt(sk.BulletSpeed)+g.TrueRecoil+1)
	g.AnimMotion += g.LastShot.Power

	// Apply boost if we should
	if speed.SquaredLength() != 0 && g.Body.Velocity.SquaredLength() != 0 {
		var (
			bullLen, bodyLen float64 = speed.Length(), g.Body.Velocity.Length()
			extraBoost       float64 = max(0, speed.X*g.Body.Velocity.X+speed.Y*g.Body.Velocity.Y) / bullLen / bodyLen
		)

		if extraBoost > 0 {
			speed.X += bodyLen * extraBoost * speed.X / bullLen
			speed.Y += bodyLen * extraBoost * speed.Y / bullLen
		}
	}

	var (
		offset, direction float64 = g.Offset.Length(), g.Offset.Direction()
		gpAngle, gAngle   float64 = g.Body.Facing + g.Angle + direction, g.Body.Facing + g.Angle
		sizeTuner         float64 = 1.5*g.Length - g.BaseWidth*g.Settings.Size/2
		gx, gy            float64 = offset*math.Cos(gpAngle) + sizeTuner*math.Cos(gAngle), offset*math.Sin(gpAngle) + sizeTuner*math.Sin(gAngle)
	)

	// Create the bullet
	/*var o *Entity =*/
	NewEntity(g.Body.Game, vector.NewVec2(g.Body.Position.X+g.Body.Size*gx-speed.X, g.Body.Position.Y+g.Body.Size*gy-speed.Y), g.Master.Master)

}
