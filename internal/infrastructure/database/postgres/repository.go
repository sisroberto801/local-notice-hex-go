package postgres

import (
	"context"
	"local-notice-hex-go/internal/domain/user"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) user.Repository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, u *user.User) (*user.User, error) {
	query := `
        INSERT INTO users (username, password, status, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, username, password, status, created_at, updated_at
    `

	var createdUser user.User
	err := r.db.QueryRow(ctx, query,
		u.Username, u.Password, u.Status, time.Now(), time.Now()).
		Scan(&createdUser.ID, &createdUser.Username, &createdUser.Password,
			&createdUser.Status, &createdUser.CreatedAt, &createdUser.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &createdUser, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id int64) (*user.User, error) {
	query := `
        SELECT id, username, password, status, created_at, updated_at
        FROM users
        WHERE id = $1
    `

	var u user.User
	err := r.db.QueryRow(ctx, query, id).
		Scan(&u.ID, &u.Username, &u.Password, &u.Status, &u.CreatedAt, &u.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*user.User, error) {
	query := `
        SELECT id, username, password, status, created_at, updated_at
        FROM users
        WHERE username = $1
    `

	var u user.User
	err := r.db.QueryRow(ctx, query, username).
		Scan(&u.ID, &u.Username, &u.Password, &u.Status, &u.CreatedAt, &u.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *UserRepository) FindAll(ctx context.Context) ([]*user.User, error) {
	query := `
        SELECT id, username, password, status, created_at, updated_at
        FROM users
        ORDER BY created_at DESC
    `

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*user.User
	for rows.Next() {
		var u user.User
		err := rows.Scan(&u.ID, &u.Username, &u.Password, &u.Status, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) Update(ctx context.Context, id int64, u *user.User) (*user.User, error) {
	query := `
        UPDATE users
        SET username = $1, password = $2, status = $3, updated_at = $4
        WHERE id = $5
        RETURNING id, username, password, status, created_at, updated_at
    `

	var updatedUser user.User
	err := r.db.QueryRow(ctx, query,
		u.Username, u.Password, u.Status, time.Now(), id).
		Scan(&updatedUser.ID, &updatedUser.Username, &updatedUser.Password,
			&updatedUser.Status, &updatedUser.CreatedAt, &updatedUser.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := r.db.Exec(ctx, query, id)
	return err
}
