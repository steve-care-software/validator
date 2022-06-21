package applications

import (
	"github.com/steve-care-software/validator/domain/grammars"
	"github.com/steve-care-software/validator/domain/results"
)

// NewApplication creates a new application
func NewApplication() Application {
	grammarAdapter := grammars.NewAdapter()
	resultBuilder := results.NewBuilder()
	resultTokenBuilder := results.NewTokenBuilder()
	resultBlockBuilder := results.NewBlockBuilder()
	resultLineBuilder := results.NewLineBuilder()
	resultElementWithCardinalityBuilder := results.NewElementWithCardinalityBuilder()
	resultElementBuilder := results.NewElementBuilder()
	return createApplication(
		grammarAdapter,
		resultBuilder,
		resultTokenBuilder,
		resultBlockBuilder,
		resultLineBuilder,
		resultElementWithCardinalityBuilder,
		resultElementBuilder,
	)
}

// Application represents the grammar application
type Application interface {
	Compile(script string) (grammars.Grammar, error)
	Execute(script string, data []byte, canHavePrefix bool) (results.Result, error)
}
