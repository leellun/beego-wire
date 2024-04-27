package models

import "time"

type Carousel struct {
	Id_RENAME   int64     `orm:"column(id);null"`
	Title       string    `orm:"column(title);null"`
	Description string    `orm:"column(description);null"`
	FileId      int64     `orm:"column(file_id);null"`
	VideoId     int64     `orm:"column(video_id);null"`
	CreateTime  time.Time `orm:"column(createTime);type(timestamp without time zone);null"`
	UpdateTime  time.Time `orm:"column(updateTime);type(timestamp without time zone);null"`
}
