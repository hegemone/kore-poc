package main

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/thefirstofthe300/kore-poc/koredata-goa/app"
)

// QuoteController implements the quote resource.
type QuoteController struct {
	*goa.Controller
	db *gorm.DB
}

// NewQuoteController creates a quote controller.
func NewQuoteController(service *goa.Service, db *gorm.DB) *QuoteController {
	return &QuoteController{Controller: service.NewController("QuoteController"), db: db}
}

// List runs the list action.
func (c *QuoteController) List(ctx *app.ListQuoteContext) error {

	response := &app.JSON{}

	c.db.Find(&response.Quotes)

	return ctx.OK(response)
}

// ListByID runs the list by ID action.
func (c *QuoteController) ListByID(ctx *app.ListByIDQuoteContext) error {
	response := &app.JSON{}

	var quote app.Quote

	c.db.Where("name = ?", ctx.UserID).First(&quote)

	response.Quotes = append(response.Quotes, &quote)

	return ctx.OK(response)
}
