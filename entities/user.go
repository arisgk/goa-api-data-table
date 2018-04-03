package entities

import (
	"github.com/satori/go.uuid"
)

// User entity. Encapsulates information about a user.
type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Age       int
}
