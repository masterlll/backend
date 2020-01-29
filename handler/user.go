package handler

import (
	"log"
	"net/http"

	"backend/lid/auth"
	"backend/lid/middleware"
	"backend/lid/util"
	"backend/model"

	"backend/database"

	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

func UserCreate(w http.ResponseWriter, r *http.Request, urlValues map[string]string, db *sqlx.DB) {

	user := struct {
		Name string `json:"name" validate:"required"`
	}{}

	err := util.ValidType(r, &user)
	if err != nil {
		middleware.SendResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	userID := uuid.NewV4().String()
	database.UserCreate(db, userID, user.Name)
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

// UserGet...
func UserGetOne(r *http.Request, urlValues map[string]string, db *sqlx.Tx, userid string) (statusCode int, err error, output interface{}) {
	user := model.User{}
	log.Println(urlValues["userId"])
	statusCode, err = database.UserGetOne(&user, db, urlValues["user"])
	return statusCode, err, user
}
