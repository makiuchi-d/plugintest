package main

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"github.com/makiuchi-d/plugintest"
	"github.com/makiuchi-d/plugintest/plugins"
	"github.com/makiuchi-d/plugintest/records"
)

var schema = `create table account (
  id   integer primary key,
  name varchar,
  double_name varchar default '-',
  double_id   integer default 0)`

type FullAccount struct {
	Id    int    `db:"id"`
	Name  string `db:"name"`
	DName string `db:"double_name"`
	DId   int    `db:"double_id"`
}

func main() {
	var err error
	var ac plugins.AppendCols
	if len(os.Args) > 1 {
		ac, err = plugins.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
	}

	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(schema)
	if err != nil {
		panic(err)
	}

	repo := plugintest.Repo{
		Db: db,
		AC: ac,
	}

	accounts := []*records.Account{
		{Id: 1, Name: "One"},
		{Id: 2, Name: "Two"},
		{Id: 3, Name: "Three"},
	}

	for _, a := range accounts {
		fmt.Println(a)

		err = repo.Insert(a)
		if err != nil {
			panic(err)
		}
	}

	fas := []FullAccount{}
	err = repo.Db.Select(&fas, "select * from account")
	if err != nil {
		panic(err)
	}
	fmt.Println(fas)
}
