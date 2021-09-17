package repositories

import (
	"context"
	"regis_class_go/grpc/class-grpc/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ClassRepository interface {
	CreateClass(ctx context.Context, model *models.Class)(*models.Class, error)
}
type dbmanager struct {
	*gorm.DB
}
func NewDBManager()(ClassRepository, error){
	db, err := gorm.Open(postgres.Open("host=localhost user=admin password=admin dbname=class port=5433 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"))
	if err != nil {
		return nil, err
	}
	err1 := db.AutoMigrate(
		&models.Class{},
	)
	if err1 != nil {
		return nil, err
	}
	return &dbmanager{db.Debug()}, nil
}
func (db *dbmanager) CreateClass(ctx context.Context, model *models.Class)(*models.Class, error){
	if err := db.Create(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}