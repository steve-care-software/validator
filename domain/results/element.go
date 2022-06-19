package results

type element struct {
	pValue *byte
	token  Token
}

func createElementWithValue(
	pValue *byte,
) Element {
	return createElementInternally(pValue, nil)
}

func createElementWithToken(
	token Token,
) Element {
	return createElementInternally(nil, token)
}

func createElementInternally(
	pValue *byte,
	token Token,
) Element {
	out := element{
		pValue: pValue,
		token:  token,
	}

	return &out
}

// Discovered returns the amount of value discovered
func (obj *element) Discovered() uint {
	if obj.IsValue() {
		return 1
	}

	return obj.Token().Block().Discovered()
}

// IsValue returns true if there is a value, false otherwise
func (obj *element) IsValue() bool {
	return obj.pValue != nil
}

// Value returns the value, if any
func (obj *element) Value() *byte {
	return obj.pValue
}

// IsToken returns true if there is a token, false otherwise
func (obj *element) IsToken() bool {
	return obj.token != nil
}

// Token returns the token, if any
func (obj *element) Token() Token {
	return obj.token
}

// IsSuccess returns true if successful, false otherwise
func (obj *element) IsSuccess() bool {
	if obj.IsValue() {
		return true
	}

	return obj.Token().IsSuccess()
}

// Channels returns the amount of channels
func (obj *element) Channels() uint {
	if obj.IsValue() {
		return 0
	}

	return obj.Token().Channels()
}
