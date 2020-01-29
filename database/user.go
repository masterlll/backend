package database

import (
	"database/sql"
	"net/http"

	"github.com/jmoiron/sqlx"
)

func UserGetOne(user interface{}, db *sqlx.Tx, userid string) (statusCode int, err error) {
	sqlString := `SELECT 
    user_id ,name  ,create_at ,update_at
FROM
    user_info
WHERE
    user_id = ?;`
	err = db.Get(user, sqlString, userid)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, err
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, err
}

func UserCreate(db *sqlx.DB, userid, name string) (statusCode int, err error) {
	sqlString := `	insert into user_info(user_id, name)
	select ?, ? 
	where not exists (select 1 from user_info where user_id = ?);`
	_, err = db.Exec(sqlString, userid, name, userid)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, err
}
