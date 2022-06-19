package channels

import "errors"

type builder struct {
	list []Channel
}

func createBuilder() Builder {
	out := builder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithList adds a list of channel to the builder
func (app *builder) WithList(list []Channel) Builder {
	app.list = list
	return app
}

// Now builds a new Channels instance
func (app *builder) Now() (Channels, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Channel instance in the list in order to build a Channels instance")
	}

	return createChannels(app.list), nil
}
