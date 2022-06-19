package cardinality

type cardinality struct {
	min  uint8
	pMax *uint8
}

func createCardinality(
	min uint8,
) Cardinality {
	return createCardinalityInternally(min, nil)
}

func createCardinalityWithMaximum(
	min uint8,
	pMax *uint8,
) Cardinality {
	return createCardinalityInternally(min, pMax)
}

func createCardinalityInternally(
	min uint8,
	pMax *uint8,
) Cardinality {
	out := cardinality{
		min:  min,
		pMax: pMax,
	}

	return &out
}

// Min returns the minimum
func (obj *cardinality) Min() uint8 {
	return obj.min
}

// HasMax returns true if there is a max, false otherwise
func (obj *cardinality) HasMax() bool {
	return obj.pMax != nil
}

// Max returns the maximum, if any
func (obj *cardinality) Max() *uint8 {
	return obj.pMax
}
