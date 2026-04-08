package model

import (
	"time"

	"github.com/google/uuid"
)

type TagEntity struct {
	TagID     uuid.UUID  `json:"tag_id"`
	Name      string     `json:"name"`
	Color     string     `json:"color"`
	UserID    uuid.UUID  `json:"user_id"`
	DeletedAt *time.Time `json:"deleted_at"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type CreateTagDTO struct {
	UserID uuid.UUID

	Name  string `json:"name"`
	Color string `json:"color"`
}

type ListTagsFromUserDTO struct {
	UserID uuid.UUID
}

type FindTagFromUserDTO struct {
	UserID uuid.UUID
	TagID  uuid.UUID
}

type UpdateTagDTO struct {
	UserID uuid.UUID
	TagID  uuid.UUID
	Name   string `json:"name"`
	Color  string `json:"color"`
}

type DeleteTagDTO struct {
	UserID uuid.UUID
	TagID  uuid.UUID
}
