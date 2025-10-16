package handlers

import (
	"net/http"
	"simple-setup/internal/middlewares"
	"simple-setup/internal/models"
	"simple-setup/internal/services"

	"github.com/go-chi/chi/v5"
)

type BoardHandler struct {
	service     *services.BoardService
	authService *middlewares.AuthService
}

func NewBoardHandler(s *services.BoardService, auth *middlewares.AuthService) *BoardHandler {
	return &BoardHandler{
		service:     s,
		authService: auth,
	}
}

func (h *BoardHandler) RegisterRoutes(r chi.Router) {
	r.Route("/boards", func(boards chi.Router) {
		boards.Use(h.authService.Middleware)

		boards.Post("/", h.CreateBoard)
		boards.Get("/", h.GetAllBoards)
		boards.Get("/my", h.GetMyBoards)
		boards.Get("/{id}", h.GetBoard)
		boards.Put("/{id}", h.UpdateBoard)
		boards.Delete("/{id}", h.DeleteBoard)
	})
}

func (h *BoardHandler) CreateBoard(w http.ResponseWriter, r *http.Request) {
	req, err := tryDecodeJSON[CreateBoardRequest](r.Body)
	if err != nil {
		respondError(w, err, http.StatusBadRequest)
		return
	}

	// Get authenticated user (optional - can be used for ownership tracking)
	_, err = middlewares.GetUserFromContext(r.Context())
	if err != nil {
		respondError(w, err, http.StatusUnauthorized)
		return
	}

	board := models.Board{
		Name: req.Name,
	}

	if err = h.service.CreateBoard(&board); err != nil {
		respondError(w, err, http.StatusBadRequest)
		return
	}

	respondJSON(w, board, http.StatusCreated)
}

func (h *BoardHandler) GetBoard(w http.ResponseWriter, r *http.Request) {
	id, err := tryGetUUIDParam("id", r)
	if err != nil {
		respondError(w, err, http.StatusBadRequest)
		return
	}

	board, err := h.service.GetBoard(id)
	if err != nil {
		respondError(w, err, http.StatusNotFound)
		return
	}

	respondJSON(w, board, http.StatusOK)
}

func (h *BoardHandler) GetAllBoards(w http.ResponseWriter, r *http.Request) {
	boards, err := h.service.GetAllBoards()
	if err != nil {
		respondError(w, err, http.StatusInternalServerError)
		return
	}

	respondJSON(w, boards, http.StatusOK)
}

func (h *BoardHandler) GetMyBoards(w http.ResponseWriter, r *http.Request) {
	_, err := middlewares.GetUserFromContext(r.Context())
	if err != nil {
		respondError(w, err, http.StatusUnauthorized)
		return
	}

	// TODO: Get userID from database using token.UID and call GetBoardsByUser
	// For now, just return all boards
	boards, err := h.service.GetAllBoards()
	if err != nil {
		respondError(w, err, http.StatusInternalServerError)
		return
	}

	respondJSON(w, boards, http.StatusOK)
}

func (h *BoardHandler) UpdateBoard(w http.ResponseWriter, r *http.Request) {
	req, err := tryDecodeJSON[UpdateBoardRequest](r.Body)
	if err != nil {
		respondError(w, err, http.StatusBadRequest)
		return
	}

	id, err := tryGetUUIDParam("id", r)
	if err != nil {
		respondError(w, err, http.StatusBadRequest)
		return
	}

	board := models.Board{
		ID:   id,
		Name: req.Name,
	}

	if err = h.service.UpdateBoard(&board); err != nil {
		respondError(w, err, http.StatusBadRequest)
		return
	}

	respondJSON(w, board, http.StatusOK)
}

func (h *BoardHandler) DeleteBoard(w http.ResponseWriter, r *http.Request) {
	id, err := tryGetUUIDParam("id", r)
	if err != nil {
		respondError(w, err, http.StatusBadRequest)
		return
	}

	if err = h.service.DeleteBoard(id); err != nil {
		respondError(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
