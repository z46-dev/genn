package game

import "github.com/z46-dev/gamelib/vector"

// Control is a set of instructions for an entity to follow.
// To facilitate an easy "faucet passer" system, we will
// use the flags as a method to say "this value has been set"
// so we can avoid the issue of zero values without making
// non-faucet designs ugly to work with due to pointers

const (
	ControllerSetFlagsNone   ControllerSetFlags = 0
	ControllerSetFlagsTarget ControllerSetFlags = 1 << iota
	ControllerSetFlagsGoal
	ControllerSetFlagsMain
	ControllerSetFlagsAlt
	ControllerSetFlagsFire
	ControllerSetFlagsPower
)

func (c *Control) SetTarget(v *vector.Vec2[float64]) (self *Control) {
	c.Target = v
	c.SetFlags |= ControllerSetFlagsTarget

	self = c
	return
}

func (c *Control) SetGoal(v *vector.Vec2[float64]) (self *Control) {
	c.Goal = v
	c.SetFlags |= ControllerSetFlagsGoal

	self = c
	return
}

func (c *Control) SetMain(b bool) (self *Control) {
	c.Main = b
	c.SetFlags |= ControllerSetFlagsMain

	self = c
	return
}

func (c *Control) SetAlt(b bool) (self *Control) {
	c.Alt = b
	c.SetFlags |= ControllerSetFlagsAlt

	self = c
	return
}

func (c *Control) SetFire(b bool) (self *Control) {
	c.Fire = b
	c.SetFlags |= ControllerSetFlagsFire

	self = c
	return
}

func (c *Control) SetPower(f float64) (self *Control) {
	c.Power = f
	c.SetFlags |= ControllerSetFlagsPower

	self = c
	return
}
