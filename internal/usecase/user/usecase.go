package user

import (
	"context"
	"time"

	"github.com/google/uuid"
	domain "github.com/kugue99A/monokuro-cal-server-go/internal/domain/user"
)

type Usecase struct {
	repo domain.Repository
}

func NewUsecase(repo domain.Repository) *Usecase {
	return &Usecase{repo: repo}
}

type CreateInput struct {
	Email string
	Name  string
}

type UpdateInput struct {
	Name string
}

func (u *Usecase) GetUser(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	return u.repo.FindByID(ctx, id)
}

func (u *Usecase) CreateUser(ctx context.Context, in CreateInput) (*domain.User, error) {
	user, err := domain.New(in.Email, in.Name)
	if err != nil {
		return nil, err
	}
	if err := u.repo.Save(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *Usecase) UpdateUser(ctx context.Context, id uuid.UUID, in UpdateInput) (*domain.User, error) {
	user, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if in.Name == "" {
		return nil, domain.ErrEmptyName
	}
	user.Name = in.Name
	user.UpdatedAt = time.Now()
	if err := u.repo.Update(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *Usecase) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return u.repo.Delete(ctx, id)
}
