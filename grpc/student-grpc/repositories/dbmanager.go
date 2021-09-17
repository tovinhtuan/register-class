package repositories

import (
	"context"
	"regis_class_go/grpc/student-grpc/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type StudentRepository interface {
	CreateStudent(ctx context.Context, model *models.Student) (*models.Student, error)
	ReadStudentByMssv(ctx context.Context,mssv string)(*models.Student, error)
	UpdateStudent(ctx context.Context, model *models.Student)(*models.Student, error)
	ChangePassword(ctx context.Context, model *models.Student)(*models.Student, error)
}
type dbmanager struct {
	*gorm.DB
}
func NewDBManager()(StudentRepository, error){
	db, err := gorm.Open(postgres.Open("host=localhost user=admin password=admin dbname=student port=5433 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"))
	if err != nil {
		return nil, err
	}
	err1 := db.AutoMigrate(
		&models.Student{},
	)
	if err1 != nil {
		return nil, err
	}
	return &dbmanager{db.Debug()}, nil
}
func (m *dbmanager) CreateStudent(ctx context.Context, model *models.Student) (*models.Student, error){
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}
func(m *dbmanager)UpdateStudent(ctx context.Context, model *models.Student)(*models.Student, error){
	if err := m.Save(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}
func (m *dbmanager)ReadStudentByMssv(ctx context.Context, mssv string)(*models.Student, error){
	student := models.Student{}
	if err := m.Where(&models.Student{Mssv: mssv}).First(&student).Error; err != nil {
		return nil, err
	}
	return &student, nil
}
func (m *dbmanager) ChangePassword(ctx context.Context, model *models.Student)(*models.Student, error){
	if err := m.Save(model).Error; err != nil{
		return nil, err 
	}
	return model, nil
}
