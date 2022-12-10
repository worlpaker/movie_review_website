package User

import (
	ErrHandler "backend/helper_handlers/error"
	Models "backend/models"
	"io"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func UsersOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := r.Cookie("Token"); ErrHandler.Log(err) {
			render := &ErrHandler.Response{HTTPStatusCode: http.StatusUnauthorized, StatusText: http.StatusText(http.StatusUnauthorized), ErrorText: err}
			render.ErrRequest(w, r)
			return
		}
		if _, err := DB.Redis.TokenValid(r); ErrHandler.Log(err) {
			render := &ErrHandler.Response{HTTPStatusCode: http.StatusUnauthorized, StatusText: http.StatusText(http.StatusUnauthorized), ErrorText: err}
			render.ErrRequest(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func TestLogout(t *testing.T) {
	log.SetOutput(io.Discard)
	db := SetupDB()
	assert.Empty(t, db)
	router := chi.NewRouter()
	router.Use(UsersOnly)
	router.Post("/api/profile/logout", Logout)
	data := &Models.Token{Save_Token: true, Email: "test@test.com", First_name: "test", Last_name: "test", Birth_date: "01.01.1900", Gender: "male"}
	ts, err := DB.Redis.CreateToken(data)
	assert.Empty(t, err)
	cookie := &http.Cookie{
		Name:     "Token",
		Value:    ts.RefreshToken,
		HttpOnly: false,
		MaxAge:   int(time.Hour * 24 * 3),
		Path:     "/",
	}
	req, _ := http.NewRequest("POST", "/api/profile/logout", nil)
	req.AddCookie(cookie)
	response := ExecuteRequest(req, router)
	assert.Equal(t, http.StatusOK, response.Code)
}
