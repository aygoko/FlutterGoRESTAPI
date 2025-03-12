package http

import (
	"encoding/json"
	"net/http"

	repository "github.com/aygoko/FlutterGoRESTAPI/domain"
	"github.com/aygoko/FlutterGoRESTAPI/usecases/service"
	"github.com/go-chi/chi/v5"
)

// UserHandler handles user-related HTTP endpoints
type UserHandler struct {
	Service repository.UserService
}

// NewUserHandler creates a new user handler instance
func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{
		Service: s, // Direct assignment if service implements the interface
	}
}

// WithObjectHandlers registers user routes
func (h *UserHandler) WithObjectHandlers(r *chi.Mux) {
	r.Route("/api/users", func(r chi.Router) {
		r.Get("/", h.GetAllUsers)
		r.Post("/", h.CreateUser)
		r.Get("/{login}", h.GetUserByLogin)
		r.Delete("/{login}", h.DeleteUser)
	})
}

// @Summary Get all users
// @Description Retrieve a list of all registered users
// @Tags Users
// @Success 200 {array} repository.User "List of users"
// @Router /api/users [get]
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Service.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

// @Summary Create a new user
// @Description Create a user with login, email, and password
// @Tags Users
// @Param user body repository.User true "User details"
// @Success 201 {object} repository.User "Created user"
// @Failure 400 {string} error "Invalid request or duplicate user"
// @Router /api/users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user repository.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdUser, err := h.Service.Save(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

// @Summary Get user by login
// @Description Retrieve a user by their login
// @Tags Users
// @Param login path string true "User login"
// @Success 200 {object} repository.User "User details"
// @Failure 404 {string} error "User not found"
// @Router /api/users/{login} [get]
func (h *UserHandler) GetUserByLogin(w http.ResponseWriter, r *http.Request) {
	login := chi.URLParam(r, "login")
	user, err := h.Service.Get(login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// @Summary Delete user by login
// @Description Delete a user by their login
// @Tags Users
// @Param login path string true "User login to delete"
// @Success 204 "User deleted successfully"
// @Failure 404 {string} error "User not found"
// @Router /api/users/{login} [delete]
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	login := chi.URLParam(r, "login")
	err := h.Service.Delete(login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
