package tokens

import "errors"

type lineBuilder struct {
	list []ElementWithCardinality
}

func createLineBuilder() LineBuilder {
	out := lineBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *lineBuilder) Create() LineBuilder {
	return createLineBuilder()
}

// WithList adds a list of elementWithCardinality instances to the builder
func (app *lineBuilder) WithList(elements []ElementWithCardinality) LineBuilder {
	app.list = elements
	return app
}

// Now builds a line instance
func (app *lineBuilder) Now() (Line, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 ElementWithCardinality instance in order to build a Line instance")
	}

	return createLine(app.list), nil
}
