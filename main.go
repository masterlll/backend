package main

import (
	"backend/handler"
	"backend/lid/auth"
	"backend/lid/config"
	"backend/lid/logroute"
	"backend/lid/middleware"
	"backend/setting"
	"crypto/rsa"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/dgrijalva/jwt-go"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"gopkg.in/tylerb/graceful.v1"
)

func main() {
	routeHub()
}

// init
func init() {

	log.Println(" Init Start ...")
	// db
	log.Println(" Db Start ... ")
	connString :=
		config.GetStr(setting.DBUserName) +
			":" + config.GetStr(setting.DBPassword) +
			"@tcp(" + config.GetStr(setting.DBHost) +
			":" + strconv.Itoa(config.GetInt(setting.DBPort)) +
			")/" + config.GetStr(setting.DBName) +
			"?allowNativePasswords=true&parseTime=true&charset=utf8mb4"

	db, err := sqlx.Connect("mysql", connString)
	if err != nil {
		log.Panic("DB connection initialization failed", err)
	}

	db.SetMaxIdleConns(config.GetInt(setting.DBMaxIdleConn))
	db.SetMaxOpenConns(config.GetInt(setting.DBMaxOpenConn))
	// auth
	log.Println(" Auth start ... ")
	var err1 error
	var currentKey *rsa.PrivateKey = nil
	var oldKey *rsa.PrivateKey = nil

	currentKeyBytes, _ := ioutil.ReadFile(config.GetStr(setting.JwtRsaKeyLocation))
	currentKey, err1 = jwt.ParseRSAPrivateKeyFromPEM(currentKeyBytes)
	if err1 != nil {
		log.Panic(err1)
	}
	if location := config.GetStr(setting.JwtOldRsaKeyLocation); location != `''` {
		oldKeyBytes, _ := ioutil.ReadFile(location)
		oldKey, err1 = jwt.ParseRSAPrivateKeyFromPEM(oldKeyBytes)
		if err1 != nil {
			log.Panic(err1)
		}
	}

	lifetime := time.Duration(config.GetInt(setting.JwtToekenLifeTime)) * time.Minute
	auth.Init(currentKey, oldKey, lifetime)

	// middleware
	log.Println(" Middleware Start ... ")
	middleware.Init(db)

	log.Println(" Logroute  Start ... ")
	logroute.Init()

}

// RouteHub
func routeHub() {

	router := mux.NewRouter()
	uuidRegexp := `[[:alnum:]]{8}-[[:alnum:]]{4}-4[[:alnum:]]{3}-[89AaBb][[:alnum:]]{3}-[[:alnum:]]{12}`
	//  Routes
	authRoute(router, uuidRegexp)
	userRoute(router, uuidRegexp)
	//
	recovery := negroni.NewRecovery()
	recovery.PrintStack = false
	n := negroni.New(recovery, negroni.NewLogger())
	n.UseHandler(router)
	graceful.Run(":8080", 1*time.Minute, n)
}

//  Route : Auth
func authRoute(router *mux.Router, uuidRegexp string) {
	router.HandleFunc("/v1/auth/", middleware.Plain(handler.Login)).Methods("POST")
}

//  Route : User
func userRoute(router *mux.Router, uuidRegexp string) {

	router.HandleFunc("/v1/users/{userId:"+uuidRegexp+"}", middleware.Wrap(handler.UserGetOne)).Methods("GET")
	router.HandleFunc("/v1/users/", middleware.Plain(handler.UserCreate)).Methods("POST")
	router.HandleFunc("/v1/users/{userId:"+uuidRegexp+"}", middleware.Wrap(handler.UserPassWordUpdate)).Methods("PUT")
	router.HandleFunc("/v1/users/{userId:"+uuidRegexp+"}", middleware.Wrap(handler.UserDelete)).Methods("DELETE")
}
