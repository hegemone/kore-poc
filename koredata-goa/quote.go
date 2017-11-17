package main

import (
	"fmt"

	"github.com/goadesign/goa"
	"github.com/thefirstofthe300/kore-poc/koredata-goa/app"
)

// QuoteController implements the quote resource.
type QuoteController struct {
	*goa.Controller
}

// NewQuoteController creates a quote controller.
func NewQuoteController(service *goa.Service) *QuoteController {
	return &QuoteController{Controller: service.NewController("QuoteController")}
}

// List runs the list action.
func (c *QuoteController) List(ctx *app.ListQuoteContext) error {

	response := &app.JSON{}

	quotes := make(map[string][]string)

	quotes["Danny"] = append(quotes["Danny"], "Wassup peeps")
	quotes["Danny"] = append(quotes["Danny"], "Hello world")
	quotes["Jack"] = append(quotes["Jack"], "Hit the road")

	for k, v := range quotes {
		// Since the quote object requires a pointer, we need to copy the data to
		// a new variable to avoid all user IDs in the response from being the same.
		// Since this is a proof of concept, this problem shouldn't be a huge issue
		// long term.
		userID := k
		quote := &app.Quote{UserID: &userID, Quote: v}
		response.Quotes = append(response.Quotes, quote)
	}

	return ctx.OK(response)
}

// ListByID runs the list by ID action.
func (c *QuoteController) ListByID(ctx *app.ListByIDQuoteContext) error {
	response := &app.JSON{}

	quotes := make(map[string][]string)

	quotes["Danny"] = append(quotes["Danny"], "Wassup peeps")
	quotes["Danny"] = append(quotes["Danny"], "Hello world")
	quotes["Jack"] = append(quotes["Jack"], "Hit the road")

	fmt.Println(ctx.UserID, quotes[ctx.UserID])

	response.Quotes = append(response.Quotes, &app.Quote{UserID: &ctx.UserID, Quote: quotes[ctx.UserID]})

	return ctx.OK(response)
}
