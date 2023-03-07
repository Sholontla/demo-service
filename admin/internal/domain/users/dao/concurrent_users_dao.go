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

func ConcurrentRegisterUser(u models.Users, uChan chan<- models.Users) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn := db.MongoConn.Database("admin-demo-service")
	col := conn.Collection("users")

	select {
	case uChan <- u:
	default:
		fmt.Printf("channel is blocked: %v", u)
	}

	hash, err := security.SetPassword(u.Password)
	if err != nil {
		log.Println(err)
	}
	u.Password = string(hash)
	u.ID = primitive.NewObjectID()
	u.CreatedAt = time.Now().GoString()

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		fmt.Printf("failed to execute query: %v", err)
	}
	fmt.Print(result.InsertedID)
	go func() {
		// send the updated store information over the channel without blocking
		select {
		case uChan <- u:
		default:
			fmt.Printf("channel is blocked: %v", u)
		}
		close(uChan)
	}()
}

func ConcurrentFindUserByEmail(email string, emailStringChan chan<- string, usersChan chan<- models.Users) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn := db.MongoConn.Database("admin-demo-service")
	col := conn.Collection("users")

	select {
	case emailStringChan <- email:
	default:
		fmt.Printf("channel is blocked: %v", email)
	}

	var profile models.Users

	condition := bson.M{
		"user_email": email,
	}

	err := col.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""
	if err != nil {
		log.Println("No User found ...", err.Error())
	}
	go func() {
		// Send the updated store information over the channel without blocking
		select {
		case emailStringChan <- profile.UserEmail:
		default:
			fmt.Printf("channel is blocked: %v", profile.UserEmail)
		}
		close(emailStringChan)
	}()

	go func() {
		// Send the updated store information over the channel without blocking
		select {
		case usersChan <- profile:
		default:
			fmt.Printf("channel is blocked: %v", profile)
		}
		close(usersChan)
	}()

}

func ConcurrentUpdateUser(u models.Users, email string) (bool, error) {
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

func ConcurrentUpdateUserPassword(u models.Users, email string) (bool, error) {
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

func ConcurrentDeleteUser(email string) (bool, error) {
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
