package configs

import (
	"math"
	"reflect"
)

var arrasToGennScale float64 = 5.0
var definitionPtrType reflect.Type = reflect.TypeOf((*Definition)(nil))

func gunArrasToGenn(aLength, aWidth, aAspect, aX, aY, aAngle, aDelay float64) (length, width, endWidth, x, y, angle, delay float64) {
	length, width, endWidth, x, y, angle, delay = aLength*arrasToGennScale, aWidth*arrasToGennScale, aWidth*arrasToGennScale*aAspect, aX*arrasToGennScale, aY*arrasToGennScale, aAngle, aDelay
	return
}

func turretArrasToGenn(aSize, aX, aY, aAngle, aArc float64, aLayer int) (size, x, y, angle, arc float64, layer int) {
	size, x, y, angle, arc, layer = aSize*arrasToGennScale, aX*arrasToGennScale, aY*arrasToGennScale, aAngle, aArc, aLayer
	return
}

func skillSetArrasToGenn(rld, dam, pen, str, spd, atk, hlt, shi, rgn, mob float64) (out Skills) {
	var convert = func(a float64) int {
		return int(math.Round(a * float64(SkillCapNormal)))
	}

	out = Skills{
		BodyDamage:         convert(atk),
		MaxHealth:          convert(hlt),
		BulletSpeed:        convert(spd),
		BulletHealth:       convert(str),
		BulletPenetration:  convert(pen),
		BulletDamage:       convert(dam),
		Reload:             convert(rld),
		Speed:              convert(mob),
		ShieldRegeneration: convert(rgn),
		ShieldCapacity:     convert(shi),
	}

	return
}

func cloneDef(source *Definition) (out *Definition) {
	if source == nil {
		return
	}

	out = reflect.New(reflect.TypeFor[Definition]()).Interface().(*Definition)

	var sourceValue, outValue reflect.Value = reflect.ValueOf(source).Elem(), reflect.ValueOf(out).Elem()

	for i := range sourceValue.NumField() {
		outValue.Field(i).Set(cloneValue(sourceValue.Field(i)))
	}

	// Update the index
	var nowID DefinitionID = idCounter
	out.Index = &nowID

	idCounter++
	return
}

func cloneValue(source reflect.Value) (clone reflect.Value) {
	switch source.Kind() {
	case reflect.Pointer:
		if source.IsNil() {
			clone = reflect.Zero(source.Type())
			return
		}

		if source.Type() == definitionPtrType {
			clone = source
			return
		}

		clone = reflect.New(source.Type().Elem())
		clone.Elem().Set(cloneValue(source.Elem()))
	case reflect.Slice:
		if source.IsNil() {
			clone = reflect.Zero(source.Type())
			return
		}

		clone = reflect.MakeSlice(source.Type(), source.Len(), source.Len())
		for i := range source.Len() {
			clone.Index(i).Set(cloneValue(source.Index(i)))
		}
	case reflect.Array:
		clone = reflect.New(source.Type()).Elem()
		for i := range source.Len() {
			clone.Index(i).Set(cloneValue(source.Index(i)))
		}
	case reflect.Struct:
		clone = reflect.New(source.Type()).Elem()
		for i := range source.NumField() {
			clone.Field(i).Set(cloneValue(source.Field(i)))
		}
	default:
		clone = source
	}

	return
}
