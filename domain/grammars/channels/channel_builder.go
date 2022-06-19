package channels

import (
	"errors"

	"github.com/steve-care-software/validator/domain/grammars/tokens"
)

type channelBuilder struct {
	token     tokens.Token
	condition Condition
}

func createChannelBuilder() ChannelBuilder {
	out := channelBuilder{
		token:     nil,
		condition: nil,
	}

	return &out
}

// Create initializes the builder
func (app *channelBuilder) Create() ChannelBuilder {
	return createChannelBuilder()
}

// WithToken adds a token to the builder
func (app *channelBuilder) WithToken(token tokens.Token) ChannelBuilder {
	app.token = token
	return app
}

// WithCondition adds a condition to the builder
func (app *channelBuilder) WithCondition(condition Condition) ChannelBuilder {
	app.condition = condition
	return app
}

// Now builds a new Channel instance
func (app *channelBuilder) Now() (Channel, error) {
	if app.token == nil {
		return nil, errors.New("the token is mandatory in order to build a Channel instance")
	}

	if app.condition != nil {
		return createChannelWithCondition(app.token, app.condition), nil
	}

	return createChannel(app.token), nil
}
