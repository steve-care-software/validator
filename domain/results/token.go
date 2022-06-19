package results

type token struct {
	name     string
	block    Block
	channels uint
}

func createToken(
	name string,
	block Block,
	channels uint,
) Token {
	out := token{
		name:     name,
		block:    block,
		channels: channels,
	}

	return &out
}

// Name returns the name
func (obj *token) Name() string {
	return obj.name
}

// Block returns the block
func (obj *token) Block() Block {
	return obj.block
}

// Channels returns the amount of channels
func (obj *token) Channels() uint {
	return obj.channels
}

// Path retruns the path
func (obj *token) Path() []string {
	path := []string{
		obj.Name(),
	}

	if !obj.Block().HasMatch() {
		return path
	}

	matches := obj.Block().Match().Elements()
	for _, oneMatch := range matches {
		if !oneMatch.HasMatches() {
			continue
		}

		path = append(path, oneMatch.Path()...)
	}

	return path
}

// IsSuccess returns true if successful, false otherwise
func (obj *token) IsSuccess() bool {
	return obj.block.IsSuccess()
}
