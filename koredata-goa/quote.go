package main

import (
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
	// QuoteController_List: start_implement

	UserID := "danny"
	Quote := "Wassup peeps"

	quotes := app.JSON{
		UserID: &UserID,
		Quote:  &Quote,
	}

	// QuoteController_List: end_implement
	return ctx.OK(&quotes)
}

// ListByID runs the list by ID action.
//func (c *QuoteController) ListByID(ctx *app.ListByIDQuoteContext) error {
// QuoteController_ListByID: start_implement

//	quotes := map[string]string{
//		"Danny": "Wassup peeps",
//		"Jack":  "Hit the road",
//	}

//	quoteJson, err := json.Marshal(quotes)

//	if err != nil {
//		return fmt.Errorf("unable to marshal quotes to json: %s", err)
//	}

// QuoteController_ListByID: end_implement
//	return quoteJson
//}
