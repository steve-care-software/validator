package cardinality

import (
	"errors"
	"fmt"
)

type builder struct {
	pMin *uint8
	pMax *uint8
}

func createBuilder() Builder {
	out := builder{
		pMin: nil,
		pMax: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithMinimum adds a minimum to the builder
func (app *builder) WithMinimum(min uint8) Builder {
	app.pMin = &min
	return app
}

// WithMaximum adds a maximum to the builder
func (app *builder) WithMaximum(max uint8) Builder {
	app.pMax = &max
	return app
}

// Now builds a new Cardinality instance
func (app *builder) Now() (Cardinality, error) {
	if app.pMin == nil {
		return nil, errors.New("the minimum is mandatory in order to build a Cardinality instance")
	}

	if app.pMax != nil {
		if *app.pMin > *app.pMax {
			str := fmt.Sprintf("the minimum (%d), must be smaller or equal (<=) than the maximum (%d)", *app.pMin, *app.pMax)
			return nil, errors.New(str)
		}

		return createCardinalityWithMaximum(*app.pMin, app.pMax), nil
	}

	return createCardinality(*app.pMin), nil
}
