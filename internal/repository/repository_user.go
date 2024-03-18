package repository

import (
	"context"
	"math/rand"
	"server/internal/presentation"
)

func (r *repository) CreateUser(ctx context.Context, user *presentation.User) (*presentation.User, error) {
	var lastInsertId int

	query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) returning id"

	err := r.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password).Scan(&lastInsertId)

	if err != nil {
		return &presentation.User{}, err
	}

	user.ID = int64(lastInsertId)

	query = "INSERT INTO accounts(id, user_id, account_number,balance) VALUES ($1, $2, $3)"

	_, err = r.db.ExecContext(ctx, query, rand.Intn(100), user.ID, "1234243", 5000)
	if err != nil {
		return &presentation.User{}, err
	}

	return user, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*presentation.User, error) {

	u := presentation.User{}

	query := "SELECT id, username, email, password FROM users WHERE email = $1"
	err := r.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.Username, &u.Email, &u.Password)
	if err != nil {
		return &presentation.User{}, err
	}
	return &u, nil

}
