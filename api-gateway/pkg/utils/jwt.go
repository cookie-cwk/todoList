package utils

import (
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)
var jwtSecret = []byte("TodoList")


type Claims struct {
	Id uint `json:"id"`
	jwt.StandardClaims
}

func GenerateToken(id uint) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24*time.Hour)
	claims := Claims {
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: "cwk",
			IssuedAt: nowTime.Unix(),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	if strings.HasPrefix(token, "Bearer ") {
		token = token[7:]
	}
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok&&tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
