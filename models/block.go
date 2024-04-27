package models

import "time"

type Block struct {
	Id_RENAME   int64     `orm:"column(id);null"`
	Name        string    `orm:"column(name);null"`
	Description string    `orm:"column(description);null"`
	Createtime  time.Time `orm:"column(createtime);type(timestamp without time zone);null"`
	Updatetime  time.Time `orm:"column(updatetime);type(timestamp without time zone);null"`
}
