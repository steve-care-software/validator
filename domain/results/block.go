package results

type block struct {
	input []byte
	list  []Line
	match Line
}

func createBlock(
	input []byte,
	list []Line,
) Block {
	return createBlockInternally(input, list, nil)
}

func createBlockWithMatch(
	input []byte,
	list []Line,
	match Line,
) Block {
	return createBlockInternally(input, list, match)
}

func createBlockInternally(
	input []byte,
	list []Line,
	match Line,
) Block {
	out := block{
		input: input,
		list:  list,
		match: match,
	}

	return &out
}

// Discovered returns the amount of value discovered
func (obj *block) Discovered() uint {
	if obj.HasMatch() {
		return 0
	}

	return obj.Match().Discovered()
}

// Input returns the input
func (obj *block) Input() []byte {
	return obj.input
}

// List returns the list of lines
func (obj *block) List() []Line {
	return obj.list
}

// Remaining returns the remaining, if any
func (obj *block) Remaining() []byte {
	if !obj.HasMatch() {
		return obj.input
	}

	return obj.Match().Remaining()
}

// HasMatch returns true if there is a match, false otherwise
func (obj *block) HasMatch() bool {
	return obj.match != nil
}

// Match returns the match, if any
func (obj *block) Match() Line {
	return obj.match
}

// IsSuccess returns true if successful, false otherwise
func (obj *block) IsSuccess() bool {
	if !obj.HasMatch() {
		return false
	}

	return obj.match.IsSuccess()
}

// Channels returns the amount of channels
func (obj *block) Channels() uint {
	if obj.HasMatch() {
		return obj.Match().Channels()
	}

	return 0
}
