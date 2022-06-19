package results

type result struct {
	index uint
	token Token
}

func createResult(
	index uint,
	token Token,
) Result {
	out := result{
		index: index,
		token: token,
	}

	return &out
}

// Index returns the index
func (obj *result) Index() uint {
	return obj.index
}

// Cursor returns the cursor
func (obj *result) Cursor() uint {
	if obj.Token().Block().HasMatch() {
		input := obj.Token().Block().Input()
		inputWithIndex := uint(len(input)) + obj.index + obj.Token().Channels()
		elements := obj.Token().Block().Match().Elements()
		index := len(elements) - 1
		return inputWithIndex - uint(len(elements[index].Remaining()))
	}

	return 0
}

// Token returns the token
func (obj *result) Token() Token {
	return obj.token
}

// IsSuccess returns true if successful, false otherwise
func (obj *result) IsSuccess() bool {
	return obj.Token().Block().HasMatch()
}
