package model

import (
	"time"

	"github.com/google/uuid"
)

type MetaContentEntity struct {
	MetaContentID uuid.UUID `json:"content_id"`
	NotebookID    uuid.UUID `json:"notebook_id"`
	UserID        uuid.UUID
	Name          string     `json:"name"`
	Icon          string     `json:"icon"`
	DeletedAt     *time.Time `json:"deleted_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	CreatedAt     time.Time  `json:"created_at"`
}
type CreateMetaContentDTO struct {
	NotebookID uuid.UUID `json:"notebook_id"`
	UserID     uuid.UUID
	Name       string `json:"name"`
	Icon       string `json:"icon"`
}
type ListMetaContentFromUserDTO struct {
	UserID uuid.UUID
}
type ListMetaContentByIdDTO struct {
	MetaContentID uuid.UUID
	UserID        uuid.UUID
}
type UpdateMetaContentDTO struct {
	MetaContentID uuid.UUID
	UserID        uuid.UUID
	Name          string `json:"name"`
	Icon          string `json:"icon"`
}
type DeleteMetaContentDTO struct {
	MetaContentID uuid.UUID
	UserID        uuid.UUID
}
