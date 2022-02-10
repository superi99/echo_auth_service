package utils

import (
	"github.com/golang-jwt/jwt"
	"time"
)

var JWTSecret = []byte(Getenv("secret_jwt", "!!SECRET!!"))

func GenerateJWT(id uint) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, _ := token.SignedString(JWTSecret)
	return t
}
