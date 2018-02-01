package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/pop/slices"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/satori/go.uuid"
)

type Quote struct {
	ID        uuid.UUID     `json:"id" db:"id"`
	CreatedAt time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" db:"updated_at"`
	Name      string        `json:"name" db:"name"`
	Quote     string        `json:"quote" db:"quote"`
	Quotes    slices.String `json:"quotes" db:"quotes"`
}

// String is not required by pop and may be deleted
func (q Quote) String() string {
	jq, _ := json.Marshal(q)
	return string(jq)
}

// Quotes is not required by pop and may be deleted
type Quotes []Quote

// String is not required by pop and may be deleted
func (q Quotes) String() string {
	jq, _ := json.Marshal(q)
	return string(jq)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (q *Quote) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: q.Name, Name: "Name"},
		&validators.StringIsPresent{Field: q.Quote, Name: "Quote"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (q *Quote) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (q *Quote) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
