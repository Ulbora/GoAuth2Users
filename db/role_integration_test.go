// +build integration move to top

package db

import (
	"fmt"
	"testing"

	dbi "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

var dbbi dbi.Database
var mydb mdb.MyDB
var rid int64

func TestUserDBRolei_Con(t *testing.T) {
	// var mydb mdb.MyDB
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "go_auth2_users"
	dbbi = &mydb
	dbbi.Connect()

}

func TestUserDBRolei_AddRole(t *testing.T) {
	// var dbbi dbi.Database
	// // var mydb mdb.MyDB
	// mydb.Host = "localhost:3306"
	// mydb.User = "admin"
	// mydb.Password = "admin"
	// mydb.Database = "ulbora_oauth2_server"
	// dbbi = &mydb
	var db UserDatabase
	var udb UserDB
	udb.DB = dbbi
	db = &udb

	var r Role
	r.Role = "testRole"
	suc, id := db.AddRole(&r)
	if !suc && id <= 0 {
		t.Fail()
	} else {
		rid = id
	}
}

func TestUserDBRolei_GetRole(t *testing.T) {
	// var dbbi dbi.Database
	// // var mydb mdb.MyDB
	// mydb.Host = "localhost:3306"
	// mydb.User = "admin"
	// mydb.Password = "admin"
	// mydb.Database = "ulbora_oauth2_server"
	// dbbi = &mydb
	var db UserDatabase
	var udb UserDB
	udb.DB = dbbi
	db = &udb

	role := db.GetRole(rid)
	if role.Role != "testRole" {
		t.Fail()
	}
}

func TestUserDBRolei_GetRoleList(t *testing.T) {
	// var dbbi dbi.Database
	// // var mydb mdb.MyDB
	// mydb.Host = "localhost:3306"
	// mydb.User = "admin"
	// mydb.Password = "admin"
	// mydb.Database = "ulbora_oauth2_server"
	// dbbi = &mydb
	var db UserDatabase
	var udb UserDB
	udb.DB = dbbi
	db = &udb

	roles := db.GetRoleList()
	fmt.Println("roles: ", roles)
	if len(*roles) == 0 {
		t.Fail()
	}
}

func TestUserDBRolei_DeleteRole(t *testing.T) {
	// var dbbi dbi.Database
	// // var mydb mdb.MyDB
	// mydb.Host = "localhost:3306"
	// mydb.User = "admin"
	// mydb.Password = "admin"
	// mydb.Database = "ulbora_oauth2_server"
	// dbbi = &mydb
	var db UserDatabase
	var udb UserDB
	udb.DB = dbbi
	db = &udb

	var r Role
	r.Role = "testRole"
	suc := db.DeleteRole(rid)
	if !suc {
		t.Fail()
	}
}
