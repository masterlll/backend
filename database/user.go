package database

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
)

func UserCreate(db *sqlx.DB, userid, email, password string) (statusCode int, err error) {
	sqlString := `	insert into users(id, email, password)
	select ?, ? ,?
	where not exists (select 1 from users  where id = ?);`
	_, err = db.Exec(sqlString, userid, email, password, userid)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, err
}

func UserGetOne(user interface{}, db *sqlx.DB, email string) (statusCode int, err error) {
	sqlString := `SELECT 
     id ,nickname  , password
     FROM
      users
     WHERE
      email = ?;`
	err = db.Get(user, sqlString, email)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, err
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, err
}

func UserForIdGetOne(user interface{}, db *sqlx.Tx, userid string) (statusCode int, err error) {
	sqlString := `SELECT 
     id ,nickname  , password
     FROM
      users
     WHERE
	  id  = ?;`
	err = db.Get(user, sqlString, userid)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, err
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, err
}

func UserPasswordUpdate(db *sqlx.Tx, userid, password string) (statusCode int, err error) {

	t := time.Now().UTC()
	sqlString := `UPDATE users 
	SET 
	password = ? , 
	updated_at = ?
	WHERE
		id = ?`
	_, err = db.Exec(sqlString, password, t, userid)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, err
}

func UserDelete(db *sqlx.Tx, userid string) (statusCode int, err error) {

	sqlString := `DELETE FROM users 
	WHERE
		id = ?`
	_, err = db.Exec(sqlString, userid)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, err
}
