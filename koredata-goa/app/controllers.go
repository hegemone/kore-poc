// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "koredata": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/thefirstofthe300/kore-poc/koredata-goa/design
// --out=$(GOPATH)/src/github.com/thefirstofthe300/kore-poc/koredata-goa
// --version=v1.3.0

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// QuoteController is the controller interface for the Quote actions.
type QuoteController interface {
	goa.Muxer
	Create(*CreateQuoteContext) error
	List(*ListQuoteContext) error
	ListByID(*ListByIDQuoteContext) error
	Login(*LoginQuoteContext) error
}

// MountQuoteController "mounts" a Quote resource controller on the given service.
func MountQuoteController(service *goa.Service, ctrl QuoteController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateQuoteContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateQuotePayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleSecurity("jwt", h, "api:write")
	service.Mux.Handle("POST", "/quotes", ctrl.MuxHandler("create", h, unmarshalCreateQuotePayload))
	service.LogInfo("mount", "ctrl", "Quote", "action", "Create", "route", "POST /quotes", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListQuoteContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/quotes", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Quote", "action", "List", "route", "GET /quotes")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListByIDQuoteContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.ListByID(rctx)
	}
	h = handleSecurity("jwt", h, "api:read")
	service.Mux.Handle("GET", "/quotes/:userId", ctrl.MuxHandler("list by ID", h, nil))
	service.LogInfo("mount", "ctrl", "Quote", "action", "ListByID", "route", "GET /quotes/:userId", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewLoginQuoteContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Login(rctx)
	}
	h = handleSecurity("BasicAuth", h)
	service.Mux.Handle("POST", "/quotes/login", ctrl.MuxHandler("login", h, nil))
	service.LogInfo("mount", "ctrl", "Quote", "action", "Login", "route", "POST /quotes/login", "security", "BasicAuth")
}

// unmarshalCreateQuotePayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateQuotePayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createQuotePayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}
