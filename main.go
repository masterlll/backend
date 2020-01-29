package main

import (
	"backend/handler"
	"backend/lid/auth"
	"backend/lid/config"
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

	//Old go compiler, it is a must to enable multithread processing
	//runtime.GOMAXPROCS(runtime.NumCPU())
	// db
	log.Println(" Db Start ... ")
	connString :=
		config.GetStr(setting.DBUserName) +
			":" + config.GetStr(setting.DBPassword) +
			"@tcp(" + config.GetStr(setting.DBHost) +
			":" + strconv.Itoa(config.GetInt(setting.DBPort)) +
			")/" + config.GetStr(setting.DBName) +
			"?allowNativePasswords=true&parseTime=true&charset=utf8mb4"

		//	 log.Println("connString", connString)

	//"root:8gm3Ncxf@tcp(47.91.251.106:3307)/group_ut?allowNativePasswords=true&parseTime=true&charset=utf8mb4
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
}

// RouteHub
func routeHub() {

	// urlRandomRegexp

	router := mux.NewRouter()

	//  Routes
	authRoute(router)
	userRoute(router)
	//

	recovery := negroni.NewRecovery()
	recovery.PrintStack = false
	n := negroni.New(recovery, negroni.NewLogger())
	n.UseHandler(router)
	graceful.Run(":8080", 1*time.Minute, n)
}

//  Route : Auth
func authRoute(router *mux.Router) {

	router.HandleFunc("/v1/auths/{auth}", middleware.Wrap(handler.AuthGet)).Methods("GET")

}

//  Route : User
func userRoute(router *mux.Router) {

	router.HandleFunc("/v1/users/{user}", middleware.Wrap(handler.UserGet)).Methods("GET")
	router.HandleFunc("/v1/users/", middleware.Plain(handler.UserCreate)).Methods("POST")
}
