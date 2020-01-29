package handler

import (
	"backend/lid/auth"
	"backend/lid/middleware"
	"backend/lid/util"
	"net/http"

	"github.com/jmoiron/sqlx"
)

// UserGet...
func Login(w http.ResponseWriter, r *http.Request, urlValues map[string]string, db *sqlx.DB) {
	//handle the input
	var input struct {
		Userid string `json:"userid" validate:"required"`
	}
	err := util.ValidType(r, &input)
	if err != nil {
		middleware.SendResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	//UserGetOne

	if newToken, err := auth.ToekenSign(input.Userid); err != nil {
		middleware.SendResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
	} else {
		// update JWT Token
		w.Header().Add("Authorization", newToken)
		//allow CORS
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		middleware.SendResponse(w, http.StatusOK, map[string]string{"userId": input.Userid})
	}
}
