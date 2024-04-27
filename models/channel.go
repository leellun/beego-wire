package models

import "time"

type Channel struct {
	Id_RENAME   int       `orm:"column(id)"`
	Title       string    `orm:"column(title)"`
	Description string    `orm:"column(description)"`
	CreateTime  time.Time `orm:"column(create_time);type(timestamp with time zone)"`
	BannerLogo  string    `orm:"column(banner_logo)"`
	UserId      int       `orm:"column(user_id)"`
	UpdateTime  time.Time `orm:"column(update_time);type(timestamp without time zone);null"`
}
