package main

import (
	"github.com/goadesign/goa"
	"github.com/hegemone/kore-poc/koredata-goa/app"
	"github.com/jinzhu/gorm"
)

// SuggestionController implements the suggestion resource.
type SuggestionController struct {
	*goa.Controller
	db *gorm.DB
}

// NewSuggestionController creates a suggestion controller.
func NewSuggestionController(service *goa.Service, db *gorm.DB) *SuggestionController {
	return &SuggestionController{Controller: service.NewController("SuggestionController")}
}

// Create runs the create action.
func (c *SuggestionController) Create(ctx *app.CreateSuggestionContext) error {
	// SuggestionController_Create: start_implement

	// Put your logic here

	suggestion := &app.Suggestion{
		ShowID:    &ctx.Payload.ShowID,
		Suggester: &ctx.Payload.Suggester,
		Title:     &ctx.Payload.Title,
	}

	c.db.Create(suggestion)

	return nil
	// SuggestionController_Create: end_implement
}

// List runs the list action.
func (c *SuggestionController) List(ctx *app.ListSuggestionContext) error {
	// SuggestionController_List: start_implement

	// Put your logic here

	res := &app.Suggestions{}

	c.db.Where("name = ?", ctx.ShowID).Find(res)

	return ctx.OK(res)
	// SuggestionController_List: end_implement
}
