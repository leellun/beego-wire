package login

import "zhiqu/models"

type Service struct {
	dao *Dao
}

func NewService(dao *Dao) *Service {
	return &Service{dao}
}
func (s Service) Login(v Req) *models.User {
	return nil
}
