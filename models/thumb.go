package models

import "time"

type Thumb struct {
	Id_RENAME  int64     `orm:"column(id);null"`
	UserId     int64     `orm:"column(user_id);null"`
	CommentId  int64     `orm:"column(comment_id);null"`
	VideoId    int64     `orm:"column(video_id);null"`
	CreateTime time.Time `orm:"column(create_time);type(timestamp without time zone);null"`
	UpdateTime time.Time `orm:"column(update_time);type(timestamp without time zone);null"`
}
