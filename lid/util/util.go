package util

import (
	"backend/lid/validate"
	"encoding/json"
	"net/http"
)

// ValidType :
func ValidType(r *http.Request, param interface{}) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(param); err != nil {
		return err
	}
	err := validate.Validater.Struct(param)
	return err
}
