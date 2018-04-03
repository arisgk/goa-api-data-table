package mongodb

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/connstring"
)

// CreateClient creates a new Mongo client
func CreateClient(uri string) (client *mongo.Client, err error) {
	connString, err := connstring.Parse(uri)
	if err != nil {
		return nil, err
	}

	cl, err := mongo.NewClientFromConnString(connString)

	return cl, err
}
