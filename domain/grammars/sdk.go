package grammars

import (
	"github.com/steve-care-software/validator/domain/grammars/cardinality"
	"github.com/steve-care-software/validator/domain/grammars/channels"
	"github.com/steve-care-software/validator/domain/grammars/tokens"
)

// NewAdapter creates a new adapter
func NewAdapter() Adapter {
	grammarBuilder := NewBuilder()
	channelsBuilder := channels.NewBuilder()
	channelBuilder := channels.NewChannelBuilder()
	conditionBuilder := channels.NewConditionBuilder()
	tokenBuilder := tokens.NewTokenBuilder()
	linesBuilder := tokens.NewLinesBuilder()
	lineBuilder := tokens.NewLineBuilder()
	elementWithCardinalityBuilder := tokens.NewElementWithCardinalityBuilder()
	elementBuilder := tokens.NewElementBuilder()
	cardinalityBuilder := cardinality.NewBuilder()
	rootPrefix := []byte("%")[0]
	rootSuffix := []byte(";")[0]
	channelPrefix := []byte("-")[0]
	channelSuffix := []byte(";")[0]
	channelConditionPrevious := []byte("<")[0]
	channelConditionNext := []byte(">")[0]
	channelConditionAnd := []byte("&")[0]
	tokenNamePrefix := []byte(".")[0]
	bytePrefix := []byte("$")[0]
	linesPrefix := []byte(":")[0]
	linesSuffix := []byte(";")[0]
	lineDelimiter := []byte("|")[0]
	cardinalityNonZeroMultiple := []byte("+")[0]
	cardinalityZeroMultiple := []byte("*")[0]
	cardinalityOptional := []byte("?")[0]
	cardinalityRangePrefix := []byte("[")[0]
	cardinalityRangeSuffix := []byte("]")[0]
	cardinalityRangeSeparator := []byte(",")[0]
	commentPrefix := []byte(";")[0]
	commentSuffix := []byte(";")[0]
	numbersCharacters := []byte("0123456789")
	tokenNameCharacters := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	channelCharacters := []byte{
		[]byte("\t")[0],
		[]byte("\n")[0],
		[]byte("\r")[0],
		[]byte(" ")[0],
	}

	return createAdapter(
		grammarBuilder,
		channelsBuilder,
		channelBuilder,
		conditionBuilder,
		tokenBuilder,
		linesBuilder,
		lineBuilder,
		elementWithCardinalityBuilder,
		elementBuilder,
		cardinalityBuilder,
		rootPrefix,
		rootSuffix,
		channelPrefix,
		channelSuffix,
		channelConditionPrevious,
		channelConditionNext,
		channelConditionAnd,
		tokenNamePrefix,
		bytePrefix,
		linesPrefix,
		linesSuffix,
		lineDelimiter,
		cardinalityNonZeroMultiple,
		cardinalityZeroMultiple,
		cardinalityOptional,
		cardinalityRangePrefix,
		cardinalityRangeSuffix,
		cardinalityRangeSeparator,
		commentPrefix,
		commentSuffix,
		numbersCharacters,
		tokenNameCharacters,
		channelCharacters,
	)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents the grammar adapter
type Adapter interface {
	ToGrammar(script string) (Grammar, error)
}

// Builder represents the grammar builder
type Builder interface {
	Create() Builder
	WithRoot(root tokens.Token) Builder
	WithChannels(channels channels.Channels) Builder
	Now() (Grammar, error)
}

// Grammar represents a lexer grammar
type Grammar interface {
	Root() tokens.Token
	HasChannels() bool
	Channels() channels.Channels
}
