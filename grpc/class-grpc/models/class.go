package models

import "github.com/google/uuid"

type Class struct {
	Id            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey`
	Name          string	`gorm:"type:varchar(250)"`
	ClassID       string	`gorm:"type:varchar(20);not null"`
	AvailableSlot int64		`gorm:"type:bigint"`
	DayOfWeek     string	`gorm:"type:varchar(20)"`
}
