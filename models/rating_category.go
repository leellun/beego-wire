package models

type RatingCategory struct {
	Id_RENAME   int    `orm:"column(id)"`
	Description string `orm:"column(description)"`
	Enabled     bool   `orm:"column(enabled)"`
	Title       string `orm:"column(title)"`
}
