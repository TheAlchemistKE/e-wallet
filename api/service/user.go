package service

import (
	"blog/api/repository"
	"blog/models"
)

//UserService UserService struct
type UserService struct {
	repository repository.UserRepository
}

//NewUserService : returns the UserService struct instance
func NewUserService(r repository.UserRepository) UserService {
	return UserService{
		repository: r,
	}
}

//Save -> calls user repository save method
func (u UserService) Save(user models.User) error {
	return u.repository.Save(user)
}

//FindAll -> calls user repo find all method
func (u UserService) FindAll(user models.User, keyword string) (*[]models.User, int64, error) {
	return u.repository.FindAll(user, keyword)
}

// Update -> calls user repo update method
func (u UserService) Update(user models.User) error {
	return u.repository.Update(user)
}

// Delete -> calls post repo delete method
func (u UserService) Delete(id int64) error {
	var user models.User
	user.Id = int(id)
	return u.repository.Delete(user)
}

// Find -> calls post repo find method
func (u UserService) Find(user models.User) (models.User, error) {
	return u.repository.Find(user)
}
