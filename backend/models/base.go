package Models

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	Id            uuid.UUID  `json:"Id"`
	CreatedAt     time.Time  `json:"CreatedAt"`
	UpdatedAt     time.Time  `json:"UpdatedAt"`
	DeletedAt     *time.Time `json:"DeletedAt"`
	Authorization string     `json:"Authorization"`
}