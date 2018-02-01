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

type Droplet struct {
	ID        uuid.UUID     `json:"id" db:"id"`
	CreatedAt time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" db:"updated_at"`
	Droplet   string        `json:"droplet" db:"droplet"`
	Droplets  slices.String `json:"droplets" db:"droplets"`
	DoToken   string        `json:"doToken" db:"doToken"`
}

// String is not required by pop and may be deleted
func (d Droplet) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// Droplets is not required by pop and may be deleted
type Droplets []Droplet

// String is not required by pop and may be deleted
func (d Droplets) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (d *Droplet) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: d.Droplet, Name: "Droplet"},
		&validators.StringIsPresent{Field: d.DoToken, Name: "DoToken"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (d *Droplet) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (d *Droplet) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
