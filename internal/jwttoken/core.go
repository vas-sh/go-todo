package jwttoken

import (
	"encoding/json"

	"github.com/golang-jwt/jwt/v5"
	"github.com/vas-sh/todo/internal/models"
)

type srv struct {
	secretJWT string
}

func New(secretJWT string) *srv {
	return &srv{secretJWT: secretJWT}
}

func (s *srv) validateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(t *jwt.Token) (any, error) {
		_, isValid := t.Method.(*jwt.SigningMethodHMAC)
		if !isValid {
			return nil, models.ErrInvalidToken
		}
		return []byte(s.secretJWT), nil
	})
}

func (s *srv) GetUser(auth string) (*models.User, error) {
	if len(auth) < 5 {
		return nil, models.ErrInvalidToken
	}
	token, err := s.validateToken(auth[4:])
	if err != nil {
		return nil, err
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, models.ErrInvalidToken
	}
	userStr, ok := claim["user"].(string)
	if !ok {
		return nil, models.ErrInvalidToken
	}
	var user models.User
	err = json.Unmarshal([]byte(userStr), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
