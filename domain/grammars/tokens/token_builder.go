package tokens

import "errors"

type tokenBuilder struct {
	name  string
	lines Lines
}

func createTokenBuilder() TokenBuilder {
	out := tokenBuilder{
		name:  "",
		lines: nil,
	}

	return &out
}

// Create initializes the builder
func (app *tokenBuilder) Create() TokenBuilder {
	return createTokenBuilder()
}

// WithName adds a name to the builder
func (app *tokenBuilder) WithName(name string) TokenBuilder {
	app.name = name
	return app
}

// WithLines adds a Lines to the builder
func (app *tokenBuilder) WithLines(lines Lines) TokenBuilder {
	app.lines = lines
	return app
}

// Now builds a new Token instance
func (app *tokenBuilder) Now() (Token, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Token instance")
	}

	if app.lines == nil {
		return nil, errors.New("the lines is mandatory in order to build a Token instance")
	}

	return createToken(app.name, app.lines), nil
}
