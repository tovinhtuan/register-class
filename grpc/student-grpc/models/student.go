package models

import "github.com/google/uuid"

type Student struct {
	Id          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey`
	Name        string    `gorm:"type:varchar(256)"`
	Class       string    `gorm:"type:varchar(50)"`
	Address     string    `gorm:"type:varchar(256)"`
	PhoneNumber string    `gorm:"type:varchar(10)"`
	Email       string    `gorm:"type:varchar(256)"`
	Password    string    `gorm:"type:varchar(20)"`
	Mssv        string    `gorm:"type:varchar(8);not null"`
}
