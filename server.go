package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	jv "github.com/Ulbora/GoAuth2JwtValidator"
	udb "github.com/Ulbora/GoAuth2Users/db"
	han "github.com/Ulbora/GoAuth2Users/handlers"
	m "github.com/Ulbora/GoAuth2Users/managers"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
	"github.com/gorilla/mux"
)

//GO111MODULE=on go mod init github.com/Ulbora/GoAuth2Users
func main() {
	var dbi db.Database
	var mydb mdb.MyDB
	var goAuth2UsersHost string
	var goAuth2UsersDbUser string
	var goAuth2UsersDbPassword string
	var goAuth2UsersDatabase string

	if os.Getenv("GO_AUTH2_USERS_HOST") != "" {
		goAuth2UsersHost = os.Getenv("GO_AUTH2_USERS_HOST")
	} else {
		goAuth2UsersHost = "localhost:3306"
	}

	if os.Getenv("GO_AUTH2_USERS_DB_USER") != "" {
		goAuth2UsersDbUser = os.Getenv("GO_AUTH2_USERS_DB_USER")
	} else {
		goAuth2UsersDbUser = "admin"
	}

	if os.Getenv("GO_AUTH2_USERS_DB_PASSWORD") != "" {
		goAuth2UsersDbPassword = os.Getenv("GO_AUTH2_USERS_DB_PASSWORD")
	} else {
		goAuth2UsersDbPassword = "admin"
	}

	if os.Getenv("GO_AUTH2_USERS_DATABASE") != "" {
		goAuth2UsersDatabase = os.Getenv("GO_AUTH2_USERS_DATABASE")
	} else {
		goAuth2UsersDatabase = "go_auth2_users"
	}

	mydb.Host = goAuth2UsersHost
	mydb.User = goAuth2UsersDbUser
	mydb.Password = goAuth2UsersDbPassword
	mydb.Database = goAuth2UsersDatabase
	dbi = &mydb
	dbi.Connect()
	var userDB udb.UserDB
	userDB.DB = dbi
	var uman m.UserManager
	uman.UserDB = &userDB
	var uh han.UserHandler
	uh.Manager = &uman
	var mc jv.OauthClient
	//var proxy gp.GoProxy
	//mc.Proxy = &proxy
	//var compress cp.JwtCompress
	//mc.JwtCompress = compress
	uh.ValidatorClient = mc.GetNewClient()
	h := uh.GetNew()

	router := mux.NewRouter()
	port := "3001"
	envPort := os.Getenv("PORT")
	if envPort != "" {
		portInt, _ := strconv.Atoi(envPort)
		if portInt != 0 {
			port = envPort
		}
	}

	router.HandleFunc("/rs/role/add", h.AddRole).Methods("POST")
	router.HandleFunc("/rs/role/get/{id}", h.GetRole).Methods("GET")
	router.HandleFunc("/rs/role/list", h.GetRoleList).Methods("GET")
	router.HandleFunc("/rs/role/delete/{id}", h.DeleteRole).Methods("DELETE")

	router.HandleFunc("/rs/user/add", h.AddUser).Methods("POST")
	router.HandleFunc("/rs/user/update", h.UpdateUser).Methods("PUT")
	router.HandleFunc("/rs/user/get/{username}/{clientId}", h.GetUser).Methods("GET")
	router.HandleFunc("/rs/user/list", h.GetUserList).Methods("GET")
	router.HandleFunc("/rs/user/search/{clientId}", h.SearchUserList).Methods("GET")
	router.HandleFunc("/rs/user/delete/{username}/{clientId}", h.DeleteUser).Methods("DELETE")

	router.HandleFunc("/rs/client/user/add", h.ClientAddUser).Methods("POST")
	router.HandleFunc("/rs/client/user/update", h.ClientUpdateUser).Methods("PUT")
	router.HandleFunc("/rs/client/user/get/{username}", h.ClientGetUser).Methods("GET")
	router.HandleFunc("/rs/client/user/search", h.ClientSearchUserList).Methods("GET")
	router.HandleFunc("/rs/client/user/delete/{username}", h.ClientDeleteUser).Methods("DELETE")

	router.HandleFunc("/rs/user/login", h.LoginUser).Methods("POST")

	fmt.Println("Starting server Oauth2 Server on " + port)
	http.ListenAndServe(":"+port, router)

}
