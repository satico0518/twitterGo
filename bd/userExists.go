package bd

import (
	"context"
	"time"

	"github.com/satico0518/twitterGo/models"
	"go.mongodb.org/mongo-driver/bson"
)

func UserExists(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitterGO")
	col := db.Collection("users")

	filter := bson.M{"email": email}

	var user models.User

	err := col.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return user, false, ""
	}
	ID := user.ID.Hex()
	return user, true, ID
}
