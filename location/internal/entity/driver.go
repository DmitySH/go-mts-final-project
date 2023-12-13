package entity

type Driver struct {
	Id   string `db:"id"`
	Name string `db:"name"`
	Auto string `db:"auto"`
}
