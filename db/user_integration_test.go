// +build integration move to top

package db

import (
	"fmt"
	"testing"
	"time"

	dbi "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

var dbbiu dbi.Database
var mydbu mdb.MyDB
var uid int64

func TestUserDBi_Con(t *testing.T) {
	// var mydb mdb.MyDB
	mydbu.Host = "localhost:3306"
	mydbu.User = "admin"
	mydbu.Password = "admin"
	mydbu.Database = "go_auth2_users"
	dbbiu = &mydbu
	dbbiu.Connect()

}

func TestUserDBi_AddUser(t *testing.T) {

	var db UserDatabase
	var udb UserDB
	udb.DB = dbbiu
	db = &udb

	var u User
	u.Username = "tester1"
	u.Password = "tester1"
	u.Enabled = true
	u.Entered = time.Now()
	u.Email = "tester1@tester.com"
	u.FirstName = "test"
	u.LastName = "er1"
	u.RoleID = 1
	u.ClientID = 444

	suc := db.AddUser(&u)
	fmt.Println("suc :", suc)
	if !suc {
		t.Fail()
	}
}

func TestUserDBi_UpdateUser(t *testing.T) {

	var db UserDatabase
	var udb UserDB
	udb.DB = dbbiu
	db = &udb

	var u User
	u.Username = "tester1"
	u.Password = "tester11"
	u.Enabled = false
	u.Email = "tester11@tester.com"
	u.FirstName = "test1"
	u.LastName = "er11"
	u.RoleID = 2
	u.ClientID = 444

	suc := db.UpdateUser(&u)
	fmt.Println("suc :", suc)
	if !suc {
		t.Fail()
	}
}

func TestUserDBi_GetUser(t *testing.T) {

	var db UserDatabase
	var udb UserDB
	udb.DB = dbbiu
	db = &udb

	user := db.GetUser("tester1", 444)
	fmt.Println("user: ", user)
	if user.Email != "tester11@tester.com" {
		t.Fail()
	}
}

func TestUserDBi_GetUserList(t *testing.T) {

	var db UserDatabase
	var udb UserDB
	udb.DB = dbbiu
	db = &udb

	users := db.GetUserList()
	fmt.Println("users: ", users)
	if len(*users) == 0 {
		t.Fail()
	}
}

func TestUserDBi_SearchUserList(t *testing.T) {

	var db UserDatabase
	var udb UserDB
	udb.DB = dbbiu
	db = &udb

	users := db.SearchUserList(444)
	fmt.Println("users searched: ", users)
	if len(*users) != 1 {
		t.Fail()
	}
}

func TestUserDBi_DeleteUser(t *testing.T) {
	var db UserDatabase
	var udb UserDB
	udb.DB = dbbiu
	db = &udb

	suc := db.DeleteUser("tester1", 444)
	if !suc {
		t.Fail()
	}
}
