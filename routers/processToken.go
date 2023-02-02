package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/satico0518/twitterGo/bd"
	"github.com/satico0518/twitterGo/models"
)

var Email string
var UserID string

func ProcessToken(token string) (*models.Claims, bool, string, error) {
	secret := []byte("secretForJWTinGO")
	claims := &models.Claims{}

	splitedToken := strings.Split(token, "Bearer")
	if len(splitedToken) < 2 {
		return claims, false, "", errors.New("Invalid Token")
	}
	finalToken := strings.TrimSpace(splitedToken[1])
	tkn, err := jwt.ParseWithClaims(finalToken, claims, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil || !tkn.Valid {
		return claims, false, "", errors.New("Invalid Token")
	}
	_, found, _ := bd.UserExists(claims.Email)
	if found {
		Email = claims.Email
		UserID = claims.ID.Hex()
		return claims, true, UserID, nil
	}
	return claims, false, "", errors.New("User not Found")
}
