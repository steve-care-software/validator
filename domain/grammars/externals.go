package grammars

import (
	"errors"
	"fmt"
)

type externals struct {
	list []External
	mp   map[string]External
}

func createExternals(
	list []External,
	mp map[string]External,
) Externals {
	out := externals{
		list: list,
		mp:   mp,
	}

	return &out
}

// List returns the list of externals
func (obj *externals) List() []External {
	return obj.list
}

// Find finds an external by name
func (obj *externals) Find(name string) (External, error) {
	if ins, ok := obj.mp[name]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the external token (name: %s) is undefined", name)
	return nil, errors.New(str)
}
