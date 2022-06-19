package applications

import "github.com/steve-care-software/validator/domain/grammars/tokens"

type tokenData struct {
	token tokens.Token
	data  []byte
}

func createTokenData(
	token tokens.Token,
	data []byte,
) *tokenData {
	out := tokenData{
		token: token,
		data:  data,
	}

	return &out
}

// Token returns the token
func (obj *tokenData) Token() tokens.Token {
	return obj.token
}

// Data returns the data
func (obj *tokenData) Data() []byte {
	return obj.data
}
