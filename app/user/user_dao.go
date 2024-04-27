package user

import (
	"github.com/beego/beego/v2/client/orm"
	"zhiqu/models"
)

type Dao struct {
	ormer orm.Ormer
}

func NewDao(ormer orm.Ormer) *Dao {
	return &Dao{ormer}
}

func (d Dao) GetUserById(id int) (*models.User, error) {
	v := &models.User{Id: id}
	if err := d.ormer.Read(v); err != nil {
		return nil, err
	}
	return v, nil
}
