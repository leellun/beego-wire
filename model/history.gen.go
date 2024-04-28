// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameHistory = "history"

// History mapped from table <history>
type History struct {
	ID         int64      `gorm:"column:id;type:bigint;primaryKey;comment:播放历史id" json:"id"`                              // 播放历史id
	MediaID    *int64     `gorm:"column:media_id;type:bigint;comment:视频id" json:"media_id"`                               // 视频id
	UserID     *int64     `gorm:"column:user_id;type:bigint;comment:用户id" json:"user_id"`                                 // 用户id
	Count      *int32     `gorm:"column:count;type:integer;comment:播放次数" json:"count"`                                    // 播放次数
	CreateTime *time.Time `gorm:"column:create_time;type:timestamp(6) without time zone;comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime *time.Time `gorm:"column:update_time;type:int unsigned;autoUpdateTime" json:"update_time"`                 // 更新时间
}

// TableName History's table name
func (*History) TableName() string {
	return TableNameHistory
}