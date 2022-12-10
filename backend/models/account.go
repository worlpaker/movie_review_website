package Models

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	Base
	Email           string    `json:"Email"`
	Password        string    `json:"Password"`
	First_name      string    `json:"First_name"`
	Last_name       string    `json:"Last_name"`
	Birth_date      string    `json:"Birth_date"`
	Gender          string    `json:"Gender"`
	Profile_picture uuid.UUID `json:"Profile_picture"`
}

type ChangePassword struct {
	Email        string `json:"Email"`
	Old_Password string `json:"Old_Password"`
	New_Password string `json:"New_Password"`
}

type UpdateAccount struct {
	Email      string    `json:"Email"`
	UpdatedAt  time.Time `json:"UpdatedAt"`
	First_name string    `json:"First_name"`
	Last_name  string    `json:"Last_name"`
	Birth_date string    `json:"Birth_date"`
	Gender     string    `json:"Gender"`
}
