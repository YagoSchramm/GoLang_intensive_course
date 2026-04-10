package repository

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/YagoSchramm/base-auth-v2-session/model"
	"github.com/google/uuid"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(d *sql.DB) *UserRepository {
	return &UserRepository{db: d}
}

//go:embed _query/user/create_user.sql
var createUserQuery string

//go:embed _query/user/find_by_email_user.sql
var findByEmailUserQuery string

//go:embed _query/user/find_by_id_user.sql
var findByIDUserQuery string

//go:embed _query/user/list_user.sql
var listUserQuery string

//go:embed _query/user/update_user.sql
var updateUserQuery string

//go:embed _query/user/delete_user.sql
var deleteUserQuery string

func (r *UserRepository) CreateUser(ctx context.Context, u *model.User) (*model.User, error) {
	if u.ID == "" {
		u.ID = uuid.NewString()
	}

	_, err := r.db.ExecContext(ctx, createUserQuery, u.ID, u.Email, u.Password)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) GetUser(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	row := r.db.QueryRowContext(ctx, findByEmailUserQuery, email)
	if err := row.Scan(&user.ID, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	row := r.db.QueryRowContext(ctx, findByIDUserQuery, id)
	if err := row.Scan(&user.ID, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) ListUsers(ctx context.Context) ([]*model.User, error) {
	rows, err := r.db.QueryContext(ctx, listUserQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Email, &u.Password); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, u *model.User) error {
	_, err := r.db.ExecContext(ctx, updateUserQuery, u.ID, u.Email, u.Password)
	return err
}

func (r *UserRepository) DeleteUser(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, deleteUserQuery, id)
	return err
}
