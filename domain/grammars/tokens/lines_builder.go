package tokens

import "errors"

type linesBuilder struct {
	list []Line
}

func createLinesBuilder() LinesBuilder {
	out := linesBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *linesBuilder) Create() LinesBuilder {
	return createLinesBuilder()
}

// WithList adds a list of elementWithCardinality instances to the builder
func (app *linesBuilder) WithList(elements []Line) LinesBuilder {
	app.list = elements
	return app
}

// Now builds a line instance
func (app *linesBuilder) Now() (Lines, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Line instance in order to build a Lines instance")
	}

	return createLines(app.list), nil
}
