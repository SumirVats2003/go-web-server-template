package repository

import (
	"context"
	"fmt"

	"github.com/SumirVats2003/go-web-server-template/models"
	"github.com/jackc/pgx/v5"
)

type AuthRepository struct {
	db    *pgx.Conn
	table string
	ctx   context.Context
}

func InitAuthRepository(db *pgx.Conn, ctx context.Context) AuthRepository {
	a := AuthRepository{db: db, ctx: ctx}
	a.table = "users"
	return a
}

func (a AuthRepository) LoginUser(email string) {
	res := a.db.QueryRow(a.ctx, "SELECT * FROM users WHERE email=$1", email)
	fmt.Println(res)
}

func (a AuthRepository) SignupUser(signupRequest models.SignupRequest) (models.User, error) {
	a.db.Query(
		a.ctx,
		"INSERT INTO users VALUES name=$1, email=$2, password=$3, created_at=$4",
		signupRequest.Name,
		signupRequest.Email,
		signupRequest.Password,
		signupRequest.CreatedAt,
	)
	return models.User{}, nil
}

func (a AuthRepository) findUser(email string) {
}
