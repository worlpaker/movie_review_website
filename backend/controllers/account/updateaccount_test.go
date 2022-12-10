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

func TestChangePassword(t *testing.T) {
	log.SetOutput(io.Discard)
	db := SetupDB()
	assert.Empty(t, db)
	router := chi.NewRouter()
	router.Use(UsersOnly)
	router.Post("/api/profile/changepassword", ChangePassword)
	CreateUserData := &Models.Account{Email: "test@test.com", Password: "test123"}
	_, err := DB.Mongo.CreateUser(CreateUserData)
	assert.Empty(t, err)
	CreateTokenData := &Models.Token{Save_Token: true, Email: "test@test.com", First_name: "test", Last_name: "test", Birth_date: "01.01.1900", Gender: "male"}
	ts, err := DB.Redis.CreateToken(CreateTokenData)
	assert.Empty(t, err)
	cookie := &http.Cookie{
		Name:     "Token",
		Value:    ts.RefreshToken,
		HttpOnly: false,
		MaxAge:   int(time.Hour * 24 * 3),
		Path:     "/",
	}
	ChangeData := &Models.ChangePassword{Email: "test@test.com", Old_Password: "test123", New_Password: "test456"}
	databuf := ConvertJsontoBuf(t, ChangeData)
	req, _ := http.NewRequest("POST", "/api/profile/changepassword", databuf)
	req.AddCookie(cookie)
	response := ExecuteRequest(req, router)
	assert.Equal(t, http.StatusOK, response.Code)
	remove_data := &Models.DeleteDataModel{CollName: "user", Filter: "email", Data: CreateUserData.Email}
	err = DB.Mongo.DeleteData(remove_data)
	assert.Empty(t, err)
	err = DB.Redis.DeleteToken(ts.RefreshToken)
	assert.Empty(t, err)
}

func TestUpdateProfile(t *testing.T) {
	log.SetOutput(io.Discard)
	db := SetupDB()
	assert.Empty(t, db)
	router := chi.NewRouter()
	router.Use(UsersOnly)
	router.Post("/api/profile/updateprofile", UpdateProfile)
	CreateUserData := &Models.Account{Email: "test@test.com", Password: "test123"}
	_, err := DB.Mongo.CreateUser(CreateUserData)
	assert.Empty(t, err)
	CreateTokenData := &Models.Token{Save_Token: true, Email: "test@test.com"}
	ts, err := DB.Redis.CreateToken(CreateTokenData)
	assert.Empty(t, err)
	cookie := &http.Cookie{
		Name:     "Token",
		Value:    ts.RefreshToken,
		HttpOnly: false,
		MaxAge:   int(time.Hour * 24 * 3),
		Path:     "/",
	}
	ChangeData := &Models.UpdateAccount{Email: "test@test.com", First_name: "Test", Last_name: "Test"}
	databuf := ConvertJsontoBuf(t, ChangeData)
	req, _ := http.NewRequest("POST", "/api/profile/updateprofile", databuf)
	req.AddCookie(cookie)
	response := ExecuteRequest(req, router)
	assert.Equal(t, http.StatusOK, response.Code)
	remove_data := &Models.DeleteDataModel{CollName: "user", Filter: "email", Data: CreateUserData.Email}
	err = DB.Mongo.DeleteData(remove_data)
	assert.Empty(t, err)
	err = DB.Redis.DeleteToken(ts.RefreshToken)
	assert.Empty(t, err)
}
