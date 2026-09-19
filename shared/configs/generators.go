package configs

import (
	"math"

	"github.com/z46-dev/gamelib/vector"
)

// ==============================
// ======== Definitions =========
// ==============================

func New() (b *DefBuilder) {
	var nowID DefinitionID = idCounter

	b = &DefBuilder{
		Definition: &Definition{
			Index:    &nowID,
			Upgrades: [UpgradeTier_SENTINEL][]*Definition{},
		},
	}

	idCounter++
	return
}

func NewDummy() (b *DefBuilder) {
	b = &DefBuilder{
		Definition: &Definition{
			Upgrades: [UpgradeTier_SENTINEL][]*Definition{},
		},
	}

	return
}

func (b *DefBuilder) Label(label string) (self *DefBuilder) {
	b.Definition.Label = &label
	self = b
	return
}

func (b *DefBuilder) Name(name string) (self *DefBuilder) {
	b.Definition.Name = &name
	self = b
	return
}

func (b *DefBuilder) Type(t Type) (self *DefBuilder) {
	b.Definition.Type = &t
	self = b
	return
}

func (b *DefBuilder) Parent(p ...*Definition) (self *DefBuilder) {
	b.Definition.Parents = append(b.Definition.Parents, p...)
	self = b
	return
}

func (b *DefBuilder) Gun(g ...*Gun) (self *DefBuilder) {
	b.Definition.Guns = append(b.Definition.Guns, g...)
	self = b
	return
}

func (b *DefBuilder) Turret(t ...*Turret) (self *DefBuilder) {
	b.Definition.Turrets = append(b.Definition.Turrets, t...)
	self = b
	return
}

func (b *DefBuilder) Upgrade(t UpgradeTier, d ...*Definition) (self *DefBuilder) {
	b.Definition.Upgrades[t] = append(b.Definition.Upgrades[t], d...)
	self = b
	return
}

func (b *DefBuilder) Body(body *BodyStats) (self *DefBuilder) {
	b.Definition.Body = body
	self = b
	return
}

func (b *DefBuilder) AI(ai *AISettings) (self *DefBuilder) {
	b.Definition.AI = ai
	self = b
	return
}

func (b *DefBuilder) MotionType(m MotionType) (self *DefBuilder) {
	b.Definition.MotionType = &m
	self = b
	return
}

func (b *DefBuilder) FacingType(f FacingType) (self *DefBuilder) {
	b.Definition.FacingType = &f
	self = b
	return
}

func (b *DefBuilder) Danger(d float64) (self *DefBuilder) {
	b.Definition.Danger = &d
	self = b
	return
}

func (b *DefBuilder) GiveKillMessage(k bool) (self *DefBuilder) {
	b.Definition.GiveKillMessage = &k
	self = b
	return
}

func (b *DefBuilder) DrawHealth(d bool) (self *DefBuilder) {
	b.Definition.DrawHealth = &d
	self = b
	return
}

func (b *DefBuilder) AcceptsScore(a bool) (self *DefBuilder) {
	b.Definition.AcceptsScore = &a
	self = b
	return
}

func (b *DefBuilder) CanGoOutsideRoom(c bool) (self *DefBuilder) {
	b.Definition.CanGoOutsideRoom = &c
	self = b
	return
}

func (b *DefBuilder) PersistsAfterDeath(p bool) (self *DefBuilder) {
	b.Definition.PersistsAfterDeath = &p
	self = b
	return
}

func (b *DefBuilder) CanBeOnLeaderboard(c bool) (self *DefBuilder) {
	b.Definition.CanBeOnLeaderboard = &c
	self = b
	return
}

func (b *DefBuilder) SetValue(v float64) (self *DefBuilder) {
	b.Definition.Value = &v
	self = b
	return
}

func (b *DefBuilder) SetMaxChildren(m int) (self *DefBuilder) {
	b.Definition.MaxChildren = &m
	self = b
	return
}

func (b *DefBuilder) HitsOwnType(h HitsOwnType) (self *DefBuilder) {
	b.Definition.HitsOwnType = &h
	self = b
	return
}

func (b *DefBuilder) Shape(s Shape) (self *DefBuilder) {
	b.Definition.Shape = &s
	self = b
	return
}

func (b *DefBuilder) DieAtRange(d bool) (self *DefBuilder) {
	b.Definition.DieAtRange = &d
	self = b
	return
}

func (b *DefBuilder) HealthWithLevel(h bool) (self *DefBuilder) {
	b.Definition.HealthWithLevel = &h
	self = b
	return
}

func (b *DefBuilder) VariesInSize(v bool) (self *DefBuilder) {
	b.Definition.VariesInSize = &v
	self = b
	return
}

func (b *DefBuilder) AdvancedDamage(a bool) (self *DefBuilder) {
	b.Definition.AdvancedDamage = &a
	self = b
	return
}

func (b *DefBuilder) RatioEffects(r bool) (self *DefBuilder) {
	b.Definition.RatioEffects = &r
	self = b
	return
}

func (b *DefBuilder) Color(c int) (self *DefBuilder) {
	b.Definition.Color = &c
	self = b
	return
}

func (b *DefBuilder) Size(s float64) (self *DefBuilder) {
	b.Definition.Size = &s
	self = b
	return
}

func (b *DefBuilder) Value(v float64) (self *DefBuilder) {
	b.Definition.Value = &v
	self = b
	return
}

func (b *DefBuilder) Independent(i bool) (self *DefBuilder) {
	b.Definition.Independent = &i
	self = b
	return
}

func (b *DefBuilder) Build() (out *Definition) {
	out = b.Definition
	return
}

// ==============================
// ========= Gun  Stats =========
// ==============================

func NewGun(length, baseWidth, endWidth, x, y, angle, delay float64) (b *GunBuilder) {
	b = &GunBuilder{
		Gun: &Gun{
			Position: &GunPosition{
				Length:    length,
				BaseWidth: baseWidth,
				EndWidth:  endWidth,
				Offset:    *vector.NewVec2(x, y),
				Angle:     angle,
				Delay:     delay,
			},
			Properties: &GunProperties{},
		},
	}

	return
}

func (b *GunBuilder) Position(p *GunPosition) (self *GunBuilder) {
	b.Gun.Position = p
	self = b
	return
}

func (b *GunBuilder) Shoots(stats GunStats, shoots *Definition) (self *GunBuilder) {
	b.Gun.Properties.ShootSettings = stats
	b.Gun.Properties.Shoots = shoots
	self = b
	return
}

func (b *GunBuilder) StatCalculator(calc GunCalcName) (self *GunBuilder) {
	b.Gun.Properties.StatCalculator = &calc
	self = b
	return
}

func (b *GunBuilder) Build() (out *Gun) {
	out = b.Gun
	return
}

// ==============================
// ======== Turret Stats ========
// ==============================

func NewTurret(size, x, y, angle, arc float64, layer int) (b *TurretBuilder) {
	b = &TurretBuilder{
		Turret: &Turret{
			Position: &TurretPosition{
				Size:   size,
				Offset: *vector.NewVec2(x, y),
				Angle:  angle,
				Arc:    arc,
				Layer:  layer,
			},
		},
	}

	return
}

func (b *TurretBuilder) Position(p *TurretPosition) (self *TurretBuilder) {
	b.Turret.Position = p
	self = b
	return
}

func (b *TurretBuilder) Definitions(d ...*Definition) (self *TurretBuilder) {
	b.Turret.Definitions = append(b.Turret.Definitions, d...)
	self = b
	return
}

func (b *TurretBuilder) Build() (out *Turret) {
	out = b.Turret
	return
}

// ==============================
// ========= Body Stats =========
// ==============================

func NewBody() (b *BodyBuilder) {
	b = &BodyBuilder{
		Body: &BodyStats{},
	}

	return
}

func (b *BodyBuilder) Health(h float64) (self *BodyBuilder) {
	b.Body.Health = &h
	self = b
	return
}

func (b *BodyBuilder) Damage(d float64) (self *BodyBuilder) {
	b.Body.Damage = &d
	self = b
	return
}

func (b *BodyBuilder) Penetration(p float64) (self *BodyBuilder) {
	b.Body.Penetration = &p
	self = b
	return
}

func (b *BodyBuilder) Shield(s float64) (self *BodyBuilder) {
	b.Body.Shield = &s
	self = b
	return
}

func (b *BodyBuilder) Regeneration(r float64) (self *BodyBuilder) {
	b.Body.Regeneration = &r
	self = b
	return
}

func (b *BodyBuilder) Heterogeneity(h float64) (self *BodyBuilder) {
	b.Body.Heterogeneity = &h
	self = b
	return
}

func (b *BodyBuilder) Density(d float64) (self *BodyBuilder) {
	b.Body.Density = &d
	self = b
	return
}

func (b *BodyBuilder) Pushability(p float64) (self *BodyBuilder) {
	b.Body.Pushability = &p
	self = b
	return
}

func (b *BodyBuilder) Resist(r float64) (self *BodyBuilder) {
	b.Body.Resist = &r
	self = b
	return
}

func (b *BodyBuilder) Acceleration(a float64) (self *BodyBuilder) {
	b.Body.Acceleration = &a
	self = b
	return
}

func (b *BodyBuilder) Speed(s float64) (self *BodyBuilder) {
	b.Body.Speed = &s
	self = b
	return
}

func (b *BodyBuilder) Range(r float64) (self *BodyBuilder) {
	b.Body.Range = &r
	self = b
	return
}

func (b *BodyBuilder) FOV(f float64) (self *BodyBuilder) {
	b.Body.FOV = &f
	self = b
	return
}

func (b *BodyBuilder) Build() (out *BodyStats) {
	out = b.Body
	return
}

// ==============================
// =========== Shapes ===========
// ==============================

// GenericShape generates a generic shape with the given number of sides.
// If sides is 0, it will generate a circle. If sides is negative, it will
// generate a star with the given number of points.
func GenericShape(sides int) (out Shape) {
	out.Circle = sides == 0

	if !out.Circle {
		var (
			numPoints int  = sides
			star      bool = sides < 0
		)

		if star {
			numPoints = -sides * 2
		}

		out.Points = make([]*vector.Vec2[float64], numPoints)

		for i := range numPoints {
			var angle, radius float64 = float64(i) * (2 * math.Pi / float64(numPoints)), 1.0

			if star && i%2 == 1 {
				radius = 0.5
			}

			out.Points[i] = vector.NewVec2(radius*math.Cos(angle), radius*math.Sin(angle))
		}
	}

	return
}

// ==============================
// ======== AI  Settings ========
// ==============================

func NewAI() (b *AIBuilder) {
	b = &AIBuilder{
		AI: &AISettings{},
	}

	return
}

func (b *AIBuilder) DirectAim(d bool) (self *AIBuilder) {
	b.AI.DirectAim = &d
	self = b
	return
}

func (b *AIBuilder) IgnoreGunRange(i bool) (self *AIBuilder) {
	b.AI.IgnoreGunRange = &i
	self = b
	return
}

func (b *AIBuilder) OwnerVisionBound(o bool) (self *AIBuilder) {
	b.AI.OwnerVisionBound = &o
	self = b
	return
}

func (b *AIBuilder) TargetAllDangerLevels(t bool) (self *AIBuilder) {
	b.AI.TargetAllDangerLevels = &t
	self = b
	return
}

func (b *AIBuilder) FullView(f bool) (self *AIBuilder) {
	b.AI.FullView = &f
	self = b
	return
}

func (b *AIBuilder) IgnoreFood(i bool) (self *AIBuilder) {
	b.AI.IgnoreFood = &i
	self = b
	return
}

func (b *AIBuilder) RandomizeStrafe(r bool) (self *AIBuilder) {
	b.AI.RandomizeStrafe = &r
	self = b
	return
}

func (b *AIBuilder) Build() (out *AISettings) {
	out = b.AI
	return
}
