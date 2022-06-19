package tokens

import "errors"

type elementBuilder struct {
	pByte     *byte
	token     Token
	reference string
}

func createElementBuilder() ElementBuilder {
	out := elementBuilder{
		pByte:     nil,
		token:     nil,
		reference: "",
	}

	return &out
}

// Create initializes the builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder()
}

// WithByte adds a byte value to the builder
func (app *elementBuilder) WithByte(byteValue byte) ElementBuilder {
	app.pByte = &byteValue
	return app
}

// WithToken adds a token to the builder
func (app *elementBuilder) WithToken(token Token) ElementBuilder {
	app.token = token
	return app
}

// WithReference adds a token reference to the builder
func (app *elementBuilder) WithReference(reference string) ElementBuilder {
	app.reference = reference
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.pByte != nil {
		return createElementWithByte(app.pByte), nil
	}

	if app.token != nil {
		return createElementWithToken(app.token), nil
	}

	if app.reference != "" {
		return createElementWithReference(app.reference), nil
	}

	return nil, errors.New("the Element is invalid")
}
