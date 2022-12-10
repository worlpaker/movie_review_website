package RedisDB

import (
	config "backend/config"
	ErrHandler "backend/helper_handlers/error"
	Models "backend/models"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

// SaveToken saves token metadata to Redis
func (s *service) SaveToken(refresh_token string, td *Models.TokenDetails) error {
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()
	atCreated, err := client.Set(refresh_token, "Refresh_Token: "+refresh_token+"\n Inserted_Time: "+now.String(), rt.Sub(now)).Result()
	if ErrHandler.Log(err) {
		return err
	}
	if atCreated == "0" {
		err = errors.New("no record inserted")
		ErrHandler.Log(err)
		return err
	}
	return nil
}

// CreateToken for given data if save_token is true saves them to redis
func (s *service) CreateToken(u *Models.Token) (*Models.TokenDetails, error) {
	var err error
	td := &Models.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Hour * 24).Unix()     //expires after 24 hour
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix() //expires after 7 day
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["Email"] = u.Email
	atClaims["First_name"] = u.First_name
	atClaims["Last_name"] = u.Last_name
	atClaims["Birth_Date"] = u.Birth_date
	atClaims["Gender"] = u.Gender
	atClaims["Profile_picture"] = u.Pp_id
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(config.Access_Secret))
	if ErrHandler.Log(err) {
		return nil, err
	}
	//Creating Refresh Token
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	rtClaims := jwt.MapClaims{}
	rtClaims["Email"] = u.Email
	rtClaims["First_name"] = u.First_name
	rtClaims["Last_name"] = u.Last_name
	rtClaims["Birth_Date"] = u.Birth_date
	rtClaims["Gender"] = u.Gender
	rtClaims["Profile_picture"] = u.Pp_id
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(config.Refresh_Secret))
	if ErrHandler.Log(err) {
		return nil, err
	}
	if u.Save_Token {
		if err := s.SaveToken(td.RefreshToken, td); ErrHandler.Log(err) {
			return nil, err
		}
	}
	return td, nil
}

// DeleteToken deletes old token in redis
func (s *service) DeleteToken(refresh_token string) error {
	deleted, err := client.Del(refresh_token).Result()
	if ErrHandler.Log(err) || deleted == 0 {
		return err
	}
	return nil
}

// TokenValid check if token is expired and valid, lastly verify token
func (s *service) TokenValid(r *http.Request) (*jwt.Token, error) {
	token, err := s.VerifyToken(r)
	if ErrHandler.Log(err) {
		return nil, err
	}
	if _, ok := token.Claims.(*jwt.StandardClaims); !ok && !token.Valid {
		ErrHandler.Log(fmt.Errorf("error in tokenvalid"))
		return nil, fmt.Errorf("error in tokenvalid")
	}
	return token, nil
}

// VerifyToken verify token method
func (s *service) VerifyToken(r *http.Request) (*jwt.Token, error) {
	refreshtoken, _ := r.Cookie("Token")
	tokenString := refreshtoken.Value
	if err := s.CheckToken(tokenString); ErrHandler.Log(err) {
		return nil, err
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			ErrHandler.Log(fmt.Errorf("unexpected signing method: %v", token.Header["alg"]))
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Refresh_Secret), nil
	})
	if ErrHandler.Log(err) {
		return nil, err
	}
	return token, nil
}

// CheckToken checks if refresh token expired
func (s *service) CheckToken(refresh_token string) error {
	_, err := client.Get(refresh_token).Result()
	if ErrHandler.Log(err) {
		return err
	}
	return nil
}

// ReadToken read jwt token and converts to *models.Token
func (s *service) ReadToken(token *jwt.Token) (*Models.Token, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		ErrHandler.Log(fmt.Errorf("error reading token"))
		return nil, fmt.Errorf("error reading token")
	}
	email, _ := claims["Email"].(string)
	firstname, _ := claims["First_name"].(string)
	lastname, _ := claims["Last_name"].(string)
	birthdate, _ := claims["Birth_Date"].(string)
	gender, _ := claims["Gender"].(string)
	pp_id, _ := claims["Profile_picture"].(string)
	data := &Models.Token{Save_Token: false,
		Email: email, First_name: firstname, Last_name: lastname, Birth_date: birthdate, Gender: gender, Pp_id: pp_id}
	return data, nil
}
