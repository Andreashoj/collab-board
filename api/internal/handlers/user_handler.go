package handlers

import (
	"net/http"
	"simple-setup/internal/models"
	"simple-setup/internal/services"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(s *services.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) RegisterRoutes(r *chi.Mux) {
	r.Route("/user", func(user chi.Router) {
		user.Post("/", h.CreateUser)
		user.Put("/{id}", h.UpdateUser)
		user.Get("/{id}", h.GetUser)
		user.Delete("/{id}", h.DeleteUser)
	})
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	req, err := tryDecodeJSON[CreateUserRequest](r.Body)

	if err != nil {
		respondError(w, err, http.StatusBadRequest)
		return
	}

	user := models.User{
		Name:  req.Name,
		Email: req.Email,
	}

	if err = h.service.CreateUser(&user); err != nil {
		respondError(w, err, http.StatusBadRequest)
		return
	}

	respondJSON(w, user, http.StatusCreated)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := tryGetUintParam("id", r)

	if err != nil {
		respondError(w, err, http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUser(id)

	if err != nil {
		respondError(w, err, http.StatusNotFound)
		return
	}

	respondJSON(w, user, http.StatusOK)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	req, err := tryDecodeJSON[UpdateUserRequest](r.Body)
	if err != nil {
		respondError(w, err, http.StatusBadRequest)
		return
	}

	id, err := tryGetUintParam("id", r)
	if err != nil {
		respondError(w, err, http.StatusNotFound)
		return
	}

	user := models.User{
		ID:    id,
		Name:  req.Name,
		Email: req.Email,
	}

	err = h.service.UpdateUser(&user)
	if err != nil {
		respondError(w, err, http.StatusBadRequest)
		return
	}

	respondJSON(w, user, http.StatusOK)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := tryGetUintParam("id", r)

	if err != nil {
		respondError(w, err, http.StatusBadRequest)
		return
	}

	err = h.service.DeleteUser(id)
	if err != nil {
		respondError(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
