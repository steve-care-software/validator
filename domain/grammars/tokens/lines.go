package tokens

type lines struct {
	list []Line
}

func createLines(
	list []Line,
) Lines {
	out := lines{
		list: list,
	}

	return &out
}

// List returns the Line list
func (obj *lines) List() []Line {
	return obj.list
}
