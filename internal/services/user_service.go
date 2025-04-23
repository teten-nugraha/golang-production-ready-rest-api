package services

import (
	"github.com/teten-nugraha/golang-production-ready-rest-api/internal/models"
	"github.com/teten-nugraha/golang-production-ready-rest-api/internal/repositories"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(id uint, user models.User) (*models.User, error)
	DeleteUser(id uint) error
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.FindAll()
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *userService) UpdateUser(id uint, user models.User) (*models.User, error) {
	return s.userRepo.Update(id, user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}
