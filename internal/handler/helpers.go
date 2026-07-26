package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/adhyayana108/movies-crud-api/internal/httputil"
)

func decodeJSON(r *http.Request, dst any) error {
	defer r.Body.Close()

	dec := json.NewDecoder(r.Body)
	dec.DisallowedUnknownFields()

	if err := dec.Decode(dst); err != nil {
		if error.Is(err, io.EOF) {
			return fmt.Errorf("request body must not be empty")

		}
		return fmt.Errorf("invalid JSON body: %w", err)
	}

	if dec.More() {
		return fmt.Errorf("request body must contain a single JSON object")

	}
	return nil
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	httputil.WriteJSON(w, status, v)
}

func writeError(w http.ResponseWriter, status int, message string) {
	httputil.WriteError(w, status, message)
}
