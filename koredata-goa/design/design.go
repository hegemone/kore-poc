package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("koredata", func() {
	Title("The Kore of the Data")
	Description("Allows users to interact with Jupiter Broadcasting's shows")
	Scheme("http")
})

var _ = Resource("quote", func() {
	BasePath("/quotes")
	DefaultMedia(Quote)

	Action("list", func() {
		Description("Returns all quotes in the quote database")
		Routing(GET("/"))
		Response(OK)
	})

	//	Action("list by ID", func() {
	//		Description("Returns all the quotes for a given person")
	//		Routing(GET("/:personID"))
	//		Params(func() {
	//			Param("personID", String, "Person ID")
	//		})
	//		Response(OK)
	//		Response(NotFound)
	//	})
})

var Quote = MediaType("application/json", func() {
	Description("A quote from the user database")
	Attributes(func() {
		Attribute("userID", String, "unique user ID")
		Attribute("quote", String, "quote")
	})
	View("default", func() {
		Attribute("userID")
		Attribute("quote")
	})
})
