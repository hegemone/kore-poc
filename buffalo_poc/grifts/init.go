package grifts

import (
	"github.com/dahendel/kore-poc/buffalo_poc/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
