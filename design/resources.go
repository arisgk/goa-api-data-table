package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("user", func() {
	DefaultMedia(UserMedia)
	BasePath("/users")

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("Retrieve all users.")
		Response(OK, CollectionOf(UserMedia))
	})

	Action("show", func() {
		Routing(
			GET("/:userID"),
		)
		Description("Retrieve user with given id.")
		Params(func() {
			Param("userID", String, "User ID")
		})
		Response(OK)
		Response(NotFound)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Create new user")
		Payload(func() {
			Member("firstName")
			Member("lastName")
			Member("age")
			Required("firstName", "lastName")
		})
		Response(Created, UserMedia)
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Routing(
			PUT("/:userID"),
		)
		Description("Update a user")
		Params(func() {
			Param("userID", String, "User ID")
		})
		Payload(func() {
			Member("firstName")
			Member("lastName")
			Member("age")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Routing(
			DELETE("/:userID"),
		)
		Params(func() {
			Param("userID", String, "User ID")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("swagger.json", "public/swagger/swagger.json")
})