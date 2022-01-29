package numberdb

type phoneNumber struct {
	Id     int    `db:"id"`
	Number string `db:"number"`
}
