package records

type Account struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}
