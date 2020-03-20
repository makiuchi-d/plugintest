package plugins

import (
	"fmt"
	"plugin"

	"github.com/makiuchi-d/plugintest/records"
)

type AppendCols = func(a *records.Account) (account interface{}, additionalCols []string)

func Open(path string) (AppendCols, error) {
	p, err := plugin.Open(path)
	if err != nil {
		return nil, err
	}

	sym, err := p.Lookup("AppendCols")
	if err != nil {
		return nil, err
	}

	f, ok := sym.(AppendCols)
	if !ok {
		return nil, fmt.Errorf("type AppendCols mismatch")
	}

	return f, nil
}
