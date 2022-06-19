package results

import "errors"

type builder struct {
	pIndex *uint
	token  Token
}

func createBuilder() Builder {
	out := builder{
		pIndex: nil,
		token:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithIndex adds an index to the builder
func (app *builder) WithIndex(index uint) Builder {
	app.pIndex = &index
	return app
}

// WithToken adds a token to the builder
func (app *builder) WithToken(token Token) Builder {
	app.token = token
	return app
}

// Now buildsa new Result instance
func (app *builder) Now() (Result, error) {
	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build a Result instance")
	}

	if app.token == nil {
		return nil, errors.New("the token is mandatory in order to build a Result instance")
	}

	return createResult(*app.pIndex, app.token), nil
}
