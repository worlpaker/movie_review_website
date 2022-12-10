package MongoDB

import (
	Models "backend/models"
	"fmt"
	"io"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Rather than mocking the mongo,
It creates real variables in mongo and then deletes them.
*/

func TestLogin(t *testing.T) {
	log.SetOutput(io.Discard)
	if db := Mongoclient(); db != nil {
		t.Errorf("error in db connection %v", db)
	}
	Handle := NewAuth()
	create_data := &Models.Account{Email: "test@test.com", Password: "test123"}
	_, err := Handle.CreateUser(create_data)
	assert.Empty(t, err)
	login_data := &Models.Account{Email: "test@test.com", Password: "test123"}
	_, err = Handle.Login(login_data)
	assert.Empty(t, err)
	remove_data := &Models.DeleteDataModel{CollName: "user", Filter: "email", Data: create_data.Email}
	err = Handle.DeleteData(remove_data)
	assert.Empty(t, err)

}

func TestCreateUser(t *testing.T) {
	log.SetOutput(io.Discard)
	Handle := NewAuth()
	if db := Mongoclient(); db != nil {
		t.Errorf("error in db connection %v", db)
	}
	TestData := []struct {
		data            *Models.Account
		expected_output error
	}{
		{&Models.Account{Email: "test@test.com", Password: "test123"}, nil},
		{&Models.Account{Email: "test@test.com", Password: "test123"}, fmt.Errorf("user already exist")},
	}

	for i, k := range TestData {
		t.Run(fmt.Sprintln("no:", i), func(t *testing.T) {
			_, err := Handle.CreateUser(k.data)
			assert.Equal(t, k.expected_output, err)
		})
	}
	remove_data := &Models.DeleteDataModel{CollName: "user", Filter: "email", Data: TestData[0].data.Email}
	err := Handle.DeleteData(remove_data)
	assert.Empty(t, err)
}

func TestChangePassword(t *testing.T) {
	log.SetOutput(io.Discard)
	if db := Mongoclient(); db != nil {
		t.Errorf("error in db connection %v", db)
	}
	Handle := NewAuth()
	create_data := &Models.Account{Email: "test@test.com", Password: "test123"}
	_, err := Handle.CreateUser(create_data)
	assert.Empty(t, err)
	change_data := &Models.ChangePassword{Email: "test@test.com", Old_Password: "test123", New_Password: "test456"}
	err = Handle.ChangePassword(change_data)
	assert.Empty(t, err)
	remove_data := &Models.DeleteDataModel{CollName: "user", Filter: "email", Data: create_data.Email}
	err = Handle.DeleteData(remove_data)
	assert.Empty(t, err)

}

func TestUpdateProfile(t *testing.T) {
	log.SetOutput(io.Discard)
	if db := Mongoclient(); db != nil {
		t.Errorf("error in db connection %v", db)
	}
	Handle := NewAuth()
	create_data := &Models.Account{Email: "test@test.com", Password: "test123"}
	_, err := Handle.CreateUser(create_data)
	assert.Empty(t, err)
	change_data := &Models.UpdateAccount{Email: "test@test.com", First_name: "Test", Last_name: "Test"}
	_, err = Handle.UpdateProfile(change_data)
	assert.Empty(t, err)
	remove_data := &Models.DeleteDataModel{CollName: "user", Filter: "email", Data: create_data.Email}
	err = Handle.DeleteData(remove_data)
	assert.Empty(t, err)
}
