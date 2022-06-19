package results

import (
	"github.com/steve-care-software/validator/domain/grammars/tokens"
)

type elementWithCardinality struct {
	missing   uint
	element   tokens.Element
	matches   []Element
	remaining []byte
}

func createElementWithCardinality(
	missing uint,
	element tokens.Element,
	remaining []byte,
) ElementWithCardinality {
	return createElementWithCardinalityInternally(missing, element, nil, remaining)
}

func createElementWithCardinalityWithMatch(
	missing uint,
	element tokens.Element,
	remaining []byte,
	matches []Element,
) ElementWithCardinality {
	return createElementWithCardinalityInternally(missing, element, matches, remaining)
}

func createElementWithCardinalityInternally(
	missing uint,
	element tokens.Element,
	matches []Element,
	remaining []byte,
) ElementWithCardinality {
	out := elementWithCardinality{
		missing:   missing,
		element:   element,
		matches:   matches,
		remaining: remaining,
	}

	return &out
}

// Discovered returns the amount of value discovered
func (obj *elementWithCardinality) Discovered() uint {
	if obj.HasMatches() {
		return 0
	}

	amount := uint(0)
	for _, oneMatch := range obj.matches {
		amount += oneMatch.Discovered()
	}

	return amount
}

// Amount returns the amount of matches
func (obj *elementWithCardinality) Amount() uint {
	if obj.HasMatches() {
		return 0
	}

	return uint(len(obj.matches))
}

// Missing returns the amount of missing occurences
func (obj *elementWithCardinality) Missing() uint {
	return obj.missing
}

// Element returns the element
func (obj *elementWithCardinality) Element() tokens.Element {
	return obj.element
}

// Path retruns the path
func (obj *elementWithCardinality) Path() []string {
	if !obj.HasMatches() {
		return nil
	}

	path := []string{}
	matches := obj.Matches()
	for _, oneMatch := range matches {
		if oneMatch.IsValue() {
			continue
		}

		path = append(path, oneMatch.Token().Path()...)
	}

	return path
}

// Remaining returns the remaining content, if any
func (obj *elementWithCardinality) Remaining() []byte {
	return obj.remaining
}

// HasMatches returns true if there is matches, false otherwise
func (obj *elementWithCardinality) HasMatches() bool {
	return obj.matches != nil
}

// Matches returns the matches, if any
func (obj *elementWithCardinality) Matches() []Element {
	return obj.matches
}

// IsSuccess returns true if successful, false otherwise
func (obj *elementWithCardinality) IsSuccess() bool {
	if obj.Missing() > 0 {
		return false
	}

	if !obj.HasMatches() {
		return false
	}

	for _, oneMatch := range obj.matches {
		if !oneMatch.IsSuccess() {
			return false
		}
	}

	return true
}

// Channels returns the amount of channels
func (obj *elementWithCardinality) Channels() uint {
	amount := uint(0)
	for _, oneMatch := range obj.matches {
		amount += oneMatch.Channels()
	}

	return amount
}
