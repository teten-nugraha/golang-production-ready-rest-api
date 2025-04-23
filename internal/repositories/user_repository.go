package repositories

import (
	"github.com/teten-nugraha/golang-production-ready-rest-api/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user models.User) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindByID(id uint) (*models.User, error)
	FindAll() ([]models.User, error)
	Update(id uint, user models.User) (*models.User, error)
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func (r userRepository) Create(user models.User) (*models.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r userRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r userRepository) Update(id uint, updatedUser models.User) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	user.Name = updatedUser.Name
	user.Email = updatedUser.Email
	if updatedUser.Password != "" {
		user.Password = updatedUser.Password
	}

	if err := r.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r userRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}
