package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null;type:varchar(191)"`
	Brand     string `gorm:"not null;type:varchar(191)"`
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Product) BeforeCreateProduct(db *gorm.DB) (err error) {
	fmt.Println("Before insert to table Product")
	if len(p.Name) < 10 {
		err = fmt.Errorf("Product name to short")
	}
	return
}

func (p *Product) PrintProduct() {
	fmt.Println("ID \t\t:", p.ID)
	fmt.Println("Name \t\t:", p.Name)
	fmt.Println("Brand \t\t:", p.Brand)
	fmt.Println("----------------------------------")
}
