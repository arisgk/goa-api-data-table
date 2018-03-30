package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var User = MediaType("application/data.table.user", func() {
	Description("A user")
	
	Attributes(func() {
		Attribute("id", String, "User unique identifier", func() {
			Example("123e4567-e89b-12d3-a456-426655440000")
		})
		Attribute("firstName", String, "First Name"), func() {
			Example("John")
		}
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
})