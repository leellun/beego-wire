package models

type MediaSubtitle struct {
	Id_RENAME    int    `orm:"column(id)"`
	SubtitleFile string `orm:"column(subtitle_file)"`
	LanguageId   int    `orm:"column(language_id)"`
	MediaId      int    `orm:"column(media_id)"`
	UserId       int    `orm:"column(user_id)"`
}
