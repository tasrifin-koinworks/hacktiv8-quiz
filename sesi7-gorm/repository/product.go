package repository

import (
	"sesi7-gorm/models"

	"gorm.io/gorm"
)

type ProductRepo interface {
	CreateProduct(*models.Product) error
	GetAllProducts() (*[]models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
	UpdateProductByID(id uint, name string, brand string) (*models.Product, error)
	DeleteProduct(id uint) error
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) CreateProduct(request *models.Product) error {
	err := r.db.Create(request).Error
	return err
}

func (r *productRepo) GetAllProducts() (*[]models.Product, error) {
	var product []models.Product

	err := r.db.Find(&product).Error
	return &product, err
}

func (r *productRepo) GetProductByID(id uint) (*models.Product, error) {
	var product models.Product

	err := r.db.First(&product, "id=?", id).Error
	return &product, err
}

func (r *productRepo) UpdateProductByID(id uint, name string, brand string) (*models.Product, error) {
	var product models.Product

	err := r.db.Model(&product).Where("id=?", id).Updates(models.Product{Name: name, Brand: brand}).Error
	return &product, err
}

func (r *productRepo) DeleteProduct(id uint) error {
	var product models.Product

	err := r.db.Where("id=?", id).Delete(&product).Error
	return err
}
