//go:generate goagen bootstrap -d github.com/thefirstofthe300/kore-poc/koredata-goa/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/thefirstofthe300/kore-poc/koredata-goa/app"
)

func main() {
	// Create service
	service := goa.New("koredata")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "quote" controller
	c := NewQuoteController(service)
	app.MountQuoteController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
