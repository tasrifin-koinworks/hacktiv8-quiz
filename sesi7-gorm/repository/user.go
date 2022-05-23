package repository

import (
	"sesi7-gorm/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(*models.User) error
	GetAllUsers() (*[]models.User, error)
	GetUserByID(id uint) (*models.User, error)
	UpdateUserByID(id uint, email string) (*models.User, error)
	DeleteUser(id uint) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) CreateUser(request *models.User) error {
	err := r.db.Create(request).Error
	return err
}

func (r *userRepo) GetAllUsers() (*[]models.User, error) {
	var users []models.User

	err := r.db.Find(&users).Error
	return &users, err
}

func (r *userRepo) GetUserByID(userId uint) (*models.User, error) {
	var user models.User

	err := r.db.First(&user, "id=?", userId).Error
	return &user, err
}

func (r *userRepo) UpdateUserByID(userId uint, newEmail string) (*models.User, error) {
	var user models.User

	err := r.db.Model(&user).Where("id=?", userId).Updates(models.User{Email: newEmail}).Error
	return &user, err
}

func (r *userRepo) DeleteUser(userId uint) error {
	var user models.User

	err := r.db.Where("id=?", userId).Delete(&user).Error

	return err
}
