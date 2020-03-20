package main

import (
	"github.com/makiuchi-d/plugintest/records"
)

type DoubleNameAccount struct {
	*records.Account
	DoubleName string `db:"double_name"`
}

func AppendCols(a *records.Account) (interface{}, []string) {
	return &DoubleNameAccount{a, a.Name + a.Name}, []string{"double_name"}
}
