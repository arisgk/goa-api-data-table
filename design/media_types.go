package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var UserMedia = MediaType("application/user", func() {
	Description("A user")
	
	Attributes(func() {
		Attribute("id", String, "User unique identifier")
		Attribute("firstName", String, "First Name", func() {
			Example("John")
		})
		Attribute("lastName", String, "Last Name", func() {
			Example("Doe")
		})
		Attribute("age", Integer, "Age", func() {
			Minimum(18)
			Maximum(150)
			Example(25)
		})

		Required("id", "firstName", "lastName")
	})

	View("default", func() {
		Attribute("id")
		Attribute("firstName")
		Attribute("lastName")
		Attribute("age")
	})
})