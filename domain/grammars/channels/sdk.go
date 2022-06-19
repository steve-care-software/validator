package channels

import (
	"github.com/steve-care-software/validator/domain/grammars/tokens"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewChannelBuilder creates a new channel builder
func NewChannelBuilder() ChannelBuilder {
	return createChannelBuilder()
}

// NewConditionBuilder creates a new condition builder
func NewConditionBuilder() ConditionBuilder {
	return createConditionBuilder()
}

// Builder represents a channels builder
type Builder interface {
	Create() Builder
	WithList(list []Channel) Builder
	Now() (Channels, error)
}

// Channels represents channels
type Channels interface {
	List() []Channel
}

// ChannelBuilder represents a channel builder
type ChannelBuilder interface {
	Create() ChannelBuilder
	WithToken(token tokens.Token) ChannelBuilder
	WithCondition(condition Condition) ChannelBuilder
	Now() (Channel, error)
}

// Channel represents a channel
type Channel interface {
	Token() tokens.Token
	HasCondition() bool
	Condition() Condition
}

// ConditionBuilder represents the condition builder
type ConditionBuilder interface {
	Create() ConditionBuilder
	WithPrevious(prev tokens.Token) ConditionBuilder
	WithNext(next tokens.Token) ConditionBuilder
	Now() (Condition, error)
}

// Condition represents a channel condition
type Condition interface {
	IsPrevious() bool
	Previous() tokens.Token
	IsNext() bool
	Next() tokens.Token
}
