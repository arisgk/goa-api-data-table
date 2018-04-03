package store

import (
	"github.com/arisgk/goa-api-data-table/entities"
)

// Store defines user store behavior
type Store interface {
	Create(user entities.User) (string, error)
	Read(id string) entities.User
	List() []entities.User
	Update(user entities.User) entities.User
	Delete(id string) string
}
