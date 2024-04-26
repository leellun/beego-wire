package service

import (
	"zhiqu/models"
)

type UsersUserService struct {
}

func (*UsersUserService) GetUser(username string) (models.UsersUser, error) {

	return models.UsersUser{}, nil
}
func (s *UsersUserService) GetViewUser(username string) any {
	var data map[string]any
	user, err := s.GetUser(username)
	if err != nil {
		panic(err)
	}
	data["user"] = user
	data["can_edit"] = true || user.IsMediacmsManager()
	data["can_delete"] = user.IsMediacmsManager()
	data["show_contact_form"] = user.AllowContact && user.IsMediacmsEditor()
	return nil
}
