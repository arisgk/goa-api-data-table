package user

import (
	"context"
	"log"

	"github.com/arisgk/goa-api-data-table/databases/mongodb"
	"github.com/arisgk/goa-api-data-table/entities/user"
	uuid "github.com/satori/go.uuid"
)

// Create inserts a user record into MongoDB
func Create(user user.User) {
	collection := mongodb.Database.Collection("users")

	res, err := collection.InsertOne(context.Background(), user)

	if err != nil {
		log.Fatal(err)
	}

	return user.ID
}

// Read gets a user record from MongoDB by ID
func Read(id uuid.UUID) {

}

// List gets a list of all user records in MongoDB
func List() {

}

// Update updates a user record in MongoDB
func Update() {

}

// Delete deletes a user record from MongoDB
func Update() {

}
