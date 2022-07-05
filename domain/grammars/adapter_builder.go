package grammars

import (
	"github.com/steve-care-software/validator/domain/grammars/cardinality"
	"github.com/steve-care-software/validator/domain/grammars/channels"
	"github.com/steve-care-software/validator/domain/grammars/tokens"
)

type adapterBuilder struct {
	grammarBuilder                Builder
	externalsBuilder              ExternalsBuilder
	channelsBuilder               channels.Builder
	channelBuilder                channels.ChannelBuilder
	conditionBuilder              channels.ConditionBuilder
	tokenBuilder                  tokens.TokenBuilder
	linesBuilder                  tokens.LinesBuilder
	lineBuilder                   tokens.LineBuilder
	elementWithCardinalityBuilder tokens.ElementWithCardinalityBuilder
	elementBuilder                tokens.ElementBuilder
	cardinalityAdapter            cardinality.Adapter
	rootPrefix                    byte
	rootSuffix                    byte
	channelPrefix                 byte
	channelSuffix                 byte
	channelConditionPrevious      byte
	channelConditionNext          byte
	channelConditionAnd           byte
	tokenNamePrefix               byte
	bytePrefix                    byte
	linesPrefix                   byte
	linesSuffix                   byte
	lineDelimiter                 byte
	commentPrefix                 byte
	commentSuffix                 byte
	tokenNameCharacters           []byte
	channelCharacters             []byte
	externals                     Externals
}

func createAdapterBuilder(
	grammarBuilder Builder,
	externalsBuilder ExternalsBuilder,
	channelsBuilder channels.Builder,
	channelBuilder channels.ChannelBuilder,
	conditionBuilder channels.ConditionBuilder,
	tokenBuilder tokens.TokenBuilder,
	linesBuilder tokens.LinesBuilder,
	lineBuilder tokens.LineBuilder,
	elementWithCardinalityBuilder tokens.ElementWithCardinalityBuilder,
	elementBuilder tokens.ElementBuilder,
	cardinalityAdapter cardinality.Adapter,
	rootPrefix byte,
	rootSuffix byte,
	channelPrefix byte,
	channelSuffix byte,
	channelConditionPrevious byte,
	channelConditionNext byte,
	channelConditionAnd byte,
	tokenNamePrefix byte,
	bytePrefix byte,
	linesPrefix byte,
	linesSuffix byte,
	lineDelimiter byte,
	commentPrefix byte,
	commentSuffix byte,
	tokenNameCharacters []byte,
	channelCharacters []byte,
) AdapterBuilder {
	out := adapterBuilder{
		grammarBuilder:                grammarBuilder,
		externalsBuilder:              externalsBuilder,
		channelsBuilder:               channelsBuilder,
		channelBuilder:                channelBuilder,
		conditionBuilder:              conditionBuilder,
		tokenBuilder:                  tokenBuilder,
		linesBuilder:                  linesBuilder,
		lineBuilder:                   lineBuilder,
		elementWithCardinalityBuilder: elementWithCardinalityBuilder,
		elementBuilder:                elementBuilder,
		cardinalityAdapter:            cardinalityAdapter,
		rootPrefix:                    rootPrefix,
		rootSuffix:                    rootSuffix,
		channelPrefix:                 channelPrefix,
		channelSuffix:                 channelSuffix,
		channelConditionPrevious:      channelConditionPrevious,
		channelConditionNext:          channelConditionNext,
		channelConditionAnd:           channelConditionAnd,
		tokenNamePrefix:               tokenNamePrefix,
		bytePrefix:                    bytePrefix,
		linesPrefix:                   linesPrefix,
		linesSuffix:                   linesSuffix,
		lineDelimiter:                 lineDelimiter,
		commentPrefix:                 commentPrefix,
		commentSuffix:                 commentSuffix,
		tokenNameCharacters:           tokenNameCharacters,
		channelCharacters:             channelCharacters,
		externals:                     nil,
	}

	return &out
}

// Create initializes the builder
func (app *adapterBuilder) Create() AdapterBuilder {
	return createAdapterBuilder(
		app.grammarBuilder,
		app.externalsBuilder,
		app.channelsBuilder,
		app.channelBuilder,
		app.conditionBuilder,
		app.tokenBuilder,
		app.linesBuilder,
		app.lineBuilder,
		app.elementWithCardinalityBuilder,
		app.elementBuilder,
		app.cardinalityAdapter,
		app.rootPrefix,
		app.rootSuffix,
		app.channelPrefix,
		app.channelSuffix,
		app.channelConditionPrevious,
		app.channelConditionNext,
		app.channelConditionAnd,
		app.tokenNamePrefix,
		app.bytePrefix,
		app.linesPrefix,
		app.linesSuffix,
		app.lineDelimiter,
		app.commentPrefix,
		app.commentSuffix,
		app.tokenNameCharacters,
		app.channelCharacters,
	)
}

// WithExternals add externals to the builder
func (app *adapterBuilder) WithExternals(externals Externals) AdapterBuilder {
	app.externals = externals
	return app
}

// Now builds a new Adapter instance
func (app *adapterBuilder) Now() (Adapter, error) {
	if app.externals != nil {
		return createAdapterWithExternals(
			app.grammarBuilder,
			app.externalsBuilder,
			app.channelsBuilder,
			app.channelBuilder,
			app.conditionBuilder,
			app.tokenBuilder,
			app.linesBuilder,
			app.lineBuilder,
			app.elementWithCardinalityBuilder,
			app.elementBuilder,
			app.cardinalityAdapter,
			app.rootPrefix,
			app.rootSuffix,
			app.channelPrefix,
			app.channelSuffix,
			app.channelConditionPrevious,
			app.channelConditionNext,
			app.channelConditionAnd,
			app.tokenNamePrefix,
			app.bytePrefix,
			app.linesPrefix,
			app.linesSuffix,
			app.lineDelimiter,
			app.commentPrefix,
			app.commentSuffix,
			app.tokenNameCharacters,
			app.channelCharacters,
			app.externals,
		), nil
	}

	return createAdapter(
		app.grammarBuilder,
		app.externalsBuilder,
		app.channelsBuilder,
		app.channelBuilder,
		app.conditionBuilder,
		app.tokenBuilder,
		app.linesBuilder,
		app.lineBuilder,
		app.elementWithCardinalityBuilder,
		app.elementBuilder,
		app.cardinalityAdapter,
		app.rootPrefix,
		app.rootSuffix,
		app.channelPrefix,
		app.channelSuffix,
		app.channelConditionPrevious,
		app.channelConditionNext,
		app.channelConditionAnd,
		app.tokenNamePrefix,
		app.bytePrefix,
		app.linesPrefix,
		app.linesSuffix,
		app.lineDelimiter,
		app.commentPrefix,
		app.commentSuffix,
		app.tokenNameCharacters,
		app.channelCharacters,
	), nil
}
