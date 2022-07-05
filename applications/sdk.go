package applications

import (
	"github.com/steve-care-software/validator/domain/grammars"
	"github.com/steve-care-software/validator/domain/results"
)

// NewBuilder creates a new application builder
func NewBuilder() Builder {
	grammarAdapterBuilder := grammars.NewAdapterBuilder()
	resultBuilder := results.NewBuilder()
	resultTokenBuilder := results.NewTokenBuilder()
	resultBlockBuilder := results.NewBlockBuilder()
	resultLineBuilder := results.NewLineBuilder()
	resultElementWithCardinalityBuilder := results.NewElementWithCardinalityBuilder()
	resultElementBuilder := results.NewElementBuilder()
	return createBuilder(
		grammarAdapterBuilder,
		resultBuilder,
		resultTokenBuilder,
		resultBlockBuilder,
		resultLineBuilder,
		resultElementWithCardinalityBuilder,
		resultElementBuilder,
	)
}

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithExternals(externals grammars.Externals) Builder
	Now() (Application, error)
}

// Application represents the grammar application
type Application interface {
	Compile(script string) (grammars.Grammar, error)
	Execute(grammar grammars.Grammar, data []byte, canHavePrefix bool) (results.Result, error)
}
