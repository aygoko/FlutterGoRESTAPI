package main

import (
	"flag"
	"log"

	"net/http"

	pkgHttp "github.com/aygoko/FlutterGoRESTAPI/go-api/user"
	repository "github.com/aygoko/FlutterGoRESTAPI/repository/ram_storage"
	"github.com/aygoko/FlutterGoRESTAPI/usecases/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	addr := flag.String("addr", ":8080", "HTTP server address")
	flag.Parse()

	userRepo := repository.NewUser()
	userService := service.NewUserService(*userRepo)
	userHandler := pkgHttp.NewUserHandler(userService)

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"*"},
	}))

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	userHandler.WithObjectHandlers(r)

	log.Printf("Starting HTTP server on %s", *addr)
	err := http.ListenAndServe(*addr, r)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
