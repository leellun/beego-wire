package models

import "time"

type Moment struct {
	Id_RENAME  string    `orm:"column(id);null"`
	Title      string    `orm:"column(title);null"`
	Content    string    `orm:"column(content);null"`
	Vid        string    `orm:"column(vid);null"`
	Userid     string    `orm:"column(userid);null"`
	Cid        string    `orm:"column(cid);null"`
	Createtime time.Time `orm:"column(createtime);type(timestamp without time zone);null"`
	Updatetime time.Time `orm:"column(updatetime);type(timestamp without time zone);null"`
}
