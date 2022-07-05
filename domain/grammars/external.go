package grammars

type external struct {
	token   string
	grammar Grammar
}

func createExternal(
	token string,
	grammar Grammar,
) External {
	out := external{
		token:   token,
		grammar: grammar,
	}

	return &out
}

// Token returns the token
func (obj *external) Token() string {
	return obj.token
}

// Grammar returns the grammar
func (obj *external) Grammar() Grammar {
	return obj.grammar
}
