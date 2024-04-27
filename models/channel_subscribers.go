package models

type ChannelSubscribers struct {
	Id_RENAME int64 `orm:"column(id)"`
	ChannelId int64 `orm:"column(channel_id)"`
	UserId    int64 `orm:"column(user_id)"`
}
