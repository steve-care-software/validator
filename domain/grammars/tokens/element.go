package tokens

type element struct {
	pByte     *byte
	token     Token
	reference string
	external  string
}

func createElementWithByte(
	pByte *byte,
) Element {
	return createElementInternally(pByte, nil, "", "")
}

func createElementWithToken(
	token Token,
) Element {
	return createElementInternally(nil, token, "", "")
}

func createElementWithReference(
	reference string,
) Element {
	return createElementInternally(nil, nil, reference, "")
}

func createElementWithExternal(
	external string,
) Element {
	return createElementInternally(nil, nil, "", external)
}

func createElementInternally(
	pByte *byte,
	token Token,
	reference string,
	external string,
) Element {
	out := element{
		pByte:     pByte,
		token:     token,
		reference: reference,
		external:  external,
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

// IsExternal returns true if there is an external token, false otherwise
func (obj *element) IsExternal() bool {
	return obj.external != ""
}

// External returns the external, if any
func (obj *element) External() string {
	return obj.external
}
