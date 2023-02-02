package bd

import (
	"github.com/satico0518/twitterGo/models"
	"golang.org/x/crypto/bcrypt"
)

func TryLogin(email string, pass string) (models.User, bool) {
	user, found, _ := UserExists(email)
	if !found {
		return user, false
	}
	passwordBytes := []byte(pass)
	dbPassBytes := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(dbPassBytes, passwordBytes)
	if err != nil {
		return user, false
	}
	return user, true
}
