//go:generate goagen bootstrap -d github.com/arisgk/goa-api-data-table/design

package main

import (
	"github.com/arisgk/goa-api-data-table/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("Data Table")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "swagger" controller
	c := NewSwaggerController(service)
	app.MountSwaggerController(service, c)
	// Mount "user" controller
	c2 := NewUserController(service)
	app.MountUserController(service, c2)

	// Start service
	if err := service.ListenAndServe(":5000"); err != nil {
		service.LogError("startup", "err", err)
	}

}
