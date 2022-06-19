package grammars

import (
	"github.com/steve-care-software/validator/domain/grammars/channels"
	"github.com/steve-care-software/validator/domain/grammars/tokens"
)

type grammar struct {
	root     tokens.Token
	channels channels.Channels
}

func createGrammar(
	root tokens.Token,
) Grammar {
	return createGrammarInternally(root, nil)
}

func createGrammarWithChannels(
	root tokens.Token,
	channels channels.Channels,
) Grammar {
	return createGrammarInternally(root, channels)
}

func createGrammarInternally(
	root tokens.Token,
	channels channels.Channels,
) Grammar {
	out := grammar{
		root:     root,
		channels: channels,
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
