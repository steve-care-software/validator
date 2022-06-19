package results

import (
	"github.com/steve-care-software/validator/domain/grammars/tokens"
)

// NewBuilder creates a new result builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewTokenBuilder creates a new token builder
func NewTokenBuilder() TokenBuilder {
	return createTokenBuilder()
}

// NewBlockBuilder creates a new block builder
func NewBlockBuilder() BlockBuilder {
	return createBlockBuilder()
}

// NewLineBuilder creates a new line builder
func NewLineBuilder() LineBuilder {
	return createLineBuilder()
}

// NewElementWithCardinalityBuilder creates a new elementWithCardinality builder
func NewElementWithCardinalityBuilder() ElementWithCardinalityBuilder {
	return createElementWithCardinalityBuilder()
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	return createElementBuilder()
}

// Builder represents the  result builder
type Builder interface {
	Create() Builder
	WithIndex(index uint) Builder
	WithToken(token Token) Builder
	Now() (Result, error)
}

// Result represents a result
type Result interface {
	Index() uint
	Cursor() uint
	Token() Token
}

// TokenBuilder represents a token builder
type TokenBuilder interface {
	Create() TokenBuilder
	WithName(name string) TokenBuilder
	WithBlock(block Block) TokenBuilder
	WithChannels(channels uint) TokenBuilder
	Now() (Token, error)
}

// Token represents a token
type Token interface {
	Name() string
	Block() Block
	Channels() uint
	Path() []string
	IsSuccess() bool
}

// BlockBuilder represents a block builder
type BlockBuilder interface {
	Create() BlockBuilder
	WithInput(input []byte) BlockBuilder
	WithList(list []Line) BlockBuilder
	Now() (Block, error)
}

// Block represents a block of lines
type Block interface {
	Discovered() uint
	Channels() uint
	IsSuccess() bool
	Input() []byte
	List() []Line
	Remaining() []byte
	HasMatch() bool
	Match() Line
}

// LineBuilder represents a line builder
type LineBuilder interface {
	Create() LineBuilder
	WithIndex(index uint) LineBuilder
	WithElements(elements []ElementWithCardinality) LineBuilder
	Now() (Line, error)
}

// Line represents a line
type Line interface {
	Discovered() uint
	Channels() uint
	IsSuccess() bool
	Index() uint
	Elements() []ElementWithCardinality
	Remaining() []byte
}

// ElementWithCardinalityBuilder represents an element with cardinality builder
type ElementWithCardinalityBuilder interface {
	Create() ElementWithCardinalityBuilder
	WithMissing(missing uint) ElementWithCardinalityBuilder
	WithElement(element tokens.Element) ElementWithCardinalityBuilder
	WithMatches(matches []Element) ElementWithCardinalityBuilder
	WithRemaining(remaining []byte) ElementWithCardinalityBuilder
	Now() (ElementWithCardinality, error)
}

// ElementWithCardinality represents an element with cardinality
type ElementWithCardinality interface {
	Discovered() uint
	Channels() uint
	IsSuccess() bool
	Amount() uint
	Missing() uint
	Element() tokens.Element
	Path() []string
	HasMatches() bool
	Matches() []Element
	Remaining() []byte
}

// ElementBuilder represents an element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithValue(value byte) ElementBuilder
	WithToken(token Token) ElementBuilder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	Discovered() uint
	Channels() uint
	IsSuccess() bool
	IsValue() bool
	Value() *byte
	IsToken() bool
	Token() Token
}
