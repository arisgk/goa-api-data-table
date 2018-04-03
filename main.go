//go:generate goagen bootstrap -d github.com/arisgk/goa-api-data-table/design

package main

import (
	"fmt"

	"github.com/arisgk/goa-api-data-table/app"
	"github.com/arisgk/goa-api-data-table/store/mongodb"
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

	// Initialize Store
	store, err := mongodb.CreateStore("mongodb://api:Starl1ght@ds231559.mlab.com:31559/goa-simple-crud")
	if err != nil {
		service.LogError("startup", "err", err)
	}

	fmt.Printf("%v", store)

	// Mount "swagger" controller
	c := NewSwaggerController(service)
	app.MountSwaggerController(service, c)
	// Mount "user" controller
	c2 := NewUserController(service, store)
	app.MountUserController(service, c2)

	// Start service
	if err := service.ListenAndServe(":5000"); err != nil {
		service.LogError("startup", "err", err)
	}

}
