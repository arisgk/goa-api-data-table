package mongodb

import (
	"context"
	"log"

	"github.com/arisgk/goa-api-data-table/entities"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/connstring"
	uuid "github.com/satori/go.uuid"
)

// MongoDB struct encapsulates MongoDB connection information
type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

// CreateStore creates a new MongoDB instance
func CreateStore(uri string) (store *MongoDB, err error) {
	connString, err := connstring.Parse(uri)
	if err != nil {
		return nil, err
	}

	client, err := mongo.NewClientFromConnString(connString)

	if err != nil {
		return nil, err
	}

	store = &MongoDB{
		Client:   client,
		Database: client.Database("users"),
	}

	return store, nil
}

// Create inserts a user record into MongoDB
func (mongo MongoDB) Create(user entities.User) uuid.UUID {
	collection := mongo.Database.Collection("users")

	_, err := collection.InsertOne(context.Background(), user)

	if err != nil {
		log.Fatal(err)
	}

	return user.ID
}

// Read gets a user record from MongoDB by ID
func (mongo MongoDB) Read(id uuid.UUID) entities.User {
	return entities.User{}
}

// List gets a list of all user records in MongoDB
func (mongo MongoDB) List() []entities.User {
	return []entities.User{}
}

// Update updates a user record in MongoDB
func (mongo MongoDB) Update(user entities.User) entities.User {
	return entities.User{}
}

// Delete deletes a user record from MongoDB
func (mongo MongoDB) Delete(id uuid.UUID) uuid.UUID {
	return id
}
