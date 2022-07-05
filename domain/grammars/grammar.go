package grammars

import (
	"github.com/steve-care-software/validator/domain/grammars/channels"
	"github.com/steve-care-software/validator/domain/grammars/tokens"
)

type grammar struct {
	root      tokens.Token
	channels  channels.Channels
	externals Externals
}

func createGrammar(
	root tokens.Token,
) Grammar {
	return createGrammarInternally(root, nil, nil)
}

func createGrammarWithChannels(
	root tokens.Token,
	channels channels.Channels,
) Grammar {
	return createGrammarInternally(root, channels, nil)
}

func createGrammarWithExternals(
	root tokens.Token,
	externals Externals,
) Grammar {
	return createGrammarInternally(root, nil, externals)
}

func createGrammarWithChannelsAndExternals(
	root tokens.Token,
	channels channels.Channels,
	externals Externals,
) Grammar {
	return createGrammarInternally(root, channels, externals)
}

func createGrammarInternally(
	root tokens.Token,
	channels channels.Channels,
	externals Externals,
) Grammar {
	out := grammar{
		root:      root,
		channels:  channels,
		externals: externals,
	}

	return &out
}

// Root returns the root token
func (obj *grammar) Root() tokens.Token {
	return obj.root
}

// HasChannels returns true if there is channels, false otherwise
func (obj *grammar) HasChannels() bool {
	return obj.channels != nil
}

// Channels returs the chanels, if any
func (obj *grammar) Channels() channels.Channels {
	return obj.channels
}

// HasExternals returns true if there is externals, false otherwise
func (obj *grammar) HasExternals() bool {
	return obj.externals != nil
}

// Externals returs the externals, if any
func (obj *grammar) Externals() Externals {
	return obj.externals
}
