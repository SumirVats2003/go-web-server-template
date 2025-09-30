package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
)

type App struct {
	Logger *log.Logger
	DB     *pgx.Conn
	Ctx    context.Context
}

func InitApp(ctx context.Context) (*App, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	logger.Println("Successfully connected to Database")

	app := &App{
		Logger: logger,
		DB:     conn,
		Ctx:    ctx,
	}
	return app, nil
}

func (a *App) Heartbeat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status is available")
}
