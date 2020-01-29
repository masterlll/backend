package handler

import (
	"net/http"

	"backend/lid/auth"
	"backend/lid/middleware"

	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

// UserGet...
func UserGet(r *http.Request, urlValues map[string]string, db *sqlx.Tx, userid string) (statusCode int, err error, output interface{}) {

	// middleware.SendResponse(w, http.StatusOK, map[string]string{"userId": userid})

	statusCode = 200
	err = nil
	output = "ok"
	return
}

func UserCreate(w http.ResponseWriter, r *http.Request, urlValues map[string]string, db *sqlx.DB) {

	userID := uuid.NewV4().String()
	if newToken, err := auth.ToekenSign(userID); err != nil {
		middleware.SendResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
	} else {
		// update JWT Token
		w.Header().Add("Authorization", newToken)
		//allow CORS
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		middleware.SendResponse(w, http.StatusOK, map[string]string{"userId": userID})
	}
}
