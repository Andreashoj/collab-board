package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func tryDecodeJSON[T any](body io.ReadCloser) (*T, error) {
	defer body.Close()

	var msg T
	if err := json.NewDecoder(body).Decode(&msg); err != nil {
		return nil, err
	}

	return &msg, nil
}

func tryGetUintParam(param string, r *http.Request) (uint, error) {
	idStr := chi.URLParam(r, param)
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		return 0, err
	}

	return uint(id), nil
}
