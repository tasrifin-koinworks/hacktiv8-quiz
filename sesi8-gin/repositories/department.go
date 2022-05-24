package repositories

import (
	"sesi8-gin/models"

	"github.com/jinzhu/gorm"
)

type DepartmentRepo interface {
	CreateDepartment(department *models.Department) error
	GetAllDepartments() (*[]models.Department, error)
}

type departmentRepo struct {
	db *gorm.DB
}

func NewDepartmentRepo(db *gorm.DB) DepartmentRepo {
	return &departmentRepo{db}
}

func (d *departmentRepo) CreateDepartment(department *models.Department) error {
	return d.db.Create(department).Error
}

func (d *departmentRepo) GetAllDepartments() (*[]models.Department, error) {
	var department []models.Department

	err := d.db.Find(&department).Error
	return &department, err
}
