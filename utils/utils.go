package utils

import (
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"log"
	"math/rand"
	"os"
	"time"
)

func Getenv(key string, defaultVal string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	val, exist := os.LookupEnv(key)
	if exist {
		return val
	}
	return defaultVal
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("0123456789")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GetSession(c echo.Context) (*sessions.Session, error) {
	return session.Get("SkillSpar", c)
}
