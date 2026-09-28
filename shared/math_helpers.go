package shared

import (
	"math"
	"math/rand/v2"

	"github.com/z46-dev/gamelib/vector"
	"golang.org/x/exp/constraints"
)

func Clamp[T constraints.Ordered](value, _min, _max T) (out T) {
	out = min(_max, max(_min, value))
	return
}

const TAU float64 = math.Pi * 2

// Gets the angle difference between two angles, normalized to the range [-π, π).
func AngleDifference(a, b float64) (diff float64) {
	diff = math.Remainder(a-b, TAU)

	if diff == math.Pi {
		diff = -math.Pi
	}

	return
}

// Normalizes an angle to the range [0, 2π).
func NormalizeAngle(angle float64) (out float64) {
	out = math.Remainder(angle, TAU)

	if out < 0 {
		out += TAU
	}

	return
}

// Loops smoothly from one angle to another, taking the shortest path and applying a slowness factor.
func LoopSmooth(angle, desired, slowness float64) (out float64) {
	out = AngleDifference(angle, desired) / slowness
	return
}

// Approximation, faster
func Gauss(mean, deviation float64) (out float64) {
	out = mean + deviation*((rand.Float64()+rand.Float64()+rand.Float64()+rand.Float64()+rand.Float64()+rand.Float64())-3)
	return
}

// Direct Box-Muller transform, potentially slower but more accurate
func BoxMullerGauss(mean, deviation float64) (out float64) {
	var x1, x2, w float64
	for w == 0 || w >= 1 {
		x1, x2 = 2*rand.Float64()-1, 2*rand.Float64()-1
		w = x1*x1 + x2*x2
	}

	w = math.Sqrt(-2 * math.Log(w) / w)
	out = mean + deviation*x1*w
	return
}

// GaussInverse returns a random number between min and max, with a clustering
// factor that determines how tightly the values are clustered around the mean.
func GaussInverse(min, max, clustering float64) (out float64) {
	var (
		rng float64 = max - min
		x   float64 = Gauss(0, rng/clustering)
	)

	if out = math.Mod(x, rng); out < 0 {
		out += rng
	}

	out += min
	return
}

// GaussRing returns a random point in a ring with a given radius and clustering factor.
func GaussRing(radius, clustering float64) (pos *vector.Vec2[float64]) {
	pos = vector.Vec2FromAngleMagnitude(TAU*rand.Float64(), Gauss(radius, radius*clustering))
	return
}

// Choose returns a random element from the provided slice of choices.
func Choose[T any](choices []T) (out T) {
	out = choices[rand.IntN(len(choices))]
	return
}

// ChooseN returns a slice of n random elements from the provided slice of choices.
func ChooseN[T any](choices []T, n int) (out []T) {
	n = min(n, len(choices))
	out = make([]T, n)
	var perm []int = rand.Perm(len(choices))

	for i := range n {
		out[i] = choices[perm[i]]
	}

	return
}

// ChooseChance selects an index from the probabilities slice based on their weights.
// The probabilities slice should contain non-negative values, and the function will
// return an index based on the relative weights of the values.
func ChooseChance(probabilities []float64) (choice int) {
	var totalProbability float64 = 0
	for _, p := range probabilities {
		totalProbability += p
	}

	var answer float64 = rand.Float64() * totalProbability

	for i, p := range probabilities {
		if answer < p {
			choice = i
			return
		}

		answer -= p
	}

	return
}
