package channels

import (
	"errors"

	"github.com/steve-care-software/validator/domain/grammars/tokens"
)

type conditionBuilder struct {
	prev tokens.Token
	next tokens.Token
}

func createConditionBuilder() ConditionBuilder {
	out := conditionBuilder{
		prev: nil,
		next: nil,
	}

	return &out
}

// Create initializes the builder
func (app *conditionBuilder) Create() ConditionBuilder {
	return createConditionBuilder()
}

// WithPrevious adds a previous condition to the builder
func (app *conditionBuilder) WithPrevious(prev tokens.Token) ConditionBuilder {
	app.prev = prev
	return app
}

// WithNext adds a next condition to the builder
func (app *conditionBuilder) WithNext(next tokens.Token) ConditionBuilder {
	app.next = next
	return app
}

// Now builds a new Condition instance
func (app *conditionBuilder) Now() (Condition, error) {
	if app.prev != nil {
		return createConditionWithPrevious(app.prev), nil
	}

	if app.next != nil {
		return createConditionWithNext(app.next), nil
	}

	return nil, errors.New("the Condition is invalid")
}
