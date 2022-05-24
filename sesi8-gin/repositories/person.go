package repositories

import (
	"sesi8-gin/models"

	"github.com/jinzhu/gorm"
)

type PersonRepo interface {
	CreatePerson(person *models.Person) error
	GetAllPersons() (*[]models.Person, error)
}

type personRepo struct {
	db *gorm.DB
}

func NewPersonRepo(db *gorm.DB) PersonRepo {
	return &personRepo{db}
}

func (p *personRepo) CreatePerson(person *models.Person) error {
	return p.db.Create(person).Error
}

func (p *personRepo) GetAllPersons() (*[]models.Person, error) {
	var person []models.Person
	err := p.db.Find(&person).Error
	return &person, err
}
