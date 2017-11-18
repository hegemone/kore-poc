// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "koredata": Application User Types
//
// Command:
// $ goagen
// --design=github.com/thefirstofthe300/kore-poc/koredata-goa/design
// --out=$(GOPATH)/src/github.com/thefirstofthe300/kore-poc/koredata-goa
// --version=v1.3.0

package app

// All quotes for a given user ID
type quote struct {
	// ID of the user
	ID *int `form:"ID,omitempty" json:"ID,omitempty" xml:"ID,omitempty"`
	// User ID of quoter
	Name *string `form:"Name,omitempty" json:"Name,omitempty" xml:"Name,omitempty"`
	// The actual quotes of the quoter
	Quote *string `form:"Quote,omitempty" json:"Quote,omitempty" xml:"Quote,omitempty"`
}

// Publicize creates Quote from quote
func (ut *quote) Publicize() *Quote {
	var pub Quote
	if ut.ID != nil {
		pub.ID = ut.ID
	}
	if ut.Name != nil {
		pub.Name = ut.Name
	}
	if ut.Quote != nil {
		pub.Quote = ut.Quote
	}
	return &pub
}

// All quotes for a given user ID
type Quote struct {
	// ID of the user
	ID *int `form:"ID,omitempty" json:"ID,omitempty" xml:"ID,omitempty"`
	// User ID of quoter
	Name *string `form:"Name,omitempty" json:"Name,omitempty" xml:"Name,omitempty"`
	// The actual quotes of the quoter
	Quote *string `form:"Quote,omitempty" json:"Quote,omitempty" xml:"Quote,omitempty"`
}