package channels

import "github.com/steve-care-software/validator/domain/grammars/tokens"

type condition struct {
	prev tokens.Token
	next tokens.Token
}

func createConditionWithPrevious(
	prev tokens.Token,
) Condition {
	return createConditionInternally(prev, nil)
}

func createConditionWithNext(
	next tokens.Token,
) Condition {
	return createConditionInternally(nil, next)
}

func createConditionInternally(
	prev tokens.Token,
	next tokens.Token,
) Condition {
	out := condition{
		prev: prev,
		next: next,
	}

	return &out
}

// IsPrevious returns true if there is a previous condition, false otherwise
func (obj *condition) IsPrevious() bool {
	return obj.prev != nil
}

// Previous returns the previous condition, if any
func (obj *condition) Previous() tokens.Token {
	return obj.prev
}

// IsNext returns true if there is a next condition, false otherwise
func (obj *condition) IsNext() bool {
	return obj.next != nil
}

// Next returns the next condition, if any
func (obj *condition) Next() tokens.Token {
	return obj.next
}
