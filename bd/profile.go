package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/satico0518/twitterGo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Profile(Id string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twitterGO")
	col := db.Collection("users")

	var profile models.User

	objId, _ := primitive.ObjectIDFromHex(Id)

	filter := bson.M{"_id": objId}
	err := col.FindOne(ctx, filter).Decode(&profile)
	profile.Password = ""
	if err != nil {
		fmt.Println("User not found " + err.Error())
		return profile, err
	}
	return profile, nil
}
