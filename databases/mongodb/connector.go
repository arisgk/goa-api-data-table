package mongodb

import (
	"log"

	"github.com/mongodb/mongo-go-driver/mongo"
)

// MongoDB struct encapsulates mongo connection data
type MongoDB struct {
	Database *mongo.Database
}

// Connect initiates the connection to MongoDB
func (mongoDB MongoDB) Connect() {
	client, err := mongo.NewClient("mongodb://api:St@rl1ght!@ds231559.mlab.com:31559/goa-simple-crud")
	if err != nil {
		log.Fatal(err)
	}

	mongoDB.Database = client.Database("goa-simple-crud")
}
