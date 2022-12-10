package RedisDB

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

func TestSaveToken(t *testing.T) {
	log.SetOutput(io.Discard)
	err := Redisclient()
	assert.Empty(t, err)
	Handler := NewAuth()
	TestData := struct {
		data *Models.Token
	}{
		data: &Models.Token{Save_Token: false, Email: "test@test.com", First_name: "test", Last_name: "test", Birth_date: "01.01.1900", Gender: "male"},
	}
	ts, err := Handler.CreateToken(TestData.data)
	assert.Empty(t, err)
	err = Handler.SaveToken(ts.RefreshToken, ts)
	assert.Empty(t, err)
	err = Handler.DeleteToken(ts.RefreshToken)
	assert.Empty(t, err)
}

func TestCreateToken(t *testing.T) {
	log.SetOutput(io.Discard)
	err := Redisclient()
	assert.Empty(t, err)
	Handler := NewAuth()
	TestData := struct {
		data *Models.Token
	}{
		data: &Models.Token{Save_Token: true, Email: "test@test.com", First_name: "test", Last_name: "test", Birth_date: "01.01.1900", Gender: "male"},
	}
	ts, err := Handler.CreateToken(TestData.data)
	assert.Empty(t, err)
	err = Handler.DeleteToken(ts.RefreshToken)
	assert.Empty(t, err)
}

func TestDeleteToken(t *testing.T) {
	log.SetOutput(io.Discard)
	err := Redisclient()
	assert.Empty(t, err)
	Handler := NewAuth()
	TestData := struct {
		data *Models.Token
	}{
		data: &Models.Token{Save_Token: true, Email: "test@test.com", First_name: "test", Last_name: "test", Birth_date: "01.01.1900", Gender: "male"},
	}
	ts, err := Handler.CreateToken(TestData.data)
	assert.Empty(t, err)
	err = Handler.DeleteToken(ts.RefreshToken)
	assert.Empty(t, err)
}

// Test TokenValid, CheckToken, VerifyToken and ReadToken
func TestTokenValid(t *testing.T) {
	log.SetOutput(io.Discard)
	db := Redisclient()
	assert.Empty(t, db)
	Handler := NewAuth()
	router := chi.NewRouter()
	router.Get("/", nil)
	TestData := struct {
		data *Models.Token
	}{
		data: &Models.Token{Save_Token: true, Email: "test@test.com", First_name: "test", Last_name: "test", Birth_date: "01.01.1900", Gender: "male"},
	}

	ts, err := Handler.CreateToken(TestData.data)
	assert.Empty(t, err)
	cookie := &http.Cookie{
		Name:     "Token",
		Value:    ts.RefreshToken,
		HttpOnly: false,
		MaxAge:   int(time.Hour * 24 * 3),
		Path:     "/",
	}
	req, _ := http.NewRequest("GET", "/", nil)
	req.AddCookie(cookie)
	token, err := Handler.TokenValid(req)
	assert.Empty(t, err)
	_, err = Handler.ReadToken(token)
	assert.Empty(t, err)
	err = Handler.DeleteToken(ts.RefreshToken)
	assert.Empty(t, err)
}
