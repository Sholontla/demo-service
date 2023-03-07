package dao

import (
	"context"
	"fmt"
	"log"
	"service/admin/demo/internal/db"
	"service/admin/demo/internal/domain/users/models"
	"service/admin/demo/internal/security"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterUser(u models.Users) (models.Users, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn := db.MongoConn.Database("admin-demo-service")
	col := conn.Collection("users")

	hash, err := security.SetPassword(u.Password)
	if err != nil {
		log.Println(err)
	}
	u.Password = string(hash)
	u.ID = primitive.NewObjectID()
	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return u, false, err
	}
	fmt.Print(result.InsertedID)
	return u, true, err
}

func ExistsUser(email string) (models.Users, bool, string) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	conn := db.MongoConn.Database("admin-demo-service")
	col := conn.Collection("users")

	condition := bson.M{"email": email}

	var result models.Users

	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}

func FindUserByEmail(email string) (models.Users, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn := db.MongoConn.Database("admin-demo-service")
	col := conn.Collection("users")

	var profile models.Users

	condition := bson.M{
		"user_email": email,
	}

	err := col.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""
	if err != nil {
		log.Println("No User found ...", err.Error())
	}

	return profile, nil
}

func UpdateUser(u models.Users, email string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn := db.MongoConn.Database("admin-demo-service")
	col := conn.Collection("users")

	register := make(map[string]interface{})

	if len(u.UserName) > 0 {
		register["user_name"] = u.UserName
	}
	if len(u.UserLastName) > 0 {
		register["user_last_name"] = u.UserLastName
	}

	updateData := bson.M{
		"$set": register,
	}

	filter := bson.M{"user_email": bson.M{"$eq": email}}

	_, err := col.UpdateOne(ctx, filter, updateData)
	if err != nil {
		return false, err
	}

	return true, nil
}

func UpdateUserPassword(u models.Users, email string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn := db.MongoConn.Database("admin-demo-service")
	col := conn.Collection("users")

	register := make(map[string]interface{})

	if register["password"] != register["password_confirm"] {
		return false, nil
	}
	if len(u.Password) > 0 {
		register["password"] = u.Password
	}

	updateData := bson.M{
		"$set": register,
	}

	filter := bson.M{"user_email": bson.M{"$eq": email}}

	_, err := col.UpdateOne(ctx, filter, updateData)
	if err != nil {
		return false, err
	}

	return true, nil
}

func DeleteUser(email string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn := db.MongoConn.Database("admin-demo-service")
	col := conn.Collection("users")

	condition := bson.M{
		"user_email": email,
	}

	_, err := col.DeleteOne(ctx, condition)
	if err != nil {
		return false, err
	}

	return true, nil
}
