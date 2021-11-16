package users

import (
	"gorm.io/gorm"
	"time"
)

type UserRepository interface {
	Save() error
}

type repository struct {
	db gorm.DB
}

type Users struct {
	ID        string `gorm:"primaryKey;"`
	Name  string `gorm:"type:varchar(255);unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewRepository(db gorm.DB) *repository {
	return &repository{db: db}
}

func NewUsers(ID string, name string) *Users {
	return &Users{ID: ID, Name: name}
}
func (d repository) Save(user Users) error {
	err := d.db.Debug().Create(&user).Error
	if err != nil {
		return err
	}
	return err
}