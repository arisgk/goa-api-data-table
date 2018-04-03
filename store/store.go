package store

import (
	"github.com/arisgk/goa-api-data-table/entities"
	"github.com/satori/go.uuid"
)

// Store defines user store behavior
type Store interface {
	Create(user entities.User) uuid.UUID
	Read(id uuid.UUID) entities.User
	List() []entities.User
	Update(user entities.User) entities.User
	Delete(id uuid.UUID) uuid.UUID
}
