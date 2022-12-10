package RedisDB

import (
	Models "backend/models"
	"net/http"

	"github.com/golang-jwt/jwt"
)

type service struct{}

func NewAuth() *service {
	return &service{}
}

type IHandlers interface {
	SaveToken(string, *Models.TokenDetails) error
	CreateToken(*Models.Token) (*Models.TokenDetails, error)
	DeleteToken(string) error
	TokenValid(*http.Request) (*jwt.Token, error)
	VerifyToken(*http.Request) (*jwt.Token, error)
	CheckToken(string) error
	ReadToken(*jwt.Token) (*Models.Token, error)
}

var _ IHandlers = &service{}
