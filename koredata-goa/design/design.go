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
	DefaultMedia(quotes)

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
		Payload(QuotePayload, func() {
			Required("Name", "Quote")
		})
		Response(OK)
		Response(Unauthorized)
	})

	Action("list by ID", func() {
		Description("Returns all the quotes for a given person")
		Routing(GET("/:userId"))
<<<<<<< HEAD
		NoSecurity()
=======
>>>>>>> upstream/master
		Params(func() {
			Param("userId", String, "User ID")
		})
		Response(OK)
<<<<<<< HEAD
=======
		Response(Unauthorized)
>>>>>>> upstream/master
		Response(NotFound)
	})
})

var BasicAuth = BasicAuthSecurity("BasicAuth", func() {
	Description("Use client ID and client secret to authenticate")
})

var quotes = MediaType("application/json", func() {
	Description("A quote from the user database")
	Attributes(func() {
		Attribute("Quotes", ArrayOf(userQuotes), "Quote")
	})
	View("default", func() {
		Attribute("Quotes")
	})
})

var QuotePayload = Type("QuotePayload", func() {
	Attribute("Name", func() {
		MinLength(2)
		Example("Number 8")
	})
	Attribute("Quote", func() {
		MinLength(2)
		Example("Asti")
	})
})

var userQuotes = Type("Quote", func() {
	Description("All quotes for a given user ID")
	Attribute("ID", Integer, "ID of the user")
	Attribute("Name", String, "User ID of quoter")
	Attribute("Quote", String, "The actual quotes of the quoter")
})
