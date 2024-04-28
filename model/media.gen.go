// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameMedium = "media"

// Medium mapped from table <media>
type Medium struct {
	ID             int64      `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:媒体文件的基本信息id" json:"id"`          // 媒体文件的基本信息id
	Description    string     `gorm:"column:description;type:text;not null;comment:描述信息" json:"description"`                      // 描述信息
	Dislikes       int32      `gorm:"column:dislikes;type:integer;not null;comment:不喜欢数量" json:"dislikes"`                        // 不喜欢数量
	Duration       int32      `gorm:"column:duration;type:integer;not null;comment:时长" json:"duration"`                           // 时长
	HlsFile        string     `gorm:"column:hls_file;type:character varying(1000);not null;comment:媒体的 HLS 文件路径" json:"hls_file"` // 媒体的 HLS 文件路径
	Likes          int32      `gorm:"column:likes;type:integer;not null;comment:表示媒体文件的喜欢数量" json:"likes"`                        // 表示媒体文件的喜欢数量
	Md5sum         *string    `gorm:"column:md5sum;type:character varying(50);comment:媒体文件的 MD5 校验和" json:"md5sum"`               // 媒体文件的 MD5 校验和
	MediaInfo      string     `gorm:"column:media_info;type:text;not null;comment:媒体文件的信息" json:"media_info"`                     // 媒体文件的信息
	MediaType      string     `gorm:"column:media_type;type:character varying(20);not null;comment:媒体文件的类型" json:"media_type"`    // 媒体文件的类型
	Poster         string     `gorm:"column:poster;type:character varying(500);not null;comment:媒体文件的封面图" json:"poster"`          // 媒体文件的封面图
	Size           *string    `gorm:"column:size;type:character varying(20);comment:媒体文件大小" json:"size"`                          // 媒体文件大小
	Thumbnail      string     `gorm:"column:thumbnail;type:character varying(500);not null;comment:缩略图路径" json:"thumbnail"`       // 缩略图路径
	VideoHeight    int32      `gorm:"column:video_height;type:integer;not null;comment:视频高度" json:"video_height"`                 // 视频高度
	Views          int32      `gorm:"column:views;type:integer;not null;comment:媒体文件的观看次数" json:"views"`                          // 媒体文件的观看次数
	ChannelID      *int32     `gorm:"column:channel_id;type:integer;comment:媒体文件所属频道的 ID" json:"channel_id"`                      // 媒体文件所属频道的 ID
	UserID         int32      `gorm:"column:user_id;type:integer;not null;comment:媒体文件上传者的用户 ID" json:"user_id"`                  // 媒体文件上传者的用户 ID
	CreateTime     *time.Time `gorm:"column:create_time;type:timestamp(6) with time zone;comment:添加时间" json:"create_time"`        // 添加时间
	UpdateTime     time.Time  `gorm:"column:update_time;type:int unsigned;autoUpdateTime" json:"update_time"`                     // 更新时间
	Status         int16      `gorm:"column:status;type:smallint;not null;comment:编码状态" json:"status"`                            // 编码状态
	Name           string     `gorm:"column:name;type:character varying(100);not null;comment:媒体文件标题" json:"name"`                // 媒体文件标题
	Identification string     `gorm:"column:identification;type:uuid;not null;comment:媒体文件的唯一标识符" json:"identification"`          // 媒体文件的唯一标识符
}

// TableName Medium's table name
func (*Medium) TableName() string {
	return TableNameMedium
}