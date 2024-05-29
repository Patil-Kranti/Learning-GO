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
	router.Handle("/*", public())
	router.Get("/", handlers.Make(handlers.HandleHome))
	fmt.Println("Hello, World!1")

	listenAdder := os.Getenv("LISTEN_ADDR")
	slog.Info("Http server started", "listenAddr", listenAdder)
	http.ListenAndServe(listenAdder, router)
}
