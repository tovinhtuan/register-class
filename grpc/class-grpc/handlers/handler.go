package handlers

import (
	"context"
	"regis_class_go/grpc/class-grpc/models"
	"regis_class_go/grpc/class-grpc/repositories"
	"regis_class_go/pb"

	"github.com/google/uuid"
)

type ClassHandler struct {
	pb.UnimplementedHUSTClassServer
	ClassRepository repositories.ClassRepository
}

func NewClassHandler(classRepository repositories.ClassRepository)(*ClassHandler, error){
	return &ClassHandler{
		ClassRepository:              classRepository,
	}, nil
}

func (h *ClassHandler)CreateClass(ctx context.Context, in *pb.Class)(*pb.Class, error){
	cRequest := &models.Class{
		Id:            uuid.New(),
		Name:          in.Name,
		ClassID:       in.ClassId,
		AvailableSlot: in.AvailableSlot,
		DayOfWeek:     in.DayOfWeek,
	}
	class, err := h.ClassRepository.CreateClass(ctx,cRequest)
	if err != nil {
		return nil, err
	}
	cResponse := &pb.Class{
		Id:            class.Id.String(),
		Name:          class.Name,
		ClassId:       class.ClassID,
		AvailableSlot: class.AvailableSlot,
		DayOfWeek:     class.DayOfWeek,
	}
	return cResponse, nil
}