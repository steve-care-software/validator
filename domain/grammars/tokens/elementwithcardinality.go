package tokens

import "github.com/steve-care-software/validator/domain/grammars/cardinality"

type elementWithCardinality struct {
	element     Element
	cardinality cardinality.Cardinality
}

func createElementWithCardinality(
	element Element,
	cardinality cardinality.Cardinality,
) ElementWithCardinality {
	out := elementWithCardinality{
		element:     element,
		cardinality: cardinality,
	}

	return &out
}

// Element returns the element
func (obj *elementWithCardinality) Element() Element {
	return obj.element
}

// Cardinality returns the cardinality
func (obj *elementWithCardinality) Cardinality() cardinality.Cardinality {
	return obj.cardinality
}
