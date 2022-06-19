package results

import "errors"

type blockBuilder struct {
	input []byte
	list  []Line
}

func createBlockBuilder() BlockBuilder {
	out := blockBuilder{
		input: nil,
		list:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *blockBuilder) Create() BlockBuilder {
	return createBlockBuilder()
}

// WithInput adds an input to the builder
func (app *blockBuilder) WithInput(input []byte) BlockBuilder {
	app.input = input
	return app
}

// WithList adds a list to the builder
func (app *blockBuilder) WithList(list []Line) BlockBuilder {
	app.list = list
	return app
}

// Now builds a new Block instance
func (app *blockBuilder) Now() (Block, error) {
	if app.input != nil && len(app.input) <= 0 {
		app.input = nil
	}

	if app.input == nil {
		return nil, errors.New("the input data is mandatory in order to build a Block instance")
	}

	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Line in order to build a Block instance")
	}

	var match Line
	for _, oneLine := range app.list {
		matches := true
		elements := oneLine.Elements()
		for _, oneElement := range elements {
			if !oneElement.IsSuccess() {
				matches = false
				break
			}
		}

		if matches {
			match = oneLine
			break
		}
	}

	if match != nil {
		return createBlockWithMatch(app.input, app.list, match), nil
	}

	return createBlock(app.input, app.list), nil
}
