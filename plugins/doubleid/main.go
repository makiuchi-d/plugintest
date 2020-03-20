package main

import (
	"github.com/makiuchi-d/plugintest/records"
)

type DoubleNameAccount struct {
	*records.Account
	DoubleId int `db:"double_id"`
}

func AppendCols(a *records.Account) (interface{}, []string) {
	return &DoubleNameAccount{a, a.Id * 2}, []string{"double_id"}
}
