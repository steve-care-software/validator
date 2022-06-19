package tokens

import (
	"errors"

	"github.com/steve-care-software/validator/domain/grammars/cardinality"
)

type elementWithCardinalityBuilder struct {
	element     Element
	cardinality cardinality.Cardinality
}

func createElementWithCardinalityBuilder() ElementWithCardinalityBuilder {
	out := elementWithCardinalityBuilder{
		element:     nil,
		cardinality: nil,
	}

	return &out
}

// Create initializes the builder
func (app *elementWithCardinalityBuilder) Create() ElementWithCardinalityBuilder {
	return createElementWithCardinalityBuilder()
}

// WithElement adds an element to the builder
func (app *elementWithCardinalityBuilder) WithElement(element Element) ElementWithCardinalityBuilder {
	app.element = element
	return app
}

// WithCardinality adds a cardinality to the builder
func (app *elementWithCardinalityBuilder) WithCardinality(cardinality cardinality.Cardinality) ElementWithCardinalityBuilder {
	app.cardinality = cardinality
	return app
}

// Now builds a new ElementWithCardinality instance
func (app *elementWithCardinalityBuilder) Now() (ElementWithCardinality, error) {
	if app.element == nil {
		return nil, errors.New("the element is mandatory in order to build an ElementWithCardinality instance")
	}

	if app.cardinality == nil {
		return nil, errors.New("the cardinality is mandatory in order to build an ElementWithCardinality instance")
	}

	return createElementWithCardinality(app.element, app.cardinality), nil
}
