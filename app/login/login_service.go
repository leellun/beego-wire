package login

type Service struct {
	dao *Dao
}

func NewService(dao *Dao) *Service {
	return &Service{dao}
}
