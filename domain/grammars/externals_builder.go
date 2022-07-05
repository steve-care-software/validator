package grammars

import (
	"errors"
	"fmt"
)

type externalsBuilder struct {
	list []External
}

func createExternalsBuilder() ExternalsBuilder {
	out := externalsBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *externalsBuilder) Create() ExternalsBuilder {
	return createExternalsBuilder()
}

// WithList adds a list to the builder
func (app *externalsBuilder) WithList(list []External) ExternalsBuilder {
	app.list = list
	return app
}

// Now builds a new Externals instance
func (app *externalsBuilder) Now() (Externals, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 External in order to build an Externals instance")
	}

	mp := map[string]External{}
	for _, oneExternal := range app.list {
		keyname := oneExternal.Token()
		if _, ok := mp[keyname]; ok {
			str := fmt.Sprintf("the external token (name: %s) is duplicate", keyname)
			return nil, errors.New(str)
		}

		mp[keyname] = oneExternal
	}

	return createExternals(app.list, mp), nil
}
