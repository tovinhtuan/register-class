package requests

import "github.com/google/uuid"

type DeleteClassRequest struct {
	Id      uuid.UUID
	ClassID string
}