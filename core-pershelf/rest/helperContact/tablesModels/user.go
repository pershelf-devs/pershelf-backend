package tablesModels

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username    string             `bson:"username" json:"username"`
	Password    string             `bson:"password" json:"password"`
	Email       string             `bson:"email" json:"email"`
	Age         int                `bson:"age" json:"age"`
	Phone       string             `bson:"phone" json:"phone"`
	Description string             `bson:"description" json:"description"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
	Name        string             `bson:"name" json:"name"`
	Surname     string             `bson:"surname" json:"surname"`
}

func (u *User) Save() error {
	// MongoDB bağlantısı
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Printf("Error connecting to MongoDB: %v", err)
		return err
	}
	defer client.Disconnect(ctx)

	// Zaman damgalarını ayarla
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	// Collection'a kaydet
	collection := client.Database("pershelf").Collection("users")
	_, err = collection.InsertOne(ctx, u)
	if err != nil {
		log.Printf("Error saving user: %v", err)
		return err
	}

	return nil
}
