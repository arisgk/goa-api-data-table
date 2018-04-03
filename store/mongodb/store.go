package mongodb

import (
	"context"

	"github.com/arisgk/goa-api-data-table/entities"
	"github.com/davecgh/go-spew/spew"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/connstring"
)

// MongoDB struct encapsulates MongoDB connection information
type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

// CreateStore creates a new MongoDB instance
func CreateStore(uri string) (store *MongoDB, err error) {
	connString, err := connstring.Parse(uri)

	spew.Dump(connString)

	if err != nil {
		return nil, err
	}

	client, err := mongo.NewClientFromConnString(connString)

	if err != nil {
		return nil, err
	}

	store = &MongoDB{
		Client:   client,
		Database: client.Database("goa-simple-crud"),
	}

	return store, nil
}

// Create inserts a user record into MongoDB
func (mongo MongoDB) Create(user entities.User) (string, error) {
	collection := mongo.Database.Collection("users")

	_, err := collection.InsertOne(context.Background(), user)

	if err != nil {
		return "", err
	}

	return user.ID, nil
}

// Read gets a user record from MongoDB by ID
func (mongo MongoDB) Read(id string) entities.User {
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
func (mongo MongoDB) Delete(id string) string {
	return id
}
