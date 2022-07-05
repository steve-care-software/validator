package applications

import (
	"github.com/steve-care-software/validator/domain/grammars"
	"github.com/steve-care-software/validator/domain/results"
)

type builder struct {
	grammarAdapterBuilder               grammars.AdapterBuilder
	resultBuilder                       results.Builder
	resultTokenBuilder                  results.TokenBuilder
	resultBlockBuilder                  results.BlockBuilder
	resultLineBuilder                   results.LineBuilder
	resultElementWithCardinalityBuilder results.ElementWithCardinalityBuilder
	resultElementBuilder                results.ElementBuilder
	externals                           grammars.Externals
}

func createBuilder(
	grammarAdapterBuilder grammars.AdapterBuilder,
	resultBuilder results.Builder,
	resultTokenBuilder results.TokenBuilder,
	resultBlockBuilder results.BlockBuilder,
	resultLineBuilder results.LineBuilder,
	resultElementWithCardinalityBuilder results.ElementWithCardinalityBuilder,
	resultElementBuilder results.ElementBuilder,
) Builder {
	out := builder{
		grammarAdapterBuilder:               grammarAdapterBuilder,
		resultBuilder:                       resultBuilder,
		resultTokenBuilder:                  resultTokenBuilder,
		resultBlockBuilder:                  resultBlockBuilder,
		resultLineBuilder:                   resultLineBuilder,
		resultElementWithCardinalityBuilder: resultElementWithCardinalityBuilder,
		resultElementBuilder:                resultElementBuilder,
		externals:                           nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.grammarAdapterBuilder,
		app.resultBuilder,
		app.resultTokenBuilder,
		app.resultBlockBuilder,
		app.resultLineBuilder,
		app.resultElementWithCardinalityBuilder,
		app.resultElementBuilder,
	)
}

// WithExternals add externals to the builder
func (app *builder) WithExternals(externals grammars.Externals) Builder {
	app.externals = externals
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	grammarAdapterBuilder := app.grammarAdapterBuilder.Create()
	if app.externals != nil {
		grammarAdapterBuilder.WithExternals(app.externals)
	}

	grammarAdapter, err := grammarAdapterBuilder.Now()
	if err != nil {
		return nil, err
	}

	return createApplication(
		grammarAdapter,
		app.resultBuilder,
		app.resultTokenBuilder,
		app.resultBlockBuilder,
		app.resultLineBuilder,
		app.resultElementWithCardinalityBuilder,
		app.resultElementBuilder,
	), nil
}
