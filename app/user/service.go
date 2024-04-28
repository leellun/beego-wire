package user

import (
	"zhiqu/infrastructure/extension/errox"
	"zhiqu/model"
)

type Service struct {
	dao *Dao
}

func NewService(dao *Dao) *Service {
	return &Service{dao}
}
func (s Service) GetUserById(id int64) model.User {
	user, err := s.dao.GetUserById(id)
	if err != nil {
		panic(errox.WithMessage("用户不存在"))
	}
	return *user
}
