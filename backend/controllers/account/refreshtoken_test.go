package User

import (
	Models "backend/models"
	"io"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestRefresh(t *testing.T) {
	log.SetOutput(io.Discard)
	db := SetupDB()
	assert.Empty(t, db)
	router := chi.NewRouter()
	router.Use(UsersOnly)
	router.Post("/api/profile/refresh", Refresh)
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
	req, _ := http.NewRequest("POST", "/api/profile/refresh", nil)
	req.AddCookie(cookie)
	response := ExecuteRequest(req, router)
	assert.Equal(t, http.StatusOK, response.Code)
	err = DB.Redis.DeleteToken(ts.RefreshToken)
	assert.Empty(t, err)
}
