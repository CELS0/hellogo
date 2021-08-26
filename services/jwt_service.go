package services

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtService struct {
	secreteKey string
	issure     string
}

func NewJWTService() *jwtService {
	return &jwtService{
		secreteKey: "secret-key",
		issure:     "book-api",
	}
}

type Claim struct {
	Sum uint `json: "sum"`
	jwt.StandardClaims
}

func (s *jwtService) GenerateToken(id uint) (string, error) {
	clain := &Claim{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    s.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clain)

	t, err := token.SignedString([]byte(s.secreteKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (s *jwtService) ValidadeToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", token)
		}
		return []byte(s.secreteKey), nil
	})
	return err == nil
}
