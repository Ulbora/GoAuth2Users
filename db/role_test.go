//Package db ...
package db

import (
	"fmt"
	"testing"

	dbi "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

func TestUserDBRole_AddRole(t *testing.T) {
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
	mydb.MockInsertID1 = 1

	var r Role
	r.Role = "testRole"
	suc, id := db.AddRole(&r)
	if !suc && id != 1 {
		t.Fail()
	}
}

func TestUserDBRole_AddRoleTestRow(t *testing.T) {
	var dbbi dbi.Database
	var mydb mdb.MyDBMock
	dbbi = &mydb
	var db UserDatabase
	var udb UserDB
	udb.DB = dbbi
	db = &udb

	var mTestRow dbi.DbRow
	mTestRow.Row = []string{"1"}
	mydb.MockTestRow = &mTestRow

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 1

	var r Role
	r.Role = "testRole"
	suc, id := db.AddRole(&r)
	if !suc && id != 1 {
		t.Fail()
	}
}

func TestUserDBRole_AddRoleBadTest(t *testing.T) {
	var dbbi dbi.Database
	var mydb mdb.MyDBMock
	dbbi = &mydb
	var db UserDatabase
	var udb UserDB
	udb.DB = dbbi
	db = &udb

	var mTestRow dbi.DbRow
	mTestRow.Row = []string{"a"}
	mydb.MockTestRow = &mTestRow

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 1

	var r Role
	r.Role = "testRole"
	suc, id := db.AddRole(&r)
	if !suc && id != 1 {
		t.Fail()
	}
}

func TestUserDBRole_GetRole(t *testing.T) {
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
	mGetRow.Row = []string{"1", "testRole"}
	mydb.MockRow1 = &mGetRow

	role := db.GetRole(1)
	if role.Role != "testRole" {
		t.Fail()
	}
}

func TestUserDBRole_GetRoleList(t *testing.T) {
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
	row1 := []string{"1", "testRole"}
	rows = append(rows, row1)
	var dbrows dbi.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	roles := db.GetRoleList()
	fmt.Println("roles: ", roles)
	if len(*roles) != 1 {
		t.Fail()
	}
}

func TestUserDBRole_DeleteRole(t *testing.T) {
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

	var r Role
	r.Role = "testRole"
	suc := db.DeleteRole(1)
	if !suc {
		t.Fail()
	}
}
