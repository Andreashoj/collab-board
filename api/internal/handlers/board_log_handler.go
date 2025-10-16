package handlers

import (
	"net/http"
	"simple-setup/internal/middlewares"
	"simple-setup/internal/models"
	"simple-setup/internal/services"

	"github.com/go-chi/chi/v5"
)

type BoardLogHandler struct {
	service     *services.BoardLogService
	authService *middlewares.AuthService
}

func NewBoardLogHandler(s *services.BoardLogService, auth *middlewares.AuthService) *BoardLogHandler {
	return &BoardLogHandler{
		service:     s,
		authService: auth,
	}
}

func (h *BoardLogHandler) RegisterRoutes(r chi.Router) {
	r.Route("/board-logs", func(logs chi.Router) {
		logs.Use(h.authService.Middleware)

		logs.Post("/", h.CreateLog)
		logs.Get("/board/{boardId}", h.GetLogsByBoard)
		logs.Get("/{id}", h.GetLog)
		logs.Delete("/{id}", h.DeleteLog)
	})
}

func (h *BoardLogHandler) CreateLog(w http.ResponseWriter, r *http.Request) {
	req, err := tryDecodeJSON[CreateLogRequest](r.Body)
	if err != nil {
		respondError(w, err, http.StatusBadRequest)
		return
	}

	log := models.BoardLog{
		BoardID: req.BoardID,
		UserID:  req.UserID,
		Change:  req.Change,
	}

	if err = h.service.CreateLog(&log); err != nil {
		respondError(w, err, http.StatusBadRequest)
		return
	}

	respondJSON(w, log, http.StatusCreated)
}

func (h *BoardLogHandler) GetLogsByBoard(w http.ResponseWriter, r *http.Request) {
	boardID, err := tryGetUUIDParam("boardId", r)
	if err != nil {
		respondError(w, err, http.StatusBadRequest)
		return
	}

	logs, err := h.service.GetLogsByBoard(boardID)
	if err != nil {
		respondError(w, err, http.StatusInternalServerError)
		return
	}

	respondJSON(w, logs, http.StatusOK)
}

func (h *BoardLogHandler) GetLog(w http.ResponseWriter, r *http.Request) {
	id, err := tryGetUUIDParam("id", r)
	if err != nil {
		respondError(w, err, http.StatusBadRequest)
		return
	}

	log, err := h.service.GetLog(id)
	if err != nil {
		respondError(w, err, http.StatusNotFound)
		return
	}

	respondJSON(w, log, http.StatusOK)
}

func (h *BoardLogHandler) DeleteLog(w http.ResponseWriter, r *http.Request) {
	id, err := tryGetUUIDParam("id", r)
	if err != nil {
		respondError(w, err, http.StatusBadRequest)
		return
	}

	if err = h.service.DeleteLog(id); err != nil {
		respondError(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
