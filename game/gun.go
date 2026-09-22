package game

import (
	"fmt"
	"math"
	"math/rand/v2"

	"github.com/z46-dev/gamelib/vector"
	"github.com/z46-dev/genn/shared/configs"
)

func NewGun(body *Entity, info *configs.Gun) (g *Gun) {
	g = &Gun{
		Body:     body,
		Master:   body.Source,
		Children: make(map[uint64]*Entity),
		CanShoot: info.Properties != nil && len(info.Properties.Shoots) > 0,
	}

	if g.CanShoot {
		g.Settings = info.Properties.ShootSettings
		g.BulletTypes = info.Properties.Shoots
		g.Autofire = info.Properties.Autofire
		g.AltFire = info.Properties.AltFire
		g.Calculator = info.Properties.StatCalculator
		g.WaitToCycle = info.Properties.WaitToCycle
		g.MaxChildren = info.Properties.MaxChildren
		g.SyncSkills = info.Properties.SyncSkills

		var natural *configs.BodyBuilder = configs.NewBody()
		for _, def := range g.BulletTypes {
			if def.Body != nil {
				natural.OverwriteWith(def.Body)
			}
		}

		g.Natural = natural.Build()
	}

	g.Length = info.Position.Length / 10
	g.BaseWidth = info.Position.BaseWidth / 10
	g.EndWidth = info.Position.EndWidth / 10
	g.Offset = vector.NewVec2(info.Position.Offset.X/10, info.Position.Offset.Y/10)
	g.Angle = info.Position.Angle * math.Pi / 180
	g.Delay = info.Position.Delay

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
	var o *Entity = NewEntity(g.Body.Game, vector.NewVec2(g.Body.Position.X+g.Body.Size*gx-speed.X, g.Body.Position.Y+g.Body.Size*gy-speed.Y), g.Master.Master)
	o.Velocity.X = speed.X
	o.Velocity.Y = speed.Y
	g.BulletInit(o)
}

func (g *Gun) BulletInit(o *Entity) {
	// Define it by its natural properties
	for _, def := range g.BulletTypes {
		o.Define(def)
	}

	// Pass the gun attributes
	o.Define(&configs.Definition{
		Body:   g.Interpret(),
		Skills: g.GetSkillRaw(),
		Size:   ptr(g.Body.Size * g.BaseWidth * g.Settings.Size / 2),
		Label:  ptr(fmt.Sprintf("%s %s", g.Master.Label, o.Label)),
	})

	o.Color = g.Body.Master.Color

	if g.MaxChildren > 0 {
		o.Parent = gParent(g)
		g.Children[o.ID] = o
	} else if g.Body.MaxChildren > 0 {
		o.Parent = eParent(g.Body)
		g.Body.Children[o.ID] = o
	}

	o.Source = g.Body
	o.Facing = o.Velocity.Direction()

	o.RefreshBodyAttributes()
	o.Life()
}

func (g *Gun) SyncChildren() {
	if !g.SyncSkills {
		return
	}

	var (
		interpret *configs.BodyStats = g.Interpret()
		skillRaw  *configs.Skills    = g.GetSkillRaw()
	)

	for _, child := range g.Children {
		child.Define(&configs.Definition{
			Body:   interpret,
			Skills: skillRaw,
		})
	}
}

func (g *Gun) Interpret() (b *configs.BodyStats) {
	var (
		body       *configs.BodyBuilder = configs.NewBody()
		sizeFactor float64              = 1
		sk         *Skill               = g.Body.Skill
	)

	body.
		Speed(g.Settings.MaxSpeed * sk.BulletSpeed).
		Health(g.Settings.Health * sk.BulletHealth).
		Resist(g.Settings.Resist * sk.BulletResist).
		Damage(g.Settings.Damage * sk.BulletDamage).
		Penetration(max(1, g.Settings.Penetration*sk.BulletPenetration)).
		Range(g.Settings.Range / math.Sqrt(sk.BulletSpeed)).
		Density(g.Settings.Density * sk.BulletPenetration * sk.BulletPenetration / sizeFactor).
		Pushability(1 / sk.BulletPenetration).
		Heterogeneity(3 - 2.8*sk.Ghost)

	// Special calculations for certain gun types
	switch g.Calculator {
	case configs.GunCalcNameThruster:
		g.TrueRecoil = g.Settings.Recoil * math.Sqrt(sk.Reload*sk.BulletSpeed)
	case configs.GunCalcNameSustained:
		body.Range(g.Settings.Range)
	case configs.GunCalcNameSwarm:
		body.Penetration(max(1, g.Settings.Penetration*(0.5*(sk.BulletPenetration-1)+1)))
		body.Health(*body.Body.Health / math.Pow(sk.BulletPenetration, 0.5))
	case configs.GunCalcNameTrap:
		body.Pushability(1 / math.Pow(sk.BulletPenetration, 0.5))
		body.Range(g.Settings.Range)
	case configs.GunCalcNameDrone, configs.GunCalcNameNecro:
		body.Pushability(1)
		body.Penetration(max(1, g.Settings.Penetration*(0.5*(sk.BulletPenetration-1)+1)))
		body.Health((g.Settings.Health*sk.BulletHealth + sizeFactor) / math.Pow(sk.BulletPenetration, 0.8))
		body.Damage(g.Settings.Damage * sk.BulletDamage * math.Sqrt(sizeFactor) * g.Settings.Penetration * sk.BulletPenetration)
		body.Range(g.Settings.Range * math.Sqrt(sizeFactor))
	}

	// Respect the natural properties of the bullet
	body.MultiplyInto(g.Natural)
	return
}

func (g *Gun) GetSkillRaw() (s *configs.Skills) {
	s = &configs.Skills{
		BulletSpeed:       g.Body.Skill.Raw[SkcnvBulletSpeed],
		BulletHealth:      g.Body.Skill.Raw[SkcnvBulletHealth],
		BulletPenetration: g.Body.Skill.Raw[SkcnvBulletPenetration],
		BulletDamage:      g.Body.Skill.Raw[SkcnvBulletDamage],
		Reload:            g.Body.Skill.Raw[SkcnvReload],
	}

	return
}
