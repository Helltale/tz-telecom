package usecase

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/Helltale/tz-telecom/internal/domain"
)

type UserRepository interface {
	Save(ctx context.Context, u *domain.User) error
}

type UserUseCaseInterface interface {
	RegisterUser(ctx context.Context, user *domain.User) error
}

type UserUseCase struct {
	repo UserRepository
}

func NewUserUseCase(r UserRepository) *UserUseCase {
	return &UserUseCase{repo: r}
}

func (uc *UserUseCase) RegisterUser(ctx context.Context, u *domain.User) error {
	if u.Age < 18 {
		return errors.New("user must be 18+")
	}
	if len(u.Password) < 8 {
		return errors.New("password too short")
	}
	hashed, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hashed)
	return uc.repo.Save(ctx, u)
}
