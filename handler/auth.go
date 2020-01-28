package handler

import (
	"fmt"
	"net/http"

	"backend/lid/middleware"

	"github.com/jmoiron/sqlx"
)

// UserGet...
func AuthGet(w http.ResponseWriter, req *http.Request, urlValues map[string]string, db *sqlx.DB) {
	fmt.Println(db)
	middleware.SendResponse(w, http.StatusOK, map[string]string{"AuthGet": urlValues["auth"]})
}
