package game

import "math"

func NewHealthType(health, resist float64, dynamic bool) (ht *HealthType) {
	ht = &HealthType{
		Max:     health,
		Amount:  health,
		Resist:  resist,
		Regen:   0,
		Dynamic: dynamic,
	}

	return
}

func (ht *HealthType) Set(health, regen float64) {
	if ht.Max > 0 {
		ht.Amount = ht.Amount / ht.Max * health
	} else {
		ht.Amount = health
	}

	ht.Max = health
	ht.Regen = regen
}

func (ht *HealthType) Ratio() (ratio float64) {
	if ht.Max > 0 {
		ratio = ht.Amount / ht.Max
	}

	return
}

func (ht *HealthType) Display() (ratio float64) {
	ratio = min(max(ht.Ratio(), 0), 1)
	return
}

func (ht *HealthType) GetDamage(amount float64, capped bool) (actual float64) {
	if ht.Dynamic {
		var permeability float64 = 0
		if ht.Max > 0 {
			permeability = min(max(ht.Amount/ht.Max, 0), 1)
		}

		actual = amount * permeability
	} else {
		actual = amount
	}

	if capped {
		actual = min(actual, ht.Amount)
	}

	return
}

func (ht *HealthType) Regenerate(boost bool) {
	var boostFactor, constant float64 = 0, 5
	if boost {
		boostFactor = 0.5
	}

	if ht.Dynamic {
		var r float64 = ht.Display()
		switch r {
		case 0:
			ht.Amount = 0.0001
		case 1:
			ht.Amount = ht.Max
		default:
			ht.Amount += constant * (ht.Regen*math.Exp(-50*math.Pow(math.Sqrt(0.5*r)-0.4, 2))/3 + r*ht.Max/10/15 + boostFactor)
		}
	} else {
		if ht.Amount < ht.Max {
			ht.Amount += constant * (ht.Max/10/60/2.5 + boostFactor)
		}
	}

	ht.Amount = min(max(ht.Amount, 0), ht.Max)
}
