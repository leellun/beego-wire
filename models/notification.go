package models

type Notification struct {
	Id_RENAME int64  `orm:"column(id)"`
	Action    string `orm:"column(action)"`
	Method    string `orm:"column(method)"`
	UserId    int64  `orm:"column(user_id)"`
	Notify    int16  `orm:"column(notify)"`
}
