//Package managers ...
package managers

import (
	"fmt"
	"testing"

	db "github.com/Ulbora/GoAuth2Users/db"
	dbi "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

func TestUserManager_AddRole(t *testing.T) {

	var dbbi dbi.Database
	var mydb mdb.MyDBMock
	dbbi = &mydb
	dbbi.Connect()
	//var udbi db.UserDatabase
	var udb db.UserDB
	udb.DB = dbbi
	udbi := udb.GetNew()

	//var man Manager
	var uman UserManager
	uman.UserDB = udbi
	man := uman.GetNew()

	var mTestRow dbi.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 55

	var r db.Role
	r.Role = "test"

	suc, id := man.AddRole(&r)
	if !suc || id == 0 {
		t.Fail()
	}
}

func TestUserManager_GetRole(t *testing.T) {

	var dbbi dbi.Database
	var mydb mdb.MyDBMock

	dbbi = &mydb
	dbbi.Connect()
	//var udbi db.UserDatabase
	var udb db.UserDB
	udb.DB = dbbi
	udbi := udb.GetNew()

	//var man Manager
	var uman UserManager
	uman.UserDB = udbi
	man := uman.GetNew()

	var mTestRow dbi.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	//mydb.MockUpdateSuccess1 = true

	var mGetRow dbi.DbRow
	mGetRow.Row = []string{"3", "someRole"}
	mydb.MockRow1 = &mGetRow

	r := man.GetRole(55)
	if r.ID != 3 || r.Role != "someRole" {
		t.Fail()
	}
}

func TestUserManager_GetRoleList(t *testing.T) {

	var dbbi dbi.Database
	var mydb mdb.MyDBMock

	dbbi = &mydb
	dbbi.Connect()
	//var udbi db.UserDatabase
	var udb db.UserDB
	udb.DB = dbbi
	udbi := udb.GetNew()

	//var man Manager
	var uman UserManager
	uman.UserDB = udbi
	man := uman.GetNew()

	var mTestRow dbi.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	//mydb.MockUpdateSuccess1 = true

	var rows [][]string
	row1 := []string{"2", "someRole"}
	rows = append(rows, row1)
	var dbrows dbi.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	rl := man.GetRoleList()
	fmt.Println("rl: ", rl)
	if len(*rl) != 1 || (*rl)[0].ID != 2 || (*rl)[0].Role != "someRole" {
		t.Fail()
	}
}

func TestUserManager_DeleteRole(t *testing.T) {

	var dbbi dbi.Database
	var mydb mdb.MyDBMock

	dbbi = &mydb
	dbbi.Connect()
	//var udbi db.UserDatabase
	var udb db.UserDB
	udb.DB = dbbi
	udbi := udb.GetNew()

	var mTestRow dbi.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow
	mydb.MockDeleteSuccess1 = true

	//var man Manager
	var uman UserManager
	uman.UserDB = udbi
	man := uman.GetNew()

	suc := man.DeleteRole(5)
	if !suc {
		t.Fail()
	}
}
