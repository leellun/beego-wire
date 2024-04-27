package models

import "time"

type Danmu struct {
	Id_RENAME  int64     `orm:"column(id);null"`
	Text       string    `orm:"column(text);null"`
	Time       int       `orm:"column(time);null"`
	VideoId    int64     `orm:"column(video_id);null"`
	UserId     int64     `orm:"column(user_id);null"`
	Createtime time.Time `orm:"column(createtime);type(timestamp without time zone);null"`
	Updatetime time.Time `orm:"column(updatetime);type(timestamp without time zone);null"`
}
