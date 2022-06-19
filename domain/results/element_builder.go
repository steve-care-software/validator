package results

import "errors"

type elementBuilder struct {
	pValue *byte
	token  Token
}

func createElementBuilder() ElementBuilder {
	out := elementBuilder{
		pValue: nil,
		token:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder()
}

// WithValue adds a value to the builder
func (app *elementBuilder) WithValue(value byte) ElementBuilder {
	app.pValue = &value
	return app
}

// WithToken adds a token to the builder
func (app *elementBuilder) WithToken(token Token) ElementBuilder {
	app.token = token
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.pValue != nil {
		return createElementWithValue(app.pValue), nil
	}

	if app.token != nil {
		return createElementWithToken(app.token), nil
	}

	return nil, errors.New("the Element is invalid")
}
