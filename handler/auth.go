package handler

import (
	"backend/database"
	"backend/lid/auth"
	"backend/lid/middleware"
	"backend/lid/util"
	"backend/model"
	"net/http"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

// UserGet...
func Login(w http.ResponseWriter, r *http.Request, urlValues map[string]string, db *sqlx.DB) {
	//handle the input
	var input struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	err := util.ValidType(r, &input)
	if err != nil {
		middleware.SendResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	user := model.User{}
	_, err = database.UserGetOne(&user, db, input.Email)
	if err != nil {
		middleware.SendResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
		middleware.SendResponse(w, http.StatusUnauthorized, map[string]string{"error": "Incorrect Email / Password"})
		return
	}

	if newToken, err := auth.ToekenSign(user.ID); err != nil {
		middleware.SendResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
	} else {
		// update JWT Token
		w.Header().Add("Authorization", newToken)
		//allow CORS
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		middleware.SendResponse(w, http.StatusOK, map[string]string{"userId": user.ID})
	}
}
