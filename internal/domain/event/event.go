package event

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID
	Title       string
	Description string
	StartAt     time.Time
	EndAt       time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func New(title, description string, startAt, endAt time.Time) (*Event, error) {
	if title == "" {
		return nil, ErrEmptyTitle
	}
	if !endAt.After(startAt) {
		return nil, ErrInvalidTimeRange
	}
	now := time.Now()
	return &Event{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
		StartAt:     startAt,
		EndAt:       endAt,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}
