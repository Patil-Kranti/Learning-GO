package main

import (
	"fmt"
	"gotth_tutorial/handlers"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	router := chi.NewMux()
	router.Get("/foo", handlers.Handlefoo)
	fmt.Println("Hello, World!")

	listenAdder := os.Getenv("LISTEN_ADDR")
	slog.Info("Http server started", "listenAddr", listenAdder)
	http.ListenAndServe(listenAdder, router)
}
