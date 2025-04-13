package adapters

import "go.mongodb.org/mongo-driver/mongo"

type UserAdapter struct {
	MongoDatabase *mongo.Database
	collection    string
}

func NewUserAdapter(mongoDB *mongo.Database) *UserAdapter {
	return &UserAdapter{
		MongoDatabase: mongoDB,
		collection:    "user",
	}
}
