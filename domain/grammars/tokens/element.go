package tokens

type element struct {
	pByte     *byte
	token     Token
	reference string
}

func createElementWithByte(
	pByte *byte,
) Element {
	return createElementInternally(pByte, nil, "")
}

func createElementWithToken(
	token Token,
) Element {
	return createElementInternally(nil, token, "")
}

func createElementWithReference(
	reference string,
) Element {
	return createElementInternally(nil, nil, reference)
}

func createElementInternally(
	pByte *byte,
	token Token,
	reference string,
) Element {
	out := element{
		pByte:     pByte,
		token:     token,
		reference: reference,
	}

	return &out
}

// IsByte returns true if there is a byte, false otherwise
func (obj *element) IsByte() bool {
	return obj.pByte != nil
}

// Byte returns the byte, if any
func (obj *element) Byte() *byte {
	return obj.pByte
}

// IsToken returns true if there is a token, false otherwise
func (obj *element) IsToken() bool {
	return obj.token != nil
}

// Token returns the token, if any
func (obj *element) Token() Token {
	return obj.token
}

// IsReference returns true if there is a reference token, false otherwise
func (obj *element) IsReference() bool {
	return obj.reference != ""
}

// Reference returns the reference, if any
func (obj *element) Reference() string {
	return obj.reference
}
