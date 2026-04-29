package event

import (
	"context"
	"time"

	"github.com/google/uuid"
	domain "github.com/kugue99A/monokuro-cal-server-go/internal/domain/event"
)

type Usecase struct {
	repo domain.Repository
}

func NewUsecase(repo domain.Repository) *Usecase {
	return &Usecase{repo: repo}
}

type CreateInput struct {
	Title       string
	Description string
	StartAt     time.Time
	EndAt       time.Time
}

type UpdateInput struct {
	Title       string
	Description string
	StartAt     time.Time
	EndAt       time.Time
}

func (u *Usecase) GetEvent(ctx context.Context, id uuid.UUID) (*domain.Event, error) {
	return u.repo.FindByID(ctx, id)
}

func (u *Usecase) ListEvents(ctx context.Context) ([]*domain.Event, error) {
	return u.repo.FindAll(ctx)
}

func (u *Usecase) CreateEvent(ctx context.Context, in CreateInput) (*domain.Event, error) {
	ev, err := domain.New(in.Title, in.Description, in.StartAt, in.EndAt)
	if err != nil {
		return nil, err
	}
	if err := u.repo.Save(ctx, ev); err != nil {
		return nil, err
	}
	return ev, nil
}

func (u *Usecase) UpdateEvent(ctx context.Context, id uuid.UUID, in UpdateInput) (*domain.Event, error) {
	ev, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if in.Title == "" {
		return nil, domain.ErrEmptyTitle
	}
	if !in.EndAt.After(in.StartAt) {
		return nil, domain.ErrInvalidTimeRange
	}
	ev.Title = in.Title
	ev.Description = in.Description
	ev.StartAt = in.StartAt
	ev.EndAt = in.EndAt
	ev.UpdatedAt = time.Now()
	if err := u.repo.Update(ctx, ev); err != nil {
		return nil, err
	}
	return ev, nil
}

func (u *Usecase) DeleteEvent(ctx context.Context, id uuid.UUID) error {
	return u.repo.Delete(ctx, id)
}
