package configs

var arrasToGennScale float64 = 5.0

func gunArrasToGenn(aLength, aWidth, aAspect, aX, aY, aAngle, aDelay float64) (length, width, endWidth, x, y, angle, delay float64) {
	length, width, endWidth, x, y, angle, delay = aLength*arrasToGennScale, aWidth*arrasToGennScale, aWidth*arrasToGennScale*aAspect, aX*arrasToGennScale, aY*arrasToGennScale, aAngle, aDelay
	return
}

func turretArrasToGenn(aSize, aX, aY, aAngle, aArc float64, aLayer int) (size, x, y, angle, arc float64, layer int) {
	size, x, y, angle, arc, layer = aSize*arrasToGennScale, aX*arrasToGennScale, aY*arrasToGennScale, aAngle, aArc, aLayer
	return
}

func cloneDef(source *Definition) (out *Definition) {
	if source == nil {
		return
	}

	// TODO: Implement a proper deep copy for Definition. This would be used to copy the top level of the definition. It would not "clone" the parents or the types of turrets/guns, but it would copy the guns/turrets. Ideally the things that *belong* to the definition would be copied so we don't have to reuse the reference.
	return
}
