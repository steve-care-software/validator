package grammars

import "github.com/steve-care-software/validator/domain/grammars/cardinality"

type scriptToken struct {
	name  string
	lines []*scriptLine
}

type scriptLine struct {
	values []*scriptValue
}

type scriptValue struct {
	pByte       *byte
	tokenName   string
	cardinality cardinality.Cardinality
}
