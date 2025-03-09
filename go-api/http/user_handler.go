package http

import (
	"encoding/json"
	"net/http"

	"github.com/aygoko/FlutterGoRESTAPI/service"
	"github.com/go-chi/chi/v5"
)

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := h.service.GetAllUsers()
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user service.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdUser, err := h.service.CreateUser(user.Login, user.Email, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

func (h *UserHandler) GetUserByLogin(w http.ResponseWriter, r *http.Request) {
	login := chi.URLParam(r, "login")
	user, err := h.service.GetUserByLogin(login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	login := chi.URLParam(r, "login")

	err := h.service.DeleteUser(login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
