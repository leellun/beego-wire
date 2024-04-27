package models

type Tag struct {
	Id_RENAME         int    `orm:"column(id)"`
	Title             string `orm:"column(title)"`
	MediaCount        int    `orm:"column(media_count)"`
	ListingsThumbnail string `orm:"column(listings_thumbnail);null"`
	UserId            int    `orm:"column(user_id);null"`
}
