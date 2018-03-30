package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("Data Table", func() {
	Title("Data Table API")
	Description("Simple CRUD API implemented with goa")
	Host("localhost:5000")
	Scheme("http")
	BasePath("/api")
})