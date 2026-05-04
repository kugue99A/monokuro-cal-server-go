package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	domain "github.com/kugue99A/monokuro-cal-server-go/internal/domain/user"
	"github.com/kugue99A/monokuro-cal-server-go/internal/infrastructure/postgres/db"
)

type UserRepository struct {
	queries *db.Queries
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{queries: db.New(pool)}
}

func (r *UserRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	row, err := r.queries.GetUser(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return userToDomain(row), nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	row, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return userToDomain(row), nil
}

func (r *UserRepository) Save(ctx context.Context, u *domain.User) error {
	_, err := r.queries.CreateUser(ctx, db.CreateUserParams{
		ID:    u.ID,
		Email: u.Email,
		Name:  u.Name,
	})
	return err
}

func (r *UserRepository) Update(ctx context.Context, u *domain.User) error {
	_, err := r.queries.UpdateUser(ctx, db.UpdateUserParams{
		ID:   u.ID,
		Name: u.Name,
	})
	return err
}

func (r *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.queries.DeleteUser(ctx, id)
}

func userToDomain(row db.User) *domain.User {
	return &domain.User{
		ID:        row.ID,
		Email:     row.Email,
		Name:      row.Name,
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
	}
}
