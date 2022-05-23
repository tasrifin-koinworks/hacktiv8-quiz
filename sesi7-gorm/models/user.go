package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"not null;unique;type:varchar(191)"`
	Products  []Product
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	fmt.Println("Before insert to table users")
	if len(u.Email) < 10 {
		err = fmt.Errorf("your email to short")
	}
	return
}

func (u *User) Print() {
	fmt.Println("ID \t\t:", u.ID)
	fmt.Println("Email \t\t:", u.Email)
	fmt.Println("----------------------------------")
}
