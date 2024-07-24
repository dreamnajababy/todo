package todo

import (
	"encoding/base64"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	IN_PROGRESS Status = "IN_PROGRESS"
	COMPLETE    Status = "COMPLETE"
)

type Todo struct {
	Id          uuid.UUID
	Title       string
	Description string
	CreatedAt   time.Time
	Image       string
	Status      Status
}

func NewTodo(id string, title, description string, createdAt string, image string, status Status) (Todo, error) {
	if id == "" {
		return Todo{}, errors.New("id cannot be empty")
	}
	uuid, uuidErr := uuid.Parse(id)
	if uuidErr != nil {
		return Todo{}, errors.New("id must be UUID format")
	}
	if title == "" {
		return Todo{}, errors.New("title cannot be empty")
	}
	if createdAt == "" {
		return Todo{}, errors.New("created_at cannot be empty")
	}
	if status == "" {
		return Todo{}, errors.New("status cannot be empty")
	}
	if status != IN_PROGRESS && status != COMPLETE {
		return Todo{}, errors.New("status must be IN_PROGRESS or COMPLETE")
	}
	parsedCreatedAt, isRFC3339 := time.Parse(time.RFC3339, createdAt)
	if isRFC3339 != nil {
		return Todo{}, errors.New("created_at must be RFC3339 format")
	}
	if len(title) > 100 {
		return Todo{}, errors.New("title must not over 100 characters")
	}
	_, base64Err := base64.StdEncoding.DecodeString(image)
	if base64Err != nil {
		return Todo{}, errors.New("image must be Base64 Encode format")
	}
	return Todo{
		Id:          uuid,
		Title:       title,
		Description: description,
		CreatedAt:   parsedCreatedAt,
		Image:       image,
		Status:      status,
	}, nil
}
