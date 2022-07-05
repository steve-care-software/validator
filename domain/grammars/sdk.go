package grammars

import (
	"github.com/steve-care-software/validator/domain/grammars/cardinality"
	"github.com/steve-care-software/validator/domain/grammars/channels"
	"github.com/steve-care-software/validator/domain/grammars/tokens"
)

// NewAdapterBuilder creates a new adapter builder
func NewAdapterBuilder() AdapterBuilder {
	grammarBuilder := NewBuilder()
	externalsBuilder := NewExternalsBuilder()
	channelsBuilder := channels.NewBuilder()
	channelBuilder := channels.NewChannelBuilder()
	conditionBuilder := channels.NewConditionBuilder()
	tokenBuilder := tokens.NewTokenBuilder()
	linesBuilder := tokens.NewLinesBuilder()
	lineBuilder := tokens.NewLineBuilder()
	elementWithCardinalityBuilder := tokens.NewElementWithCardinalityBuilder()
	elementBuilder := tokens.NewElementBuilder()
	cardinalityAdapter := cardinality.NewAdapter()
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
	commentPrefix := []byte(";")[0]
	commentSuffix := []byte(";")[0]
	tokenNameCharacters := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	channelCharacters := []byte{
		[]byte("\t")[0],
		[]byte("\n")[0],
		[]byte("\r")[0],
		[]byte(" ")[0],
	}

	return createAdapterBuilder(
		grammarBuilder,
		externalsBuilder,
		channelsBuilder,
		channelBuilder,
		conditionBuilder,
		tokenBuilder,
		linesBuilder,
		lineBuilder,
		elementWithCardinalityBuilder,
		elementBuilder,
		cardinalityAdapter,
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
		commentPrefix,
		commentSuffix,
		tokenNameCharacters,
		channelCharacters,
	)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewExternalsBuilder creates a new externals builder
func NewExternalsBuilder() ExternalsBuilder {
	return createExternalsBuilder()
}

// NewExternalBuilder creates a new external builder
func NewExternalBuilder() ExternalBuilder {
	return createExternalBuilder()
}

// AdapterBuilder represents an adapter builder
type AdapterBuilder interface {
	Create() AdapterBuilder
	WithExternals(externals Externals) AdapterBuilder
	Now() (Adapter, error)
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
	WithExternals(externals Externals) Builder
	Now() (Grammar, error)
}

// Grammar represents a lexer grammar
type Grammar interface {
	Root() tokens.Token
	HasChannels() bool
	Channels() channels.Channels
	HasExternals() bool
	Externals() Externals
}

// ExternalsBuilder represents an externals builder
type ExternalsBuilder interface {
	Create() ExternalsBuilder
	WithList(list []External) ExternalsBuilder
	Now() (Externals, error)
}

// Externals represents externals
type Externals interface {
	List() []External
	Find(name string) (External, error)
}

// ExternalBuilder represents an external builder
type ExternalBuilder interface {
	Create() ExternalBuilder
	WithToken(token string) ExternalBuilder
	WithGrammar(grammar Grammar) ExternalBuilder
	Now() (External, error)
}

// External represents an external token
type External interface {
	Token() string
	Grammar() Grammar
}
