package globals

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	MongoClient     *mongo.Client
	UsersCollection *mongo.Collection
)

func InitCollections() {
	UsersCollection = MongoClient.Database("pershelf").Collection("users")
}
