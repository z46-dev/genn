package configs

import "github.com/z46-dev/gamelib/vector"

// ==============================
// ======== Definitions =========
// ==============================

func New(name string) (b *DefBuilder) {
	b = &DefBuilder{
		Definition: &Definition{
			ID:       idCounter,
			Name:     &name,
			Upgrades: [UpgradeTier_SENTINEL][]*Definition{},
		},
	}

	idCounter++
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

func (b *BodyBuilder) FOV(f float64) (self *BodyBuilder) {
	b.Body.FOV = &f
	self = b
	return
}

func (b *BodyBuilder) Build() (out *BodyStats) {
	out = b.Body
	return
}
