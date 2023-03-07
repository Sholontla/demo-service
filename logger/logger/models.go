package logger

import (
	"context"
	"log"
	"time"

	"service/logger/service/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New(mongo *mongo.Client) LogModel {
	client = mongo

	return LogModel{
		LogEntry: LogEntry{},
	}
}

type LogModel struct {
	LogEntry LogEntry
}

type LogEntry struct {
	ID        string `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string `bson:"name" json:"name"`
	Data      string `bson:"data" json:"data"`
	CreatedAt string `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt string `bson:"udated_at,omitempty" json:"udated_at,omitempty"`
}

func (l *LogEntry) Insert(entry LogEntry) error {

	coll := config.DBCollection()
	ctx := context.TODO()
	collection := client.Database(coll).Collection(coll)

	_, err := collection.InsertOne(ctx, LogEntry{
		Name:      entry.Name,
		Data:      entry.Data,
		CreatedAt: GetNowString(),
		UpdatedAt: GetNowString(),
	})
	if err != nil {
		log.Printf("Error while inserting into Log DB: %s", err)
		return nil
	}
	return nil
}

func (l *LogEntry) All() ([]*LogEntry, error) {

	ctxB := context.Background()
	ctxTo := context.TODO()
	ctx, cancel := context.WithTimeout(ctxB, 15*time.Second)
	defer cancel()

	coll := config.DBCollection()
	collection := client.Database(coll).Collection(coll)

	opts := options.Find()
	opts.SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := collection.Find(ctxTo, bson.D{}, opts)

	if err != nil {
		log.Printf("Error While finding all Logs: %s", err)
		return nil, err
	}

	defer cursor.Close(ctx)

	var logs []*LogEntry

	for cursor.Next(ctx) {
		var item LogEntry

		err := cursor.Decode(&item)
		if err != nil {
			log.Printf("Error decoding log into slice: %s", err)
		} else {
			logs = append(logs, &item)
		}
	}

	return logs, nil
}

func (l *LogEntry) GetOne(id string) (*LogEntry, error) {
	ctxB := context.Background()
	ctx, cancel := context.WithTimeout(ctxB, 15*time.Second)
	defer cancel()

	coll := config.DBCollection()
	collection := client.Database(coll).Collection(coll)

	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		Error("Error while getting ID: %s" + err.Error())
	}

	var entry LogEntry
	err = collection.FindOne(ctx, bson.M{"_id": docID}).Decode(&entry)
	if err != nil {
		Error("Error while decoding ID: %s" + err.Error())
	}

	return &entry, nil
}

func (l *LogEntry) DropCollection() error {
	ctxB := context.Background()
	ctx, cancel := context.WithTimeout(ctxB, 15*time.Second)
	defer cancel()

	coll := config.DBCollection()
	collection := client.Database(coll).Collection(coll)

	if err := collection.Drop(ctx); err != nil {
		Error("Error while dropping ID: %s" + err.Error())
	}

	return nil
}

func (l *LogEntry) UpdateCollection() (*mongo.UpdateResult, error) {
	ctxB := context.Background()
	ctx, cancel := context.WithTimeout(ctxB, 15*time.Second)
	defer cancel()

	coll := config.DBCollection()
	collection := client.Database(coll).Collection(coll)

	docID, err := primitive.ObjectIDFromHex(l.ID)
	if err != nil {
		Error("Error while getting ID: %s" + err.Error())
	}

	result, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": docID},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "name", Value: l.Name},
				{Key: "data", Value: l.Data},
				{Key: "updated_at", Value: GetNowString()},
			}},
		},
	)

	if err != nil {
		Error("Error while Updating ID: %s" + err.Error())
	}
	return result, nil
}
