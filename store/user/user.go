package user

import (
	"github.com/arisgk/goa-api-data-table/entities/user"
	"github.com/satori/go.uuid"
)

// Store defines user store behavior
type Store interface {
	Create() uuid.UUID
	Read() user.User
	List() []user.User
	Update() user.User
	Delete() uuid.UUID
}
