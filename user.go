package main

import (
	"fmt"

	"github.com/arisgk/goa-api-data-table/app"
	"github.com/arisgk/goa-api-data-table/entities"
	"github.com/arisgk/goa-api-data-table/store"
	"github.com/goadesign/goa"
	"github.com/satori/go.uuid"
)

// ToUserMedia converts a user entity to goa UserMedia
func ToUserMedia(user entities.User) *app.User {
	return &app.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Age:       &user.Age,
	}
}

// UserController implements the user resource.
type UserController struct {
	*goa.Controller
	store store.Store
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service, store store.Store) *UserController {
	return &UserController{Controller: service.NewController("UserController"), store: store}
}

// Create runs the create action.
func (c *UserController) Create(ctx *app.CreateUserContext) error {
	// UserController_Create: start_implement

	payload := ctx.Payload
	user := entities.User{
		ID:        uuid.Must(uuid.NewV4()).String(),
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
	}

	if payload.Age != nil {
		user.Age = *payload.Age
	}

	fmt.Println(c.store)

	_, err := c.store.Create(user)

	if err != nil {
		return goa.ErrInternal(err)
	}

	return ctx.Created(ToUserMedia(user))
	// UserController_Create: end_implement
}

// Delete runs the delete action.
func (c *UserController) Delete(ctx *app.DeleteUserContext) error {
	// UserController_Delete: start_implement

	// Put your logic here

	return nil
	// UserController_Delete: end_implement
}

// List runs the list action.
func (c *UserController) List(ctx *app.ListUserContext) error {
	// UserController_List: start_implement

	// Put your logic here

	res := app.UserCollection{}
	return ctx.OK(res)
	// UserController_List: end_implement
}

// Show runs the show action.
func (c *UserController) Show(ctx *app.ShowUserContext) error {
	// UserController_Show: start_implement

	// Put your logic here

	res := &app.User{}
	return ctx.OK(res)
	// UserController_Show: end_implement
}

// Update runs the update action.
func (c *UserController) Update(ctx *app.UpdateUserContext) error {
	// UserController_Update: start_implement

	// Put your logic here

	return nil
	// UserController_Update: end_implement
}
