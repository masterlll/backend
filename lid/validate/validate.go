package validate

import "gopkg.in/go-playground/validator.v9"

var Validater *validator.Validate

func init() {
	Validater = validator.New()
}
