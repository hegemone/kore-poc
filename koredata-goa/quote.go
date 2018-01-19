package main

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/thefirstofthe300/kore-poc/koredata-goa/app"
)

// QuoteController implements the quote resource.
type QuoteController struct {
	*goa.Controller
	privateKey *rsa.PrivateKey
	db         *gorm.DB
}

// NewQuoteController creates a quote controller.
func NewQuoteController(service *goa.Service, db *gorm.DB) (*QuoteController, error) {
	b, err := ioutil.ReadFile("./jwtkey/jwt.key")
	if err != nil {
		return nil, err
	}
	privKey, err := jwtgo.ParseRSAPrivateKeyFromPEM(b)
	if err != nil {
		return nil, fmt.Errorf("jwt: failed to load private key: %s", err) // bug
	}
	return &QuoteController{
		Controller: service.NewController("QuoteController"),
		privateKey: privKey,
		db:         db,
	}, nil
}

// CreateQuote runs the create action
func (c *QuoteController) Login(ctx *app.LoginQuoteContext) error {
	// Generate JWT
	token := jwtgo.New(jwtgo.SigningMethodRS512)
	in10m := time.Now().Add(time.Duration(10) * time.Minute).Unix()
	token.Claims = jwtgo.MapClaims{
		"iss":    "Issuer",              // who creates the token and signs it
		"aud":    "Audience",            // to whom the token is intended to be sent
		"exp":    in10m,                 // time when the token will expire (10 minutes from now)
		"jti":    uuid.NewV4().String(), // a unique identifier for the token
		"iat":    time.Now().Unix(),     // when the token was issued/created (now)
		"nbf":    2,                     // time before which the token is not yet valid (2 minutes ago)
		"sub":    "subject",             // the subject/principal is whom the token is about
		"scopes": "api:read",            // token scope - not a standard claim
	}
	signedToken, err := token.SignedString(c.privateKey)
	if err != nil {
		return fmt.Errorf("failed to sign token: %s", err) // internal error
	}

	// Set auth header for client retrieval
	ctx.ResponseData.Header().Set("Authorization", "Bearer "+signedToken)

	// Send response
	return ctx.NoContent()
}

// List runs the list action.
func (c *QuoteController) List(ctx *app.ListQuoteContext) error {

	response := &app.JSON{}

	c.db.Find(&response.Quotes)

	return ctx.OK(response)
}

// CreateQuote runs the create action
func (c *QuoteController) Create(ctx *app.CreateQuoteContext) error {
	response := &app.JSON{}

	quote := &app.Quote{
		Name:  &ctx.Payload.Name,
		Quote: &ctx.Payload.Quote,
	}

	c.db.Create(quote)

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
