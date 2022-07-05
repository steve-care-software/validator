package grammars

import "errors"

type externalBuilder struct {
	token   string
	grammar Grammar
}

func createExternalBuilder() ExternalBuilder {
	out := externalBuilder{
		token:   "",
		grammar: nil,
	}

	return &out
}

// Create initializes the builder
func (app *externalBuilder) Create() ExternalBuilder {
	return createExternalBuilder()
}

// WithToken adds a token to the builder
func (app *externalBuilder) WithToken(token string) ExternalBuilder {
	app.token = token
	return app
}

// WithGrammar adds a grammar to the builder
func (app *externalBuilder) WithGrammar(grammar Grammar) ExternalBuilder {
	app.grammar = grammar
	return app
}

// Now builds a new Grammar instance
func (app *externalBuilder) Now() (External, error) {
	if app.token == "" {
		return nil, errors.New("the token is mandatory in order to build an External instance")
	}

	if app.grammar == nil {
		return nil, errors.New("the grammar is mandatory in order to build an External instance")
	}

	return createExternal(app.token, app.grammar), nil
}
