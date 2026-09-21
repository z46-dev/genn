package game

import (
	"math"

	"github.com/z46-dev/genn/shared/configs"
)

const (
	SKILLS_LENGTH    int     = 10
	SKILL_LEVEL_CAP  int     = 45
	skillCurveLogDiv float64 = 1 / 1.60943791243410037460 // 1/log(5)
)

const (
	SkcnvMaxHealth int = iota
	SkcnvBodyDamage
	SkcnvBulletSpeed
	SkcnvBulletHealth
	SkcnvBulletPenetration
	SkcnvBulletDamage
	SkcnvReload
	SkcnvMovementSpeed
	SkcnvShieldRegeneration
	SkcnvShieldCapacity
)

// skillCurve applies the Arras diminishing-return curve to a raw skill value.
func skillCurve(raw, cap int) (res float64) {
	if cap > 0 {
		res = math.Log(4*float64(raw)/float64(cap)+1) * skillCurveLogDiv
	}

	return
}

// skillApply projects a curved skill value into a gameplay multiplier.
func skillApply(f, x float64) (res float64) {
	if x < 0 {
		res = 1 / (1 - x*f)
	} else {
		res = f*x + 1
	}

	return
}

// NewSkill creates a skill block with normal caps and freshly computed multipliers.
func NewSkill() (sk *Skill) {
	sk = &Skill{}
	sk.SetAllCaps(int(configs.SkillCapNormal))
	sk.Update()

	return
}

// Reset clears progression state and recomputes derived skill multipliers.
func (sk *Skill) Reset() {
	sk.Points = 0
	sk.Score = 0
	sk.Deduction = 0
	sk.Level = 0
	sk.CanUpgrade = false
	sk.updateLevelScore()
	sk.Update()
	sk.Maintain()
}

// Update clamps raw skill values and recomputes all derived multipliers.
func (sk *Skill) Update() {
	var attrib [SKILLS_LENGTH]float64

	for i := range SKILLS_LENGTH {
		if sk.Raw[i] > sk.Caps[i] {
			sk.Points += sk.Raw[i] - sk.Caps[i]
			sk.Raw[i] = sk.Caps[i]
		}
	}

	for i := range SKILLS_LENGTH {
		attrib[i] = skillCurve(sk.Raw[i], int(configs.SkillCapNormal))
	}

	sk.applyAttributes(attrib)
}

// applyAttributes stores gameplay multipliers from curved skill attributes.
func (sk *Skill) applyAttributes(attrib [SKILLS_LENGTH]float64) {
	sk.Reload = math.Pow(0.5, attrib[SkcnvReload])
	sk.BulletPenetration = skillApply(2.5, attrib[SkcnvBulletPenetration])
	sk.BulletHealth = skillApply(2, attrib[SkcnvBulletHealth])
	sk.BulletDamage = skillApply(3, attrib[SkcnvBulletDamage])
	sk.BulletSpeed = 0.5 + skillApply(1.5, attrib[SkcnvBulletSpeed])
	sk.BodyDamage = skillApply(1, attrib[SkcnvBodyDamage])
	sk.MaxHealth = skillApply(2, attrib[SkcnvMaxHealth])
	sk.MovementSpeed = skillApply(0.8, attrib[SkcnvMovementSpeed])
	sk.ShieldRegeneration = skillApply(25, attrib[SkcnvShieldRegeneration])
	sk.ShieldCapacity = skillApply(3, attrib[SkcnvShieldCapacity])
	sk.Ghost = attrib[SkcnvBulletPenetration]
}

// Maintain updates skill progression over time.
func (sk *Skill) Maintain() (updated bool) {
	if sk.Level < SKILL_LEVEL_CAP && sk.Score-sk.Deduction >= sk.LevelScore {
		sk.Deduction += sk.LevelScore
		sk.Level++

		switch sk.Level {
		case 37, 39, 41, 43:
			// No points
		default:
			sk.Points++
		}

		switch sk.Level {
		case 15, 30, 45:
			sk.CanUpgrade = true
		}

		sk.Update()
		updated = true
	}

	updated = false
	return
}

// SetAllCaps applies one cap to every skill.
func (sk *Skill) SetAllCaps(cap int) {
	for i := range SKILLS_LENGTH {
		sk.Caps[i] = cap
	}
}

func (sk *Skill) updateLevelScore() {
	var lvl float64 = float64(sk.Level)
	sk.LevelScore = math.Ceil(1.8*math.Pow(lvl+1, 1.8) - 2*lvl + 1)
}

func (sk *Skill) Upgrade(stat int) (updated bool) {
	if sk.Points > 0 && stat >= 0 && stat < SKILLS_LENGTH && sk.Raw[stat] < sk.Caps[stat] {
		sk.Raw[stat]++
		sk.Points--
		sk.Update()
		updated = true
	}

	return
}
