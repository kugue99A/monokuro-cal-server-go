package event

import "errors"

var (
	ErrNotFound         = errors.New("event not found")
	ErrEmptyTitle       = errors.New("title is required")
	ErrInvalidTimeRange = errors.New("end time must be after start time")
)
