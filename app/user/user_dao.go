package user

import (
	"github.com/beego/beego/v2/client/orm"
	"zhiqu/model"
)

type Dao struct {
	ormer orm.Ormer
}

func NewDao(ormer orm.Ormer) *Dao {
	return &Dao{ormer}
}

func (d Dao) GetUserById(id int) (*model.User, error) {
	v := &model.User{Id: id}
	if err := d.ormer.Read(v); err != nil {
		return nil, err
	}
	return v, nil
}
