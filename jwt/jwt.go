package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/satico0518/twitterGo/models"
)

func GetToken(user models.User) (string, error) {
	secret := []byte("secretForJWTinGO")
	payload := jwt.MapClaims{
		"ID":       user.ID.Hex(),
		"expires":  time.Now().Add(24 * time.Hour).Unix(),
		"email":    user.Email,
		"lastname": user.LastName,
		"birthday": user.Birthday,
		"bio":      user.Bio,
		"banner":   user.Banner,
		"location": user.Location,
		"web":      user.Website,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}
