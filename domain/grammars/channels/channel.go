package channels

import "github.com/steve-care-software/validator/domain/grammars/tokens"

type channel struct {
	token     tokens.Token
	condition Condition
}

func createChannel(
	token tokens.Token,
) Channel {
	return createChannelInternally(token, nil)
}

func createChannelWithCondition(
	token tokens.Token,
	condition Condition,
) Channel {
	return createChannelInternally(token, condition)
}

func createChannelInternally(
	token tokens.Token,
	condition Condition,
) Channel {
	out := channel{
		token:     token,
		condition: condition,
	}

	return &out
}

// Token returns the token
func (obj *channel) Token() tokens.Token {
	return obj.token
}

// HasCondition returns true if there is a condition, false otherwise
func (obj *channel) HasCondition() bool {
	return obj.condition != nil
}

// Condition returns the condition, if any
func (obj *channel) Condition() Condition {
	return obj.condition
}
