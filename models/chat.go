package models

import "time"

type Chat struct {
	Id_RENAME     int64     `orm:"column(id);null"`
	Content       string    `orm:"column(content);null"`
	SendUserId    int64     `orm:"column(send_user_id);null"`
	IsOnline      int16     `orm:"column(is_online);null"`
	ReceiveUserId string    `orm:"column(receive_user_id);null"`
	CreateTime    time.Time `orm:"column(create_time);type(timestamp without time zone);null"`
	UpdateTime    time.Time `orm:"column(update_time);type(timestamp without time zone);null"`
}
