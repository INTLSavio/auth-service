package repositories

import (
	"github.com/INTLSavio/auth-service/internal/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (userRepo *UserRepository) CreateUser(user *entities.User) error {
	return userRepo.DB.Create(user).Error
}

func (userRepo *UserRepository) FindUserByEmail(email string) (*entities.User, error) {
	var user *entities.User
	err := userRepo.DB.Where("email = $1", email).First(&user).Error

	return user, err
}
