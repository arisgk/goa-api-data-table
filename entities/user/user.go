package user

import (
	"github.com/satori/go.uuid"
)

type User struct {
	Id uuid.UUID
	FirstName string
	LastName string
	Age int
}