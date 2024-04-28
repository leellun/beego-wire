// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameRatingCategory = "rating_category"

// RatingCategory mapped from table <rating_category>
type RatingCategory struct {
	ID          int32  `gorm:"column:id;type:integer;not null;comment:评分分类信息id" json:"id"`                     // 评分分类信息id
	Description string `gorm:"column:description;type:text;not null;comment:描述" json:"description"`            // 描述
	Enabled     bool   `gorm:"column:enabled;type:boolean;not null;comment:评分类别是否启用" json:"enabled"`           // 评分类别是否启用
	Title       string `gorm:"column:title;type:character varying(200);not null;comment:评分类别的标题" json:"title"` // 评分类别的标题
}

// TableName RatingCategory's table name
func (*RatingCategory) TableName() string {
	return TableNameRatingCategory
}