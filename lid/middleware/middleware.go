package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

func Init(database *sqlx.DB) {
	db = database
}

//PlainHandler
type PlainHandler func(res http.ResponseWriter, req *http.Request, urlValues map[string]string, db *sqlx.DB)

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
func Plain(f PlainHandler) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		f(res, req, mux.Vars(req), db)
	}

}
