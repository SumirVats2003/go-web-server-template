package repository

import (
	"context"
	"errors"

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

func (a AuthRepository) LoginUser(email string) (models.User, error) {
	query := "SELECT userid, name, email, password, created_at FROM users WHERE email=$1"

	var id string
	var userInfo models.SignupRequest

	err := a.db.QueryRow(a.ctx, query, email).Scan(
		&id,
		&userInfo.Name,
		&userInfo.Email,
		&userInfo.Password,
		&userInfo.CreatedAt,
	)

	if err != nil {
		return models.User{}, err
	}

	return models.User{
		UserId: id,
		User:   userInfo,
	}, nil
}

func (a AuthRepository) SignupUser(signupRequest models.SignupRequest) (string, error) {
	if a.isExistingUser(signupRequest.Email) {
		return "", errors.New("User Already Exists")
	}

	var userId string
	query := getSignupQuery()

	err := a.db.QueryRow(a.ctx, query,
		signupRequest.Name,
		signupRequest.Email,
		signupRequest.Password,
		signupRequest.CreatedAt,
	).Scan(&userId)

	if err != nil {
		return "", err
	}

	return userId, nil
}

func getSignupQuery() string {
	return `
		INSERT INTO users (name, email, password, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING userId
	`
}

func (a AuthRepository) isExistingUser(email string) bool {
	query := "SELECT userid FROM users WHERE email=$1"
	var id string

	err := a.db.QueryRow(a.ctx, query, email).Scan(&id)
	if err != nil {
		return false
	}

	return true
}
