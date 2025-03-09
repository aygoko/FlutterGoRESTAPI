package main

import (
	"flag"
	"log"

	"github.com/aygoko/FlutterGoRESTAPI/http"
	"github.com/aygoko/FlutterGoRESTAPI/repository"
	"github.com/aygoko/FlutterGoRESTAPI/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	addr := flag.String("addr", ":8080", "HTTP server address")
	flag.Parse()

	userRepo := repository.NewUser()
	userService := service.NewUserService(userRepo)
	userHandler := http.NewUserHandler(userService)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	userHandler.WithObjectHandlers(r)

	log.Printf("Starting HTTP server on %s", *addr)
	err := http.ListenAndServe(*addr, r)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
