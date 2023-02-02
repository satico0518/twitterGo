package bd

import (
	"context"
	"time"

	"github.com/satico0518/twitterGo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitterGO")
	col := db.Collection("users")

	user.Password, _ = encryptPassword(user.Password)

	result, err := col.InsertOne(ctx, user)
	if err != nil {
		return "", false, err
	}

	ObjId, _ := result.InsertedID.(primitive.ObjectID)
	return ObjId.String(), true, nil
}

func encryptPassword(pass string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err
}
