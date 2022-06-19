package applications

import (
	"github.com/steve-care-software/validator/domain/grammars"
	"github.com/steve-care-software/validator/domain/grammars/cardinality"
	"github.com/steve-care-software/validator/domain/grammars/channels"
	"github.com/steve-care-software/validator/domain/grammars/tokens"
)

// NewGrammarForTests creates a new grammar for tests
func NewGrammarForTests(root tokens.Token) grammars.Grammar {
	return NewGrammarWithChannelsForTests(root, nil)
}

// NewGrammarWithChannelsForTests creates a new grammar with channels for tests
func NewGrammarWithChannelsForTests(root tokens.Token, list []channels.Channel) grammars.Grammar {
	grammarBuilder := grammars.NewBuilder().Create().WithRoot(root)
	if list != nil {
		chans, err := channels.NewBuilder().WithList(list).Now()
		if err != nil {
			panic(err)
		}

		grammarBuilder.WithChannels(chans)
	}

	ins, err := grammarBuilder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewChannelForTests creates a new channel for tests
func NewChannelForTests(token tokens.Token) channels.Channel {
	return NewChannelWithConditionsForTests(token, nil)
}

// NewChannelWithConditionsForTests creates a new channel woth conditions for tests
func NewChannelWithConditionsForTests(token tokens.Token, condition channels.Condition) channels.Channel {
	builder := channels.NewChannelBuilder().Create().WithToken(token)
	if condition != nil {
		builder.WithCondition(condition)
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewConditionWithNext creates a new condition with next
func NewConditionWithNext(next tokens.Token) channels.Condition {
	ins, err := channels.NewConditionBuilder().Create().WithNext(next).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewConditionWithPrevious creates a new condition with previous
func NewConditionWithPrevious(previous tokens.Token) channels.Condition {
	ins, err := channels.NewConditionBuilder().Create().WithPrevious(previous).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewCardinalityWithSpecificForTests creates cardinality with specific for tests
func NewCardinalityWithSpecificForTests(specific uint8) cardinality.Cardinality {
	cardinality, err := cardinality.NewBuilder().Create().WithMinimum(specific).WithMaximum(specific).Now()
	if err != nil {
		panic(err)
	}

	return cardinality
}

// NewLineWithElementWithCardinalityList creates a new line with ElementWithCardinality list for tests
func NewLineWithElementWithCardinalityList(list []tokens.ElementWithCardinality) tokens.Line {
	line, err := tokens.NewLineBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return line
}

// NewElementWithCardinalityWithReferenceAndCardinalityForTests creates a new elementWithCardinality with token reference and cardinality for  tests
func NewElementWithCardinalityWithReferenceAndCardinalityForTests(refName string, cardinality cardinality.Cardinality) tokens.ElementWithCardinality {
	element, err := tokens.NewElementBuilder().Create().WithReference(refName).Now()
	if err != nil {
		panic(err)
	}

	return NewElementWithCardinalityWithElementAndCardinalityForTests(element, cardinality)
}

// NewTokenWithRangeCardinalityWithByteForTests creates a new token with range cardinality with byte for tests
func NewTokenWithRangeCardinalityWithByteForTests(tokenName string, min uint8, max uint8, byteVal byte) tokens.Token {
	element, err := tokens.NewElementBuilder().Create().WithByte(byteVal).Now()
	if err != nil {
		panic(err)
	}

	cardinality, err := cardinality.NewBuilder().Create().WithMinimum(min).WithMaximum(max).Now()
	if err != nil {
		panic(err)
	}

	return NewTokenWithSingleElementInSingleLineForTests(tokenName, element, cardinality)
}

// NewTokenWithMinimumCardinalityWithByteForTests creates a new token with min cardinality with byte for tests
func NewTokenWithMinimumCardinalityWithByteForTests(tokenName string, min uint8, byteVal byte) tokens.Token {
	element, err := tokens.NewElementBuilder().Create().WithByte(byteVal).Now()
	if err != nil {
		panic(err)
	}

	cardinality, err := cardinality.NewBuilder().Create().WithMinimum(min).Now()
	if err != nil {
		panic(err)
	}

	return NewTokenWithSingleElementInSingleLineForTests(tokenName, element, cardinality)
}

// NewTokenWithSpecificCardinalityWithByteForTests creates a new token with specific cardinality with byte for tests
func NewTokenWithSpecificCardinalityWithByteForTests(tokenName string, specific uint8, byteVal byte) tokens.Token {
	element, err := tokens.NewElementBuilder().Create().WithByte(byteVal).Now()
	if err != nil {
		panic(err)
	}

	cardinality := NewCardinalityWithSpecificForTests(specific)
	return NewTokenWithSingleElementInSingleLineForTests(tokenName, element, cardinality)
}

// NewTokenWithSingleElementInSingleLineForTests creates a new token with single element in a singleline for tests
func NewTokenWithSingleElementInSingleLineForTests(tokenName string, element tokens.Element, cardinality cardinality.Cardinality) tokens.Token {
	elementWithCardinality := NewElementWithCardinalityWithElementAndCardinalityForTests(element, cardinality)
	line := NewLineWithElementWithCardinalityList([]tokens.ElementWithCardinality{
		elementWithCardinality,
	})

	return NewTokenWithLinesForTests(tokenName, []tokens.Line{
		line,
	})
}

// NewTokenWithLinesForTests creates a new token with lines for tests
func NewTokenWithLinesForTests(tokenName string, linesList []tokens.Line) tokens.Token {
	lines, err := tokens.NewLinesBuilder().Create().WithList(linesList).Now()
	if err != nil {
		panic(err)
	}

	token, err := tokens.NewTokenBuilder().Create().WithName(tokenName).WithLines(lines).Now()
	if err != nil {
		panic(err)
	}

	return token
}

// NewElementWithTokenForTests creates a new element with token for tests
func NewElementWithTokenForTests(token tokens.Token) tokens.Element {
	element, err := tokens.NewElementBuilder().Create().WithToken(token).Now()
	if err != nil {
		panic(err)
	}

	return element
}

// NewElementWithCardinalityWithElementAndCardinalityForTests creates a new elementWithCardinality with element and cardinality for tests
func NewElementWithCardinalityWithElementAndCardinalityForTests(element tokens.Element, cardinality cardinality.Cardinality) tokens.ElementWithCardinality {
	elementWithCardinality, err := tokens.NewElementWithCardinalityBuilder().
		Create().
		WithElement(element).
		WithCardinality(cardinality).
		Now()

	if err != nil {
		panic(err)
	}

	return elementWithCardinality
}

// NewElementWithCardinalityWithTokenAndCardinalityForTests creates a new elementWithCardinality with token and cardinality for tests
func NewElementWithCardinalityWithTokenAndCardinalityForTests(token tokens.Token, cardinality cardinality.Cardinality) tokens.ElementWithCardinality {
	element := NewElementWithTokenForTests(token)
	return NewElementWithCardinalityWithElementAndCardinalityForTests(element, cardinality)
}
