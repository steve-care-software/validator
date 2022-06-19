package tokens

type line struct {
	list []ElementWithCardinality
}

func createLine(
	list []ElementWithCardinality,
) Line {
	out := line{
		list: list,
	}

	return &out
}

// List returns the elementWithCardinality list
func (obj *line) List() []ElementWithCardinality {
	return obj.list
}
