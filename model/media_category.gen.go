// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameMediaCategory = "media_category"

// MediaCategory mapped from table <media_category>
type MediaCategory struct {
	ID         int64 `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:媒体分类中间表id" json:"id"` // 媒体分类中间表id
	MediaID    int64 `gorm:"column:media_id;type:bigint;not null;comment:媒体id" json:"media_id"`               // 媒体id
	CategoryID int64 `gorm:"column:category_id;type:bigint;not null;comment:分类id" json:"category_id"`         // 分类id
}

// TableName MediaCategory's table name
func (*MediaCategory) TableName() string {
	return TableNameMediaCategory
}
