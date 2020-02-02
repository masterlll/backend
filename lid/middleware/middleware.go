package middleware

import (
	"backend/lid/auth"
	"backend/lid/logroute"
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

type Handler func(r *http.Request, urlValues map[string]string, db *sqlx.Tx, userId string) (statusCode int, err error, output interface{})

//PlainHandler
type PlainHandler func(res http.ResponseWriter, req *http.Request, urlValues map[string]string, db *sqlx.DB)

//SendResponse  send a http response to the user with JSON format
func SendResponse(res http.ResponseWriter, statusCode int, data interface{}, err error) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.WriteHeader(statusCode)
	if d, ok := data.([]byte); ok {
		res.Write(d)
	} else {
		json.NewEncoder(res).Encode(data)
	}
	if err != nil {
		logroute.LogSave(err.Error())
	}
}

func Plain(f PlainHandler) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		f(res, req, mux.Vars(req), db)
	}

}

//Wrap  a middleware to handle user authorization
func Wrap(f Handler) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		userId, err := auth.Verify(req.Header.Get("Authorization"))
		if err != nil {
			SendResponse(res, http.StatusUnauthorized, map[string]string{"error": err.Error()}, err)
			return
		} else {
			//please think carefully on this design, as it has potential security problem
			if newToken, err := auth.ToekenSign(userId); err != nil {
				SendResponse(res, http.StatusInternalServerError, map[string]string{"error": err.Error()}, err)
				return
			} else {
				res.Header().Add("Authorization", newToken) // update JWT Token
			}
		}

		//prepare a database session for the handler
		session, err := db.Beginx()
		if err != nil {
			SendResponse(res, http.StatusInternalServerError, map[string]string{"error": err.Error()}, err)
			return
		}
		//everything seems fine, goto the business logic handler
		if statusCode, err, output := f(req, mux.Vars(req), session, userId); err == nil {
			//the business logic handler return no error, then try to commit the db session
			if err := session.Commit(); err != nil {
				SendResponse(res, http.StatusInternalServerError, map[string]string{"error": err.Error()}, err)
			} else {
				SendResponse(res, statusCode, output, nil)
			}
		} else {
			session.Rollback()
			SendResponse(res, statusCode, map[string]string{"error": err.Error()}, err)
		}
	}
}
