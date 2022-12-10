package SuccessHandler

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	HTTPStatusCode int
	Message        string
}

func (e *Response) Render(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(e.HTTPStatusCode)
	w.Write([]byte(e.Message))
	return nil
}

func (e *Response) RenderJSON(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(e.HTTPStatusCode)
	msg := map[string]string{
		"message": e.Message,
	}
	json.NewEncoder(w).Encode(msg)
	return nil
}
