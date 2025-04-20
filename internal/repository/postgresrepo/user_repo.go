package postgresrepo

import (
	"context"
	"database/sql"

	"github.com/Helltale/tz-telecom/internal/domain"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Save(ctx context.Context, u *domain.User) error {
	query := `
		INSERT INTO users (first_name, last_name, age, is_married, password)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	return r.db.QueryRowContext(ctx, query,
		u.FirstName,
		u.LastName,
		u.Age,
		u.IsMarried,
		u.Password,
	).Scan(&u.ID)
}
