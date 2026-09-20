package configs

import "math"

// Base Food

var Food *Definition = New().
	Type(TypeFood).
	HitsOwnType(HitsOwnTypeRepel).
	MotionType(MotionTypeDrift).
	FacingType(FacingTypeTurnWithSpeed).
	VariesInSize(true).
	Body(NewBody().Pushability(1).Build()).
	AdvancedDamage(false).
	RatioEffects(true).
	HealthWithLevel(false).
	DrawHealth(true).
	NPCControllers(ControllerMoveInCircles).
	Build()

var Egg *Definition = New().
	Label("Egg").
	Parent(Food).
	Body(NewBody().Health(0.01).Damage(0).Density(2).Pushability(0).Build()).
	DrawHealth(false).
	Size(5).
	Color(6).
	Value(10).
	Shape(GenericShape(0)).
	Build()

var Square *Definition = New().
	Label("Square").
	Parent(Food).
	Body(NewBody().Health(BaseFoodHealth).Damage(BaseFoodDamage).Density(4).Penetration(2).Build()).
	Size(10).
	Color(13).
	Value(30).
	Shape(GenericShape(4)).
	Build()

var Triangle *Definition = New().
	Label("Triangle").
	Parent(Food).
	Body(NewBody().Health(3 * BaseFoodHealth).Damage(BaseFoodDamage).Density(6).Penetration(1.5).Resist(1.15).Build()).
	Size(9).
	Color(2).
	Value(120).
	Shape(GenericShape(3)).
	Build()

var Pentagon *Definition = New().
	Label("Pentagon").
	Parent(Food).
	Body(NewBody().Health(10 * BaseFoodHealth).Damage(1.5 * BaseFoodDamage).Density(8).Penetration(1.1).Resist(1.25).Build()).
	Size(16).
	Color(14).
	Value(400).
	Shape(GenericShape(5)).
	Build()

var BetaPentagon *Definition = New().
	Label("Beta Pentagon").
	Parent(Food).
	Body(NewBody().Health(50 * BaseFoodHealth).Damage(2 * BaseFoodDamage).Density(30).Resist(math.Pow(1.25, 2)).Shield(20 * BaseFoodHealth).Regeneration(0.2).Build()).
	Size(30).
	Color(14).
	Value(2500).
	Shape(GenericShape(5)).
	Build()

var AlphaPentagon *Definition = New().
	Label("Alpha Pentagon").
	Parent(Food).
	Body(NewBody().Health(300 * BaseFoodHealth).Damage(2 * BaseFoodDamage).Density(80).Resist(math.Pow(1.25, 3)).Shield(40 * BaseFoodHealth).Regeneration(0.6).Build()).
	Size(58).
	Color(14).
	Value(15000).
	Shape(GenericShape(-5)).
	Build()

// Rare Food

var GreenEgg *Definition = makeGreenFood(Egg)
var GreenSquare *Definition = makeGreenFood(Square)
var GreenTriangle *Definition = makeGreenFood(Triangle)
var GreenPentagon *Definition = makeGreenFood(Pentagon)
var GreenBetaPentagon *Definition = makeGreenFood(BetaPentagon)
var GreenAlphaPentagon *Definition = makeGreenFood(AlphaPentagon)

var Gem *Definition = New().
	Label("Gem").
	Parent(Food).
	Body(NewBody().Health(50 * BaseFoodHealth).Damage(1.5 * BaseFoodDamage).Density(8).Penetration(1.1).Resist(1.25).Build()).
	Size(16).
	Color(0).
	Value(3000).
	Shape(GenericShape(6)).
	Build()

var BetaGem *Definition = New().
	Label("Beta Gem").
	Parent(Food).
	Body(NewBody().Health(300 * BaseFoodHealth).Damage(1.5 * BaseFoodDamage).Density(8).Penetration(1.1).Resist(1.25).Shield(100 * BaseFoodHealth).Regeneration(1).Build()).
	Size(35).
	Color(0).
	Value(9000).
	Shape(GenericShape(6)).
	Build()

var AlphaGem *Definition = New().
	Label("Alpha Gem").
	Parent(Food).
	Body(NewBody().Health(1000 * BaseFoodHealth).Damage(1.5 * BaseFoodDamage).Density(8).Penetration(1.1).Resist(1.25).Shield(300 * BaseFoodHealth).Regeneration(1.75).Build()).
	Size(73).
	Color(0).
	Value(27000).
	Shape(GenericShape(-6)).
	Build()
