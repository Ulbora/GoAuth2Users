//Package managers ...
package managers

import (
	"fmt"
	"testing"

	db "github.com/Ulbora/GoAuth2Users/db"
	dbi "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

func TestUserManager_AddUser(t *testing.T) {

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

	var u db.User
	u.Username = "tester"
	u.Password = "tester"
	u.Enabled = true
	u.Email = "tester@ulboralabs.com"
	u.FirstName = "tester"
	u.LastName = "tester"
	u.RoleID = 1
	u.ClientID = 10
	suc := man.AddUser(&u)
	if !suc {
		t.Fail()
	}
}

func TestUserManager_UpdateUserPw(t *testing.T) {

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

	mydb.MockUpdateSuccess1 = true

	var mGetRow dbi.DbRow
	mGetRow.Row = []string{"tester1", "password", "1", "2019-12-01", "tester11@tester.com", "test", "er1", "2", "444"}
	mydb.MockRow1 = &mGetRow

	var u db.User
	u.Username = "tester"
	u.Password = "tester2"
	// u.Enabled = true
	// u.Email = "tester@ulboralabs.com"
	// u.FirstName = "tester"
	// u.LastName = "tester"
	//u.RoleID = 1
	u.ClientID = 10
	suc := man.UpdateUserPassword(&u)
	//foundU := udbi.GetUser("tester", 10)
	//pwmatch := uman.validatePassword("tester2", foundU.Password)
	if !suc {
		t.Fail()
	}
}

func TestUserManager_UpdateUserEnabled(t *testing.T) {

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

	mydb.MockUpdateSuccess1 = true

	var mGetRow dbi.DbRow
	mGetRow.Row = []string{"tester1", "password", "1", "2019-12-01", "tester11@tester.com", "test", "er1", "2", "444"}
	mydb.MockRow1 = &mGetRow

	var u db.User
	u.Username = "tester"
	//u.Password = "tester2"
	u.Enabled = false
	// u.Email = "tester@ulboralabs.com"
	// u.FirstName = "tester"
	// u.LastName = "tester"
	//u.RoleID = 1
	u.ClientID = 10
	suc := man.UpdateUserEnabled(&u)
	//foundU := udbi.GetUser("tester", 10)

	if !suc {
		t.Fail()
	}
}

func TestUserManager_GetUser(t *testing.T) {

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
	mGetRow.Row = []string{"tester1", "password", "1", "2019-12-01", "tester@tester.com", "test", "er1", "2", "444"}
	mydb.MockRow1 = &mGetRow

	u := man.GetUser("tester", 10)
	if u.Email != "tester@tester.com" {
		t.Fail()
	}
}

func TestUserManager_GetUserList(t *testing.T) {

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
	row1 := []string{"tester", "password", "1", "2019-12-01", "tester11@tester.com", "tester", "tester", "2", "444"}
	rows = append(rows, row1)
	var dbrows dbi.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	ul := man.GetUserList()
	fmt.Println("ul: ", ul)
	if len(*ul) != 1 || (*ul)[0].Username != "tester" || (*ul)[0].FirstName != "tester" || (*ul)[0].LastName != "tester" {
		t.Fail()
	}
}

func TestUserManager_SearchUserList(t *testing.T) {

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
	row1 := []string{"tester", "password", "1", "2019-12-01", "tester11@tester.com", "tester", "tester", "2", "444"}
	rows = append(rows, row1)
	var dbrows dbi.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	ul := man.SearchUserList(10)
	fmt.Println("ul: ", ul)
	if len(*ul) != 1 || (*ul)[0].Username != "tester" || (*ul)[0].FirstName != "tester" || (*ul)[0].LastName != "tester" {
		t.Fail()
	}
}

func TestUserManager_UpdateUserInfo(t *testing.T) {

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

	mydb.MockUpdateSuccess1 = true

	var mGetRow dbi.DbRow
	mGetRow.Row = []string{"tester1", "password", "1", "2019-12-01", "tester11@tester.com", "test", "er1", "2", "444"}
	mydb.MockRow1 = &mGetRow

	var u db.User
	u.Username = "tester"
	//u.Password = "tester2"
	//u.Enabled = false
	u.Email = "tester2@ulboralabs.com"
	u.FirstName = "tester2"
	u.LastName = "tester22"
	u.RoleID = 1
	u.ClientID = 10
	suc := man.UpdateUserInfo(&u)
	//foundU := udbi.GetUser("tester", 10)

	if !suc {
		t.Fail()
	}
}

func TestUserManager_DeleteUser(t *testing.T) {

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

	suc := man.DeleteUser("tester", 10)
	if !suc {
		t.Fail()
	}
}

func TestUserManager_LoginUser(t *testing.T) {

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

	var mGetRow dbi.DbRow
	mGetRow.Row = []string{"tester", "$2a$10$Cp5LWuqgayns7.Fox1hCiuQw.Ya3nmAgOH7GMYfVDWQYCirICIioS", "1", "2019-12-01", "tester@tester.com", "test", "er1", "2", "444"}
	mydb.MockRow1 = &mGetRow

	//var man Manager
	var uman UserManager
	uman.UserDB = udbi
	man := uman.GetNew()

	suc := man.ValidateUser("tester", "tester123", 10)
	if !suc {
		t.Fail()
	}
}
