package models

import "time"

type Collect struct {
	Id_RENAME   int64     `orm:"column(id);null"`
	Name        string    `orm:"column(name);null"`
	Cover       string    `orm:"column(cover);null"`
	Description string    `orm:"column(description);null"`
	UserId      int64     `orm:"column(user_id);null"`
	CreateTime  time.Time `orm:"column(create_time);type(timestamp without time zone);null"`
	UpdateTime  time.Time `orm:"column(update_time);type(timestamp without time zone);null"`
}
