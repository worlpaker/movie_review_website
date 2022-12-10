package ErrHandler

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	HTTPStatusCode int
	StatusText     string
	ErrorText      error // only for developer mode
}

//ErrRequest helper function handle error request
func (e *Response) ErrRequest(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(e.HTTPStatusCode)
	msg := map[string]string{
		"message": e.StatusText,
		"error":   e.ErrorText.Error(),
	}
	json.NewEncoder(w).Encode(msg)
	return nil
}
