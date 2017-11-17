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
	DefaultMedia(quotes)

	Action("list", func() {
		Description("Returns all quotes in the quote database")
		Routing(GET(""))
		Response(OK)
	})

	Action("list by ID", func() {
		Description("Returns all the quotes for a given person")
		Routing(GET("/:userId"))
		Params(func() {
			Param("userId", String, "User ID")
		})
		Response(OK)
		Response(NotFound)
	})
})

var quotes = MediaType("application/json", func() {
	Description("A quote from the user database")
	Attributes(func() {
		Attribute("quotes", ArrayOf(userQuotes), "quote")
	})
	View("default", func() {
		Attribute("quotes")
	})
})

var userQuotes = Type("quote", func() {
	Description("All quotes for a given user ID")
	Attribute("userID", String, "User ID of quoter")
	Attribute("quote", ArrayOf(String), "The actual quotes of the quoter")
})
