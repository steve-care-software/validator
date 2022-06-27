package cardinality

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	builder := NewBuilder()
	nonZeroMultiple := []byte("+")[0]
	zeroMultiple := []byte("*")[0]
	optional := []byte("?")[0]
	rangePrefix := []byte("[")[0]
	rangeSuffix := []byte("]")[0]
	rangeSeparator := []byte(",")[0]
	numbersCharacters := []byte("0123456789")
	return createAdapter(
		builder,
		nonZeroMultiple,
		zeroMultiple,
		optional,
		rangePrefix,
		rangeSuffix,
		rangeSeparator,
		numbersCharacters,
	)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents a cardinality adapter
type Adapter interface {
	ToCardinality(script string) (Cardinality, []byte, error)
}

// Builder represents the cardinality builder
type Builder interface {
	Create() Builder
	WithMinimum(min uint8) Builder
	WithMaximum(max uint8) Builder
	Now() (Cardinality, error)
}

// Cardinality represents the cardinality
type Cardinality interface {
	Min() uint8
	HasMax() bool
	Max() *uint8
}
