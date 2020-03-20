package plugintest

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"

	"github.com/makiuchi-d/plugintest/plugins"
	"github.com/makiuchi-d/plugintest/records"
)

var (
	cols []string
)

func init() {
	t := reflect.TypeOf(records.Account{})
	cols = make([]string, 0, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		if c := t.Field(i).Tag.Get("db"); c != "" {
			cols = append(cols, c)
		}
	}
}

type Repo struct {
	Db *sqlx.DB
	AC plugins.AppendCols
}

func (r *Repo) Insert(a *records.Account) error {
	var obj interface{} = a
	cols := cols
	if r.AC != nil {
		var addc []string
		obj, addc = r.AC(a)
		cols = append(cols, addc...)
	}
	query := fmt.Sprintf("INSERT INTO account (%s) VALUES (:%s)",
		strings.Join(cols, ","),
		strings.Join(cols, ",:"))

	fmt.Println("query: ", query)

	_, err := r.Db.NamedExec(query, obj)
	return err
}
