package middleware

import (
	"encoding/json"
	"net/http"
)

//Handler api
type Handler func(r *http.Request, urlValues map[string]string, userId string) (statusCode int, err error, output interface{})

//SendResponse  send a http response to the user with JSON format
func SendResponse(res http.ResponseWriter, statusCode int, data interface{}) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.WriteHeader(statusCode)
	if d, ok := data.([]byte); ok {
		res.Write(d)
	} else {
		json.NewEncoder(res).Encode(data)
	}
}

//Wrap  a middleware to handle user authorization
func Wrap(f Handler) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

	}

}
