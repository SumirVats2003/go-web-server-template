package routes

import (
	"github.com/SumirVats2003/go-web-server-template/internal/app"
	"github.com/SumirVats2003/go-web-server-template/internal/middleware"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.App) (*chi.Mux, error) {
	r := chi.NewRouter()
	r.Use(middleware.CorsMiddleware)

	// heartbeat
	r.Get("/heartbeat", app.Heartbeat)

	// route groups
	authRouter, err := InitAuthRoutes(app.DB, app.Ctx)
	if err != nil {
		return nil, err
	}

	r.Mount("/auth", authRouter)

	return r, nil
}
