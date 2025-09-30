package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/SumirVats2003/go-web-server-template/internal/api"
	"github.com/SumirVats2003/go-web-server-template/models"
	"github.com/SumirVats2003/go-web-server-template/utils"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

type AuthRouter struct {
	db      *pgx.Conn
	authApi api.AuthApi
}

func InitAuthRoutes(db *pgx.Conn, ctx context.Context) (chi.Router, error) {
	authApi, err := api.InitAuthApi(db, ctx)

	if err != nil {
		return nil, err
	}

	a := AuthRouter{db: db, authApi: authApi}
	r := chi.NewRouter()
	r.Get("/login", a.Login)
	r.Post("/signup", a.Signup)
	return r, nil
}

func (a AuthRouter) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest models.LoginRequest
	loginRequest, err := utils.ParseJSON(loginRequest, r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := a.authApi.Login(loginRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

func (a AuthRouter) Signup(w http.ResponseWriter, r *http.Request) {
	var signupRequest models.SignupRequest
	signupRequest, err := utils.ParseJSON(signupRequest, r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	success, err := a.authApi.Signup(signupRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"userId": success.UserId,
		"user":   success.User,
	})
}
