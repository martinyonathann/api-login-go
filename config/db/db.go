package db

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	password = "SXGz0TDpZRVOCgGP"
	dbname   = "GoLogin"
)

func GetDBCollection() (*mongo.Collection, error) {
	log.SetFormatter(&log.JSONFormatter{})

	// client, err := mongo.NewClient((options.Client().ApplyURI("mongodb://localhost:27017")))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://standard:"+password+"@cluster0.pdpui.mongodb.net/"+dbname+"?retryWrites=true&w=majority",
	))

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	//Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	collection := client.Database("GoLogin").Collection("users")
	return collection, nil
}
