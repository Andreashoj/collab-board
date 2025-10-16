package handlers

import (
	"net/http"
	"simple-setup/internal/middlewares"
	"simple-setup/internal/models"
	"simple-setup/internal/services"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	service     *services.UserService
	authService *middlewares.AuthService
}

func NewUserHandler(s *services.UserService, auth *middlewares.AuthService) *UserHandler {
	return &UserHandler{
		service:     s,
		authService: auth,
	}
}

func (h *UserHandler) RegisterRoutes(r *chi.Mux) {
	r.Route("/user", func(user chi.Router) {
		user.Use(h.authService.Middleware)

		user.Post("/", h.CreateUser)
		user.Put("/{id}", h.UpdateUser)
		user.Get("/", h.GetUser)
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
		Email:        req.Email,
		PasswordHash: req.Password, // Should be hashed
	}

	if err = h.service.CreateUser(&user); err != nil {
		respondError(w, err, http.StatusBadRequest)
		return
	}

	respondJSON(w, user, http.StatusCreated)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	user, err := middlewares.GetUserFromContext(r.Context())

	if err != nil {
		respondError(w, err, http.StatusUnauthorized)
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

	id, err := tryGetUUIDParam("id", r)
	if err != nil {
		respondError(w, err, http.StatusBadRequest)
		return
	}

	user := models.User{
		ID:    id,
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
	id, err := tryGetUUIDParam("id", r)
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
