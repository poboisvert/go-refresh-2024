package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	h "url/api"
	redisDb "url/database"

	"url/pkg" // Assuming the shortener package is imported here

	"github.com/joho/godotenv"
)

func goDotEnv(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env")
	}
	return os.Getenv(key)
}

func main() {
	repo := getDB()
	service := pkg.NewRedirectService(repo)
	handler := h.NewHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/{code}", handler.Get)
	r.Post("/", handler.Post)

	errs := make(chan error, 2)

	go func() {
		fmt.Println("Listening on port :8000")
		port := os.Getenv("PORT")
		if port == "" {
			port = "8000"
		}
		errs <- http.ListenAndServe(":"+port, r)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated %s", <-errs)
}

func getDB() pkg.RedirectManager {
	// switch case for DB redis or mongo
	switch goDotEnv("URL_DB") {
	case "redis":
		redisURL := goDotEnv("REDIS_URL")
		repo, err := redisDb.NewRedisRepository(redisURL)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	}

	// if nothing
	return nil
}
