package entity

type Driver struct {
	Lat  float64 `db:"lat"`
	Lng  float64 `db:"lng"`
	Id   string  `db:"id"`
	Name string  `db:"name"`
	Auto string  `db:"auto"`
}
