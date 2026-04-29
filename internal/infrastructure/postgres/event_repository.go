package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	domain "github.com/kugue99A/monokuro-cal-server-go/internal/domain/event"
	"github.com/kugue99A/monokuro-cal-server-go/internal/infrastructure/postgres/db"
)

type EventRepository struct {
	queries *db.Queries
}

func NewEventRepository(pool *pgxpool.Pool) *EventRepository {
	return &EventRepository{queries: db.New(pool)}
}

func (r *EventRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Event, error) {
	row, err := r.queries.GetEvent(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return toDomain(row), nil
}

func (r *EventRepository) FindAll(ctx context.Context) ([]*domain.Event, error) {
	rows, err := r.queries.ListEvents(ctx)
	if err != nil {
		return nil, err
	}
	events := make([]*domain.Event, len(rows))
	for i, row := range rows {
		events[i] = toDomain(row)
	}
	return events, nil
}

func (r *EventRepository) Save(ctx context.Context, ev *domain.Event) error {
	_, err := r.queries.CreateEvent(ctx, db.CreateEventParams{
		ID:          ev.ID,
		Title:       ev.Title,
		Description: ev.Description,
		StartAt:     pgtype.Timestamptz{Time: ev.StartAt, Valid: true},
		EndAt:       pgtype.Timestamptz{Time: ev.EndAt, Valid: true},
	})
	return err
}

func (r *EventRepository) Update(ctx context.Context, ev *domain.Event) error {
	_, err := r.queries.UpdateEvent(ctx, db.UpdateEventParams{
		ID:          ev.ID,
		Title:       ev.Title,
		Description: ev.Description,
		StartAt:     pgtype.Timestamptz{Time: ev.StartAt, Valid: true},
		EndAt:       pgtype.Timestamptz{Time: ev.EndAt, Valid: true},
	})
	return err
}

func (r *EventRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.queries.DeleteEvent(ctx, id)
}

func toDomain(row db.Event) *domain.Event {
	return &domain.Event{
		ID:          row.ID,
		Title:       row.Title,
		Description: row.Description,
		StartAt:     row.StartAt.Time,
		EndAt:       row.EndAt.Time,
		CreatedAt:   row.CreatedAt.Time,
		UpdatedAt:   row.UpdatedAt.Time,
	}
}
