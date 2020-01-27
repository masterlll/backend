package main

import (
	"log"
	"runtime"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"gopkg.in/tylerb/graceful.v1"
)

func main() {

	routeHub()

}

// init
func init() {

	log.Println("init ")

	//Old go compiler, it is a must to enable multithread processing
	runtime.GOMAXPROCS(runtime.NumCPU())
	// db
	log.Println("db ")

	// auth jwt

	//// middleware
}

// RouteHub
func routeHub() {

	// urlRandomRegexp

	router := mux.NewRouter()

	//  Routes
	//	authRoute(router)
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

}

//  Route : User
func userRoute(router *mux.Router) {

	router.HandleFunc("/v1/users", middleware.Plain(handler.UserCreate)).Methods("GET")

}
