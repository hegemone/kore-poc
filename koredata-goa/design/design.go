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

var BasicAuth = BasicAuthSecurity("BasicAuth", func() {
	Description("Use client ID and client secret to authenticate")
})

var JWT = JWTSecurity("jwt", func() {
	Header("Authorization")
	Scope("api:write", "API write access")
	Scope("api:read", "API read access")
})

var _ = Resource("quote", func() {
	BasePath("/quotes")
	Security(JWT, func() {
		Scope("api:read")
	})
	DefaultMedia(Quotes)

	Action("login", func() {
		Description("Login to the api")
		Routing(POST("/login"))
		Security(BasicAuth)
		Response(NoContent, func() {
			Headers(func() {
				Header("Authorization", String, "Generated JWT")
			})
		})
		Response(Unauthorized)
	})

	Action("list", func() {
		Description("Returns all quotes in the quote database")
		Routing(GET(""))
		NoSecurity()
		Response(OK)
		Response(Unauthorized)
	})

	Action("create", func() {
		Description("Create a quote and add it to the database")
		Security(JWT, func() {
			Scope("api:write")
		})
		Routing(POST(""))
		Payload(Quote, func() {
			Required("Name", "Quote")
		})
		Response(OK)
		Response(Unauthorized)
	})

	Action("list by ID", func() {
		Description("Returns all the quotes for a given person")
		Routing(GET("/:userId"))
		NoSecurity()
		Params(func() {
			Param("userId", String, "User ID")
		})
		Response(OK, Quote)
		Response(NotFound)
	})
})

var _ = Resource("suggestion", func() {
	BasePath("/suggestion")
	NoSecurity()
	DefaultMedia(Suggestions)

	Action("create", func() {
		Description("Create a new suggestion")
		Routing(POST("/"))
		Payload(Suggestion, func() {
			Required("ShowID", "Suggester", "Title")
		})
		Response(NoContent)
	})

	Action("list", func() {
		Description("Return all suggestions for a given show ID")
		Routing(GET("/:showId"))
		Response(OK)
		Response(NotFound)
	})
})

var Quotes = MediaType("vnd.application.io/quotes", func() {
	Description("A quote from the user database")
	Attributes(func() {
		Attribute("Quotes", ArrayOf(Quote), "Quote")
	})
	View("default", func() {
		Attribute("Quotes")
	})
})

var Quote = MediaType("vnd.application.io/quote", func() {
	Description("All quotes for a given user ID")
	Attributes(func() {
		Attribute("ID", Integer, "ID of the user")
		Attribute("Name", String, "User ID of quoter")
		Attribute("Quote", String, "The actual quotes of the quoter")
	})
	View("default", func() {
		Attribute("ID", Integer, "ID of the user")
		Attribute("Name", String, "User ID of quoter")
		Attribute("Quote", String, "The actual quotes of the quoter")
	})
})

var Suggestions = MediaType("vnd.application.io/suggestions", func() {
	Attributes(func() {
		Attribute("Suggestions", ArrayOf(Suggestion), "Suggestion")
	})
	View("default", func() {
		Attribute("Suggestions")
	})
})

var Suggestion = Type("suggestion", func() {
	Description("All suggestions for a given user ID")
	Attribute("ShowID", String, "The ID of the show")
	Attribute("Suggester", String, "Identity of suggester")
	Attribute("Title", String, "The suggested title")
})
