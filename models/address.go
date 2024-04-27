package models

type Address struct {
	Id_RENAME int64  `orm:"column(id)"`
	Email     string `orm:"column(email)"`
	UserId    int64  `orm:"column(user_id)"`
	Primary   int16  `orm:"column(primary)"`
}
