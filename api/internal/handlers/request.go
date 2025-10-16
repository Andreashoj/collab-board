package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func tryDecodeJSON[T any](body io.ReadCloser) (*T, error) {
	defer body.Close()

	var msg T
	if err := json.NewDecoder(body).Decode(&msg); err != nil {
		return nil, err
	}

	return &msg, nil
}

func tryGetUUIDParam(param string, r *http.Request) (uuid.UUID, error) {
	idStr := chi.URLParam(r, param)
	return uuid.Parse(idStr)
}
