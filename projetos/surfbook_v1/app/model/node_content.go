package model

import (
	"time"

	"github.com/google/uuid"
)

type NodeContentEntity struct {
	NodeID     uuid.UUID  `json:"node_id"`
	ContentID  uuid.UUID  `json:"content_id"`
	UserID     uuid.UUID  `json:"user_id"`
	NotebookID uuid.UUID  `json:"notebook_id"`
	DeletedAt  *time.Time `json:"deleted_at"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

type CreateNodeContentDTO struct {
	ContentID  uuid.UUID `json:"content_id"`
	NotebookID uuid.UUID `json:"notebook_id"`
	UserID     uuid.UUID
}

type ListNodeContentFromUserDTO struct {
	UserID uuid.UUID
}

type FindNodeContentFromUserDTO struct {
	UserID uuid.UUID
	NodeID uuid.UUID
}

type UpdateNodeContentDTO struct {
	UserID     uuid.UUID
	NodeID     uuid.UUID
	ContentID  *uuid.UUID `json:"content_id"`
	NotebookID *uuid.UUID `json:"notebook_id"`
}

type DeleteNodeContentDTO struct {
	UserID uuid.UUID
	NodeID uuid.UUID
}
