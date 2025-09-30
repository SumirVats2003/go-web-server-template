package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SumirVats2003/go-web-server-template/internal/app"
	"github.com/SumirVats2003/go-web-server-template/internal/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using empty variables")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	app, err := app.InitApp(ctx)
	if err != nil {
		panic(err)
	}

	defer app.DB.Close(app.Ctx)

	port := ":8080"
	r, err := routes.SetupRoutes(app)
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Addr:         fmt.Sprintf("%s", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("server running on port:8080\n")

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
