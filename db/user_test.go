//Package db ...
package db

import (
	"fmt"
	"testing"
	"time"

	dbi "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

func TestUserDB_AddUser(t *testing.T) {
	var dbbi dbi.Database
	var mydb mdb.MyDBMock
	dbbi = &mydb
	var db UserDatabase
	var udb UserDB
	udb.DB = dbbi
	db = &udb

	var mTestRow dbi.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockInsertSuccess1 = true

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

func TestUserDB_UpdateUser(t *testing.T) {

	var dbbi dbi.Database
	var mydb mdb.MyDBMock
	dbbi = &mydb
	var db UserDatabase
	var udb UserDB
	udb.DB = dbbi
	db = &udb

	var mTestRow dbi.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockUpdateSuccess1 = true

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

func TestUserDB_GetUser(t *testing.T) {

	var dbbi dbi.Database
	var mydb mdb.MyDBMock
	dbbi = &mydb
	var db UserDatabase
	var udb UserDB
	udb.DB = dbbi
	db = &udb

	var mTestRow dbi.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	var mGetRow dbi.DbRow
	mGetRow.Row = []string{"tester1", "password", "1", "2019-12-01", "tester11@tester.com", "test", "er1", "2", "444"}
	mydb.MockRow1 = &mGetRow

	user := db.GetUser("tester1", 444)
	fmt.Println("user: ", user)
	if user.Email != "tester11@tester.com" {
		t.Fail()
	}
}

func TestUserDB_GetUserList(t *testing.T) {

	var dbbi dbi.Database
	var mydb mdb.MyDBMock
	dbbi = &mydb
	var db UserDatabase
	var udb UserDB
	udb.DB = dbbi
	db = &udb

	var mTestRow dbi.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	var rows [][]string
	row1 := []string{"tester1", "password", "1", "2019-12-01", "tester11@tester.com", "test", "er1", "2", "444"}
	rows = append(rows, row1)
	var dbrows dbi.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	users := db.GetUserList()
	fmt.Println("users: ", users)
	if len(*users) != 1 {
		t.Fail()
	}
}

func TestUserDB_SearchUserList(t *testing.T) {

	var dbbi dbi.Database
	var mydb mdb.MyDBMock
	dbbi = &mydb
	var db UserDatabase
	var udb UserDB
	udb.DB = dbbi
	db = &udb

	var mTestRow dbi.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	var rows [][]string
	row1 := []string{"tester1", "password", "1", "2019-12-01", "tester11@tester.com", "test", "er1", "2", "444"}
	rows = append(rows, row1)
	var dbrows dbi.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	users := db.SearchUserList(444)
	fmt.Println("users searched: ", users)
	if len(*users) != 1 {
		t.Fail()
	}
}

func TestUserDB_DeleteUser(t *testing.T) {

	var dbbi dbi.Database
	var mydb mdb.MyDBMock
	dbbi = &mydb
	var db UserDatabase
	var udb UserDB
	udb.DB = dbbi
	db = &udb

	var mTestRow dbi.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow
	mydb.MockDeleteSuccess1 = true

	suc := db.DeleteUser("tester1", 444)
	if !suc {
		t.Fail()
	}
}
