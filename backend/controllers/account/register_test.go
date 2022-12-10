package User

import (
	Models "backend/models"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	log.SetOutput(io.Discard)
	db := SetupDB()
	assert.Empty(t, db)
	router := chi.NewRouter()
	router.Get("/api/register", CreateUser)

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
		t.Run(fmt.Sprintln("no: ", i), func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/api/register", databuf)
			response := ExecuteRequest(req, router)
			assert.Equal(t, http.StatusOK, response.Code)
		})
		remove_data := &Models.DeleteDataModel{CollName: "user", Filter: "email", Data: k.data.Email}
		err := DB.Mongo.DeleteData(remove_data)
		assert.Empty(t, err)
	}
}
