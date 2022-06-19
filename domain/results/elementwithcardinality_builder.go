package results

import (
	"errors"

	"github.com/steve-care-software/validator/domain/grammars/tokens"
)

type elementWithCardinalityBuilder struct {
	pMissing  *uint
	element   tokens.Element
	matches   []Element
	remaining []byte
}

func createElementWithCardinalityBuilder() ElementWithCardinalityBuilder {
	out := elementWithCardinalityBuilder{
		pMissing:  nil,
		element:   nil,
		matches:   nil,
		remaining: nil,
	}

	return &out
}

// Create initializes the builder
func (app *elementWithCardinalityBuilder) Create() ElementWithCardinalityBuilder {
	return createElementWithCardinalityBuilder()
}

// WithMissing adds a missing to the builder
func (app *elementWithCardinalityBuilder) WithMissing(missing uint) ElementWithCardinalityBuilder {
	app.pMissing = &missing
	return app
}

// WithElement adds an element to the builder
func (app *elementWithCardinalityBuilder) WithElement(element tokens.Element) ElementWithCardinalityBuilder {
	app.element = element
	return app
}

// WithMatches add matches to the builder
func (app *elementWithCardinalityBuilder) WithMatches(matches []Element) ElementWithCardinalityBuilder {
	app.matches = matches
	return app
}

// WithRemaining adds a remaining to the builder
func (app *elementWithCardinalityBuilder) WithRemaining(remaining []byte) ElementWithCardinalityBuilder {
	app.remaining = remaining
	return app
}

// Now builds a new ElementWithCardinality instance
func (app *elementWithCardinalityBuilder) Now() (ElementWithCardinality, error) {
	if app.pMissing == nil {
		return nil, errors.New("the missing is mandatory in order to build an ElementWithCardinality instance")
	}

	if app.element == nil {
		return nil, errors.New("the element is mandatory in order to build an ElementWithCardinality instance")
	}

	if app.matches != nil && len(app.matches) <= 0 {
		app.matches = nil
	}

	if app.remaining != nil && len(app.remaining) <= 0 {
		app.remaining = []byte{}
	}

	if app.matches != nil {
		return createElementWithCardinalityWithMatch(*app.pMissing, app.element, app.remaining, app.matches), nil
	}

	return createElementWithCardinality(*app.pMissing, app.element, app.remaining), nil
}
