package handler

import (
	"errors"
	"net/http"

	"backend/lid/auth"
	"backend/lid/middleware"
	"backend/lid/util"
	"backend/model"

	"backend/database"

	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func UserCreate(w http.ResponseWriter, r *http.Request, urlValues map[string]string, db *sqlx.DB) {

	user := struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}{}

	err := util.ValidType(r, &user)
	if err != nil {
		middleware.SendResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()}, err)
		return
	}
	userID := uuid.NewV4().String()

	// hash  password
	digest, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		middleware.SendResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()}, err)
		return
	}
	user.Password = string(digest)

	_, err = database.UserCreate(db, userID, user.Email, user.Password)
	if err != nil {
		middleware.SendResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()}, err)
		return
	}

	if newToken, err := auth.ToekenSign(userID); err != nil {
		middleware.SendResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()}, err)
	} else {
		// update JWT Token
		w.Header().Add("Authorization", newToken)
		//allow CORS
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		middleware.SendResponse(w, http.StatusOK, map[string]string{"userId": userID}, nil)
	}
}

func UserPassWordUpdate(r *http.Request, urlValues map[string]string, db *sqlx.Tx, userId string) (int, error, interface{}) {
	id := urlValues[`userId`]
	if id != userId {
		return http.StatusForbidden, errors.New("Updating others account is forbidden"), nil
	}
	input := struct {
		Password         string `json:"password" validate:"required"`
		OriginalPassword string `json:"originalpassword" validate:"required"`
	}{}

	//perform the input binding
	err := util.ValidType(r, &input)
	//bind the input
	if err != nil {
		return http.StatusBadRequest, err, nil
	}

	user := model.User{}
	statecode, err := database.UserForIdGetOne(&user, db, userId)
	if err != nil {
		return statecode, err, nil
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.OriginalPassword)) != nil {
		return http.StatusForbidden, errors.New(`The original password is invalid`), nil
	}
	// hash  password
	newpassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {

		return http.StatusInternalServerError, err, nil
	}
	stateCode, err := database.UserPasswordUpdate(db, userId, string(newpassword))
	return stateCode, err, nil
}

// UserDelete...
func UserDelete(r *http.Request, urlValues map[string]string, db *sqlx.Tx, userid string) (statusCode int, err error, output interface{}) {
	user := model.User{}
	id := urlValues[`userId`]
	if id != userid {
		return http.StatusForbidden, errors.New("Updating others account is forbidden"), nil
	}
	statusCode, err = database.UserDelete(db, userid)
	return statusCode, err, user
}

func UserGetOne(r *http.Request, urlValues map[string]string, db *sqlx.Tx, userid string) (statusCode int, err error, output interface{}) {
	user := model.User{}

	statusCode, err = database.UserForIdGetOne(&user, db, urlValues["userId"])

	data := make(map[string]interface{})

	data["userId"] = user.ID
	data["nickname"] = user.Name

	return statusCode, err, data
}
