package handler

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

// UserGet...
func AuthGet(r *http.Request, urlValues map[string]string, db *sqlx.Tx, userid string) (statusCode int, err error, output interface{}) {
	// middleware.SendResponse(w, http.StatusOK, map[string]string{"AuthGet": userid})

	statusCode = 200
	err = nil
	output = "ok"
	return

}
