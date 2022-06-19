package results

type line struct {
	index    uint
	elements []ElementWithCardinality
}

func createLine(
	index uint,
	elements []ElementWithCardinality,
) Line {
	out := line{
		index:    index,
		elements: elements,
	}

	return &out
}

// Discovered returns the amount of value discovered
func (obj *line) Discovered() uint {
	amount := uint(0)
	for _, oneElement := range obj.elements {
		amount += oneElement.Discovered()
	}

	return amount
}

// Index returns the index
func (obj *line) Index() uint {
	return obj.index
}

// Elements returns the elements
func (obj *line) Elements() []ElementWithCardinality {
	return obj.elements
}

// IsSuccess returns true if successful, false otherwise
func (obj *line) IsSuccess() bool {
	for _, oneMatch := range obj.elements {
		if !oneMatch.IsSuccess() {
			return false
		}
	}

	return true
}

// Channels returns the amount of channels
func (obj *line) Channels() uint {
	amount := uint(0)
	for _, oneElement := range obj.elements {
		amount += oneElement.Channels()
	}

	return amount
}

// Remaining returns the remaining content, if any
func (obj *line) Remaining() []byte {
	return obj.elements[len(obj.elements)-1].Remaining()
}
