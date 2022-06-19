package tokens

type token struct {
	name  string
	lines Lines
}

func createToken(
	name string,
	lines Lines,
) Token {
	out := token{
		name:  name,
		lines: lines,
	}

	return &out
}

// Name return the name
func (obj *token) Name() string {
	return obj.name
}

// Lines return the lines
func (obj *token) Lines() Lines {
	return obj.lines
}
