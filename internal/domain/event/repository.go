package event

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*Event, error)
	FindAllByUser(ctx context.Context, userID uuid.UUID) ([]*Event, error)
	Save(ctx context.Context, event *Event) error
	Update(ctx context.Context, event *Event) error
	Delete(ctx context.Context, id uuid.UUID) error
}
