package User

import (
	MongoDB "backend/database/mongo"
	RedisDB "backend/database/redis"
	Models "backend/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func SetupDB() error {
	if err := MongoDB.Mongoclient(); err != nil {
		return err
	}
	if err := RedisDB.Redisclient(); err != nil {
		return err
	}
	return nil
}

func ExecuteRequest(r *http.Request, s *chi.Mux) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	s.ServeHTTP(w, r)
	return w
}

type ModelsGeneric interface {
	*Models.Account | *Models.ChangePassword | *Models.UpdateAccount
}

func ConvertJsontoBuf[T ModelsGeneric](t *testing.T, data T) *bytes.Buffer {
	databuf := new(bytes.Buffer)
	if err := json.NewEncoder(databuf).Encode(data); err != nil {
		t.Errorf("error in convert json to buf %v", err)
	}
	return databuf
}

func TestLogin(t *testing.T) {
	log.SetOutput(io.Discard)
	db := SetupDB()
	assert.Empty(t, db)
	router := chi.NewRouter()
	router.Post("/api/login", Login)
	TestData := []struct {
		data          *Models.Account
		expected_code int
	}{
		{data: &Models.Account{Email: "test@test.com", Password: "test123"},
			expected_code: http.StatusOK},
		{data: &Models.Account{Email: "test@test.com", Password: "test123", First_name: "test", Last_name: "test", Birth_date: "01.01.1900", Gender: "male"},
			expected_code: http.StatusOK},
	}
	for i, k := range TestData {
		databuf := ConvertJsontoBuf(t, k.data)
		_, err := DB.Mongo.CreateUser(k.data)
		assert.Empty(t, err)
		t.Run(fmt.Sprintln("no: ", i), func(t *testing.T) {
			req, _ := http.NewRequest("POST", "/api/login", databuf)
			response := ExecuteRequest(req, router)
			assert.Equal(t, http.StatusOK, response.Code)
		})
		remove_data := &Models.DeleteDataModel{CollName: "user", Filter: "email", Data: k.data.Email}
		err = DB.Mongo.DeleteData(remove_data)
		assert.Empty(t, err)
	}
}
