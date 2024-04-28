package user

import (
	"gorm.io/gorm"
	"zhiqu/model"
)

type Dao struct {
	db *gorm.DB
}

func NewDao(db *gorm.DB) *Dao {
	return &Dao{db}
}

func (d Dao) GetUserById(id int64) (*model.User, error) {
	v := &model.User{ID: id}
	if err := d.db.First(v, id).Error; err != nil {
		return nil, err
	}
	return v, nil
}
