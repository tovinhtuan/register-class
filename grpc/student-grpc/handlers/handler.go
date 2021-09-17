package handlers

import (
	"context"
	"database/sql"
	"regis_class_go/grpc/student-grpc/models"
	"regis_class_go/grpc/student-grpc/repositories"
	"regis_class_go/pb"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StudentHandler struct {
	pb.UnimplementedHUSTStudentServer

	StudentRepository repositories.StudentRepository
}

func NewStudentHandler(studentRepository repositories.StudentRepository) (*StudentHandler, error) {
	return &StudentHandler{
		StudentRepository: studentRepository,
	}, nil
}
func (h *StudentHandler) CreateStudent(ctx context.Context, in *pb.Student) (*pb.Student, error) {
	checkStudent, _ := h.StudentRepository.ReadStudentByMssv(ctx, in.Mssv)
	if checkStudent != nil {
		return nil, status.Error(codes.AlreadyExists, "mssv is exist")
	}
	sRequest := &models.Student{
		Id:          uuid.New(),
		Name:        in.Name,
		Class:       in.Class,
		Address:     in.Address,
		PhoneNumber: in.PhoneNumber,
		Email:       in.Email,
		Password:    in.Password,
		Mssv:        in.Mssv,
	}
	student, err := h.StudentRepository.CreateStudent(ctx, sRequest)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	sResponse := &pb.Student{
		Id:          student.Id.String(),
		Name:        student.Name,
		Class:       student.Class,
		Address:     student.Address,
		PhoneNumber: student.PhoneNumber,
		Email:       student.Email,
		Password:    student.Password,
		Mssv:        student.Mssv,
	}
	return sResponse, nil
}
func (h *StudentHandler) UpdateStudent(ctx context.Context, in *pb.UpdateStudentRequest) (*pb.Student, error) {
	student, err := h.StudentRepository.ReadStudentByMssv(ctx, in.Mssv)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	if in.Name != "" {
		student.Name = in.Name
	}
	if in.Class != "" {
		student.Class = in.Class
	}
	if in.Address != "" {
		student.Address = in.Address
	}
	if in.PhoneNumber != "" {
		student.PhoneNumber = in.PhoneNumber
	}
	if in.Email != "" {
		student.Email = in.Email
	}
	newStudent, err := h.StudentRepository.UpdateStudent(ctx, student)
	if err != nil {
		return nil, err
	}
	sResponse := &pb.Student{
		Id:          newStudent.Id.String(),
		Name:        newStudent.Name,
		Class:       newStudent.Class,
		Address:     newStudent.Address,
		PhoneNumber: newStudent.PhoneNumber,
		Email:       newStudent.Email,
		Password:    newStudent.Password,
		Mssv:        newStudent.Mssv,
	}
	return sResponse, nil
}
func (h *StudentHandler) ChangePassword(ctx context.Context, in *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	student, err := h.StudentRepository.ReadStudentByMssv(ctx, in.Mssv)
	if err != nil {
		return nil, status.Error(codes.NotFound, "mssv not exist")
	}
	if student.Name != in.Name {
		return nil, status.Error(codes.NotFound, "name student incorrect")
	}
	if student.Password != in.OldPassword {
		return nil, status.Error(codes.NotFound, "Password incorrect")
	}
	if in.NewPassword == in.OldPassword {
		return nil, status.Error(codes.AlreadyExists, "Password is exist")
	}
	if in.NewPassword == "" {
		return nil, status.Error(codes.AlreadyExists, "Field NewPassword not null")
	}
	student.Password = in.NewPassword
	_, err1 := h.StudentRepository.ChangePassword(ctx, student)
	if err1 != nil {
		return nil, err1
	}

	sResponse := &pb.ChangePasswordResponse{
		SuccessChangePassword: true,
	}
	return sResponse, nil
}
func (h *StudentHandler) FindStudent(ctx context.Context, in *pb.FindStudentRequest) (*pb.FindStudentResponse, error) {
	student, err := h.StudentRepository.ReadStudentByMssv(ctx, in.Mssv)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	cResponse := &pb.FindStudentResponse{
		Mssv:        student.Mssv,
		Name:        student.Name,
		Class:       student.Class,
		PhoneNumber: student.PhoneNumber,
		Email:       student.Email,
		Address:     student.Address,
	}
	return cResponse, nil
}
	
