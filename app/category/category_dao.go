package category

import "github.com/beego/beego/v2/client/orm"

type Dao struct {
	ormer orm.Ormer
}

func NewDao(ormer orm.Ormer) *Dao {
	return &Dao{ormer}
}
