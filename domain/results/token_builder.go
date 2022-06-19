package results

import "errors"

type tokenBuilder struct {
	name      string
	block     Block
	pChannels *uint
}

func createTokenBuilder() TokenBuilder {
	out := tokenBuilder{
		name:      "",
		block:     nil,
		pChannels: nil,
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

// WithBlock adds a block to the builder
func (app *tokenBuilder) WithBlock(block Block) TokenBuilder {
	app.block = block
	return app
}

// WithChannels add channels to the builder
func (app *tokenBuilder) WithChannels(channels uint) TokenBuilder {
	app.pChannels = &channels
	return app
}

// Now builds a new Token instance
func (app *tokenBuilder) Now() (Token, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Token instance")
	}

	if app.block == nil {
		return nil, errors.New("the block is mandatory in order to build a Token instance")
	}

	if app.pChannels == nil {
		return nil, errors.New("the channels is mandatory in order to build a Token instance")
	}

	return createToken(app.name, app.block, *app.pChannels), nil
}
