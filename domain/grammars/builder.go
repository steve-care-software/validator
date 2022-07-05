package grammars

import (
	"errors"

	"github.com/steve-care-software/validator/domain/grammars/channels"
	"github.com/steve-care-software/validator/domain/grammars/tokens"
)

type builder struct {
	root      tokens.Token
	channels  channels.Channels
	externals Externals
}

func createBuilder() Builder {
	out := builder{
		root:      nil,
		channels:  nil,
		externals: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithRoot adds a root token to the builder
func (app *builder) WithRoot(root tokens.Token) Builder {
	app.root = root
	return app
}

// WithChannels add channels to the builder
func (app *builder) WithChannels(channels channels.Channels) Builder {
	app.channels = channels
	return app
}

// WithExternals add externals to the builder
func (app *builder) WithExternals(externals Externals) Builder {
	app.externals = externals
	return app
}

// Now builds a new Grammar instance
func (app *builder) Now() (Grammar, error) {
	if app.root == nil {
		return nil, errors.New("the root token is mandatory in order to build a Grammar instance")
	}

	if app.channels != nil && app.externals != nil {
		return createGrammarWithChannelsAndExternals(app.root, app.channels, app.externals), nil
	}

	if app.externals != nil {
		return createGrammarWithExternals(app.root, app.externals), nil
	}

	if app.channels != nil {
		return createGrammarWithChannels(app.root, app.channels), nil
	}

	return createGrammar(app.root), nil
}
