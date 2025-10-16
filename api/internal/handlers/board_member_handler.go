package handlers

import (
	"net/http"
	"simple-setup/internal/middlewares"
	"simple-setup/internal/models"
	"simple-setup/internal/services"

	"github.com/go-chi/chi/v5"
)

type BoardMemberHandler struct {
	service     *services.BoardMemberService
	authService *middlewares.AuthService
}

func NewBoardMemberHandler(s *services.BoardMemberService, auth *middlewares.AuthService) *BoardMemberHandler {
	return &BoardMemberHandler{
		service:     s,
		authService: auth,
	}
}

func (h *BoardMemberHandler) RegisterRoutes(r chi.Router) {
	r.Route("/board-members", func(members chi.Router) {
		members.Use(h.authService.Middleware)

		members.Post("/", h.AddMember)
		members.Get("/board/{boardId}", h.GetMembersByBoard)
		members.Get("/{id}", h.GetMember)
		members.Patch("/{id}/role", h.UpdateMemberRole)
		members.Delete("/{id}", h.RemoveMember)
	})
}

func (h *BoardMemberHandler) AddMember(w http.ResponseWriter, r *http.Request) {
	req, err := tryDecodeJSON[AddMemberRequest](r.Body)
	if err != nil {
		respondError(w, err, http.StatusBadRequest)
		return
	}

	member := models.BoardMember{
		UserID:  req.UserID,
		BoardID: req.BoardID,
		Role:    req.Role,
	}

	if err = h.service.AddMember(&member); err != nil {
		respondError(w, err, http.StatusBadRequest)
		return
	}

	respondJSON(w, member, http.StatusCreated)
}

func (h *BoardMemberHandler) GetMembersByBoard(w http.ResponseWriter, r *http.Request) {
	boardID, err := tryGetUUIDParam("boardId", r)
	if err != nil {
		respondError(w, err, http.StatusBadRequest)
		return
	}

	members, err := h.service.GetMembersByBoard(boardID)
	if err != nil {
		respondError(w, err, http.StatusInternalServerError)
		return
	}

	respondJSON(w, members, http.StatusOK)
}

func (h *BoardMemberHandler) GetMember(w http.ResponseWriter, r *http.Request) {
	id, err := tryGetUUIDParam("id", r)
	if err != nil {
		respondError(w, err, http.StatusBadRequest)
		return
	}

	member, err := h.service.GetMember(id)
	if err != nil {
		respondError(w, err, http.StatusNotFound)
		return
	}

	respondJSON(w, member, http.StatusOK)
}

func (h *BoardMemberHandler) UpdateMemberRole(w http.ResponseWriter, r *http.Request) {
	req, err := tryDecodeJSON[UpdateMemberRoleRequest](r.Body)
	if err != nil {
		respondError(w, err, http.StatusBadRequest)
		return
	}

	id, err := tryGetUUIDParam("id", r)
	if err != nil {
		respondError(w, err, http.StatusBadRequest)
		return
	}

	if err = h.service.UpdateMemberRole(id, req.Role); err != nil {
		respondError(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *BoardMemberHandler) RemoveMember(w http.ResponseWriter, r *http.Request) {
	id, err := tryGetUUIDParam("id", r)
	if err != nil {
		respondError(w, err, http.StatusBadRequest)
		return
	}

	if err = h.service.RemoveMember(id); err != nil {
		respondError(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
