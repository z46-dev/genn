package game

import (
	"math"

	"github.com/z46-dev/gamelib/poly"
	"github.com/z46-dev/gamelib/vector"
	"github.com/z46-dev/genn/shared"
	"github.com/z46-dev/genn/shared/configs"
)

const runSpeed float64 = 1.5

// Take the pointer of a value and return it.
func ptr[T any](v T) (pv *T) {
	pv = &v
	return
}

// Lazy ternary implementation. Returns a if cond is true, otherwise returns b.
func ternary[T any](cond bool, a, b T) (out T) {
	if cond {
		out = a
	} else {
		out = b
	}

	return
}

func eParent(e *Entity) (p *Parent) {
	p = &Parent{
		IsGun:  false,
		Entity: e,
		Gun:    nil,
	}

	return
}

func gParent(g *Gun) (p *Parent) {
	p = &Parent{
		IsGun:  true,
		Entity: nil,
		Gun:    g,
	}

	return
}

func NewEntity(g *Game, pos *vector.Vec2[float64], master *Entity) (e *Entity) {
	e = &Entity{
		Game:       g,
		ID:         g.EntitiesIDAccumulator,
		Position:   pos,
		Master:     master,
		MayCollide: true,
		Ghost:      false,
	}

	e.Source = e
	e.Parent = eParent(e)
	if master == nil {
		e.Master = e
	}

	g.EntitiesIDAccumulator++
	e.Game.Entities.Add(e)
	return
}

// Define sets the entity's properties based on a definition
func (e *Entity) Define(def *configs.Definition) {
	for _, parent := range def.Parents {
		e.Define(parent)
	}

	if def.Index != nil {
		e.Index = *def.Index
	}

	if def.Type != nil {
		e.Type = *def.Type
	}

	if def.Label != nil {
		e.Label = *def.Label
	}

	if def.Name != nil {
		e.Name = *def.Name
	}

	if def.Guns != nil {
		e.Guns = e.Guns[:0]
		for _, gunDef := range *def.Guns {
			e.Guns = append(e.Guns, NewGun(e, gunDef))
		}
	}

	if def.Turrets != nil {
		for _, o := range e.Turrets {
			o.Destroy()
		}

		e.Turrets = e.Turrets[:0]

		for _, turretDef := range *def.Turrets {
			var o *Entity = NewEntity(e.Game, e.Position.Copy(), e.Master)
			for _, turretDefDef := range turretDef.Definitions {
				o.Define(turretDefDef)
			}

			o.Bind(turretDef.Position, e)
		}
	}

	if def.Shape != nil {
		if def.Shape.Circle {
			e.Shape = nil
		} else {
			e.Shape = poly.NewPolygon(def.Shape.Points, e.Position, e.Size, e.Facing)
		}
	}

	if def.Color != nil {
		e.Color = *def.Color
	}

	if def.Controllers != nil {
		// TODO
	}

}

func (e *Entity) Bind(pos *configs.TurretPosition, master *Entity) {
	e.Bond = master
	e.Source = master
	e.Bond.Turrets = append(e.Bond.Turrets, e)
	e.Skill = e.Bond.Skill
	e.Label = e.Bond.Label + " " + e.Label
	e.MayCollide = false
	e.Ghost = true
	e.Bound = &EntityBound{
		Size:   pos.Size / 10,
		Angle:  pos.Angle * math.Pi / 180,
		Offset: vector.NewVec2(pos.Offset.X, pos.Offset.Y).Mul(0.1),
		Arc:    pos.Arc * math.Pi / 180,
	}

	e.Facing = e.Bond.Facing + e.Bound.Angle
	e.FacingType = configs.FacingTypeBound
	e.MotionType = configs.MotionTypeBound
	e.Move()
}

func (e *Entity) Move() {
	var (
		g, engine *vector.Vec2[float64] = e.Control.Goal.Subbed(e.Position), vector.Vec2_0[float64]()
		gActive   bool                  = g.X != 0 || g.Y != 0
		a         float64               = e.Acceleration
	)

	switch e.MotionType {
	case configs.MotionTypeGlide:
		e.MaxSpeed = e.TopSpeed
		e.Damp = 0.05
	case configs.MotionTypeMotor:
		e.MaxSpeed = 0
		if e.TopSpeed != 0 {
			e.Damp = a / e.TopSpeed
		}

		if gActive {
			var len float64 = g.Length()
			engine.X, engine.Y = a*g.X/len, a*g.Y/len
		}
	case configs.MotionTypeSwarm:
		e.MaxSpeed = e.TopSpeed
		var l float64 = engine.Dist(g) + 1.0
		if gActive && l > e.Size {
			var turning float64 = math.Sqrt((e.TopSpeed*max(1, e.Range) + 1) / a)
			engine = g.Mulled(e.TopSpeed / l).Sub(e.Velocity).Mul(1.0 / max(5, turning))
		} else if e.Velocity.Length() < e.TopSpeed {
			engine = e.Velocity.Mulled(a / 20)
		}
	case configs.MotionTypeChase:
		if gActive {
			var l float64 = engine.Dist(g) + 1.0
			if l > e.Size*2 {
				e.MaxSpeed = e.TopSpeed
				engine = g.Mulled(e.TopSpeed / l).Sub(e.Velocity).Mul(a)
			} else {
				e.MaxSpeed = 0
			}
		} else {
			e.MaxSpeed = 0
		}
	case configs.MotionTypeDrift:
		e.MaxSpeed = 0
		engine = g.Mulled(a)
	case configs.MotionTypeBound:
		var bLen, bDir float64 = e.Bound.Offset.Length(), e.Bound.Offset.Direction()
		e.Position.X = e.Bond.Position.X + e.Bond.Size*bLen*math.Cos(bDir+e.Bound.Angle+e.Bond.Facing)
		e.Position.Y = e.Bond.Position.Y + e.Bond.Size*bLen*math.Sin(bDir+e.Bound.Angle+e.Bond.Facing)
		e.Bond.Velocity.Add(e.DeltaV.Mulled(e.Bound.Size))
		e.FiringArc.Start, e.FiringArc.Width = e.Bond.Facing+e.Bound.Angle, e.Bound.Arc/2
		e.DeltaV.Mul(0)
	}

	e.DeltaV.Add(engine.Mul(e.Control.Power))
}

func (e *Entity) Face() {
	var oldFacing float64 = e.Facing

	switch e.FacingType {
	case configs.FacingTypeAutospin:
		e.Facing += 0.02
	case configs.FacingTypeTurnWithSpeed:
		e.Facing += e.Velocity.Length() / 90 * math.Pi
	case configs.FacingTypeTurnWithMotion:
		e.Facing = e.Velocity.Direction()
	case configs.FacingTypeSmoothWithMotion:
		e.Facing += shared.LoopSmooth(e.Facing, e.Velocity.Direction(), 4)
	case configs.FacingTypeTurnWithTarget:
		e.Facing = math.Atan2(e.Control.Target.Y, e.Control.Target.X)
	case configs.FacingTypeLocksFacing:
		if !e.Control.Alt {
			e.Facing = math.Atan2(e.Control.Target.Y, e.Control.Target.X)
		}
	case configs.FacingTypeSmoothWithTarget:
		e.Facing += shared.LoopSmooth(e.Facing, math.Atan2(e.Control.Target.Y, e.Control.Target.X), 4)
	case configs.FacingTypeBound:
		var givenAngle float64
		if e.Control.Main {
			givenAngle = math.Atan2(e.Control.Target.Y, e.Control.Target.X)
			var diff float64 = shared.AngleDifference(givenAngle, e.FiringArc.Start)
			if math.Abs(diff) >= e.FiringArc.Width {
				givenAngle = e.FiringArc.Start
			}
		} else {
			givenAngle = e.FiringArc.Start
		}

		e.Facing += shared.LoopSmooth(e.Facing, givenAngle, 2) // 4
	}

	e.vFacing = shared.AngleDifference(oldFacing, e.Facing)
}

func (e *Entity) Physics() {
	e.Velocity.Add(e.DeltaV)
	e.DeltaV.Mul(0)

	e.StepRemaining = 1
	e.Position.Add(e.Velocity) // could .Mul(e.StepRemaining) like arras did but it was always 1, so no need to waste
}

func (e *Entity) Friction() {
	var (
		motion float64 = e.Velocity.Length()
		excess float64 = motion - e.MaxSpeed
	)

	if excess > 0 && e.Damp != 0 {
		var finalVelocity = e.MaxSpeed + excess/(e.Damp+1)
		e.Velocity.X = finalVelocity * e.Velocity.X / motion
		e.Velocity.Y = finalVelocity * e.Velocity.Y / motion
	}
}

func (e *Entity) ConfinementToTheseEarthlyShackles() {
	var loc *Zone = e.Game.Room.Location(e.Position)

	if !e.Settings.CanGoOutsideRoom {
		e.DeltaV.X -= min(e.Position.X-e.Size+e.Game.Room.HalfWidth+50, 0) * e.Game.C.RoomBoundForce
		e.DeltaV.X -= max(e.Position.X+e.Size-e.Game.Room.HalfWidth-50, 0) * e.Game.C.RoomBoundForce
		e.DeltaV.Y -= min(e.Position.Y-e.Size+e.Game.Room.HalfHeight+50, 0) * e.Game.C.RoomBoundForce
		e.DeltaV.Y -= max(e.Position.Y+e.Size-e.Game.Room.HalfHeight-50, 0) * e.Game.C.RoomBoundForce

		if loc.Type == configs.RoomCellTypeBarr {
			// Push it out ONLY on the quad side it's on so we don't do some random BS
			if e.Position.X < loc.X1+e.Size {
				e.DeltaV.X += e.Game.C.RoomBoundForce
			} else if e.Position.X > loc.X2-e.Size {
				e.DeltaV.X -= e.Game.C.RoomBoundForce
			}

			if e.Position.Y < loc.Y1+e.Size {
				e.DeltaV.Y += e.Game.C.RoomBoundForce
			} else if e.Position.Y > loc.Y2-e.Size {
				e.DeltaV.Y -= e.Game.C.RoomBoundForce
			}
		}
	}

	if !e.Settings.MayGoInBase {
		if loc != nil && loc.KillZone && e.Team != loc.Team {
			e.Kill()
		}
	}
}

func (e *Entity) ContemplationOfMortality() {
	// TODO
}

func (e *Entity) Life() {
	// TODO
}

func (e *Entity) RefreshBodyAttributes() {
	// TODO
}

func (e *Entity) UpdateAABB() {
	var eX, eY float64 = e.Position.X + e.Velocity.X + e.DeltaV.X, e.Position.Y + e.Velocity.Y + e.DeltaV.Y
	e.AABB.X1 = min(e.Position.X, eX) - e.Size
	e.AABB.Y1 = min(e.Position.Y, eY) - e.Size
	e.AABB.X2 = max(e.Position.X, eX) + e.Size
	e.AABB.Y2 = max(e.Position.Y, eY) + e.Size
}

func (e *Entity) Kill() {
	e.Health.Amount = -1
}

func (e *Entity) Destroy() {
	e.Game.Entities.Remove(e.ID)

	if e.Parent != nil {
		if e.Parent.IsGun {
			delete(e.Parent.Gun.Children, e.ID)
		} else {
			delete(e.Parent.Entity.Children, e.ID)
		}
	}

	// TODO: Reconsider sources, masters, parents, children in relation to entities and whatnot
}
