package api

import (
	"context"
	"errors"
	"time"

	"github.com/SumirVats2003/go-web-server-template/internal/repository"
	"github.com/SumirVats2003/go-web-server-template/models"
	"github.com/SumirVats2003/go-web-server-template/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthApi struct {
	db             *pgx.Conn
	authRepository repository.AuthRepository
}

var jwtSecret string

func InitAuthApi(db *pgx.Conn, ctx context.Context) (AuthApi, error) {
	jwtSecret = utils.GetEnv("JWT_SECRET", "")
	if jwtSecret == "" {
		return AuthApi{}, errors.New("Error: Could not load the JWT Secret key. The env is not set correctly")
	}

	u := repository.InitAuthRepository(db, ctx)
	userApi := AuthApi{db: db, authRepository: u}
	return userApi, nil
}

func (a AuthApi) Login(loginRequest models.LoginRequest) (string, error) {
	user, err := a.authRepository.LoginUser(loginRequest.Email)
	if err != nil {
		return "", errors.New("Invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.User.Password), []byte(loginRequest.Password))
	if err != nil {
		return "", err
	}

	userToken, err := createToken(loginRequest.Email)
	if err != nil {
		return "", err
	}

	return userToken, nil
}

func (a AuthApi) Signup(signupRequest models.SignupRequest) (models.User, error) {
	hashedPassword, err := hashPassword(signupRequest.Password)
	if err != nil {
		return models.User{}, err
	}

	signupRequest.Password = hashedPassword
	signupRequest.CreatedAt = utils.GetCurrentTimestamp()
	success, err := a.authRepository.SignupUser(signupRequest)
	return success, err
}

func createToken(email string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": email,
		"iss": "formify",
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := claims.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
