package cardinality

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
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
