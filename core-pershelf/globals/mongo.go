package globals

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	MongoClient         *mongo.Client
	UsersCollection     *mongo.Collection
	BooksCollection     *mongo.Collection
	UserBooksCollection *mongo.Collection
)

func InitCollections() {
	UsersCollection = MongoClient.Database("pershelf").Collection("users")
	BooksCollection = MongoClient.Database("pershelf").Collection("books")
	UserBooksCollection = MongoClient.Database("pershelf").Collection("user_books")
}
