package login

import "zhiqu/model"

type Service struct {
	dao *Dao
}

func NewService(dao *Dao) *Service {
	return &Service{dao}
}
func (s Service) Login(v Req) *model.User {
	return nil
}
