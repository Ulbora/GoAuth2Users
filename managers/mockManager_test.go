//Package managers ...
package managers

import (
	"fmt"
	"testing"

	db "github.com/Ulbora/GoAuth2Users/db"
)

func TestMockUserManager_AddRole(t *testing.T) {
	var um MockUserManager
	man := um.GetNew()

	um.MockInsertRoleSuc = true
	um.MockInsertRoleID = 33

	var r db.Role
	r.Role = "test"

	suc, id := man.AddRole(&r)
	if !suc || id == 0 {
		t.Fail()
	}

}

func TestMockUserManager_GetRole(t *testing.T) {
	var um MockUserManager
	man := um.GetNew()
	var mr db.Role
	mr.ID = 3
	mr.Role = "someRole"
	um.MockRole = &mr

	r := man.GetRole(55)
	if r.ID != 3 || r.Role != "someRole" {
		t.Fail()
	}
}

func TestMockUserManager_GetRoleList(t *testing.T) {
	var um MockUserManager
	man := um.GetNew()

	var mr db.Role
	mr.ID = 2
	mr.Role = "someRole"
	var rll = []db.Role{mr}
	um.MockRoleList = &rll

	rl := man.GetRoleList()
	fmt.Println("rl: ", rl)
	if len(*rl) != 1 || (*rl)[0].ID != 2 || (*rl)[0].Role != "someRole" {
		t.Fail()
	}
}

func TestMockUserManager_DeleteRole(t *testing.T) {
	var um MockUserManager
	man := um.GetNew()

	um.MockDeleteRoleSuc = true

	suc := man.DeleteRole(5)
	if !suc {
		t.Fail()
	}
}

func TestMockUserManager_AddUser(t *testing.T) {
	var um MockUserManager
	um.MockInsertUserSuc = true
	m := um.GetNew()

	var u db.User
	u.Username = "tester"
	u.Password = "tester"
	u.Enabled = true
	u.Email = "tester@ulboralabs.com"
	u.FirstName = "tester"
	u.LastName = "tester"
	u.RoleID = 1
	u.ClientID = 10
	suc := m.AddUser(&u)
	if !suc {
		t.Fail()
	}
}

func TestMockUserManager_UpdateUserPw(t *testing.T) {
	var um MockUserManager
	man := um.GetNew()

	um.MockUpdatePasswordSuc = true

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

func TestMockUserManager_UpdateUserEn(t *testing.T) {
	var um MockUserManager
	man := um.GetNew()

	um.MockUpdateEnableSuc = true

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

func TestMockUserManager_UpdateUserInfo(t *testing.T) {
	var um MockUserManager
	man := um.GetNew()

	um.MockUpdateInfoSuc = true

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

func TestMockUserManager_GetUser(t *testing.T) {
	var um MockUserManager
	man := um.GetNew()

	var mu User
	mu.Username = "tester1"
	mu.Email = "tester@tester.com"

	um.MockUser = &mu

	u := man.GetUser("tester", 10)
	if u.Email != "tester@tester.com" {
		t.Fail()
	}
}

func TestMockUserManager_GetUserList(t *testing.T) {
	var um MockUserManager
	man := um.GetNew()

	var mul UserList
	mul.Username = "tester"
	mul.FirstName = "tester"
	mul.LastName = "tester"

	var mull = []UserList{mul}
	um.MockUserList = &mull

	ul := man.GetUserList()
	fmt.Println("ul: ", ul)
	if len(*ul) != 1 || (*ul)[0].Username != "tester" || (*ul)[0].FirstName != "tester" || (*ul)[0].LastName != "tester" {
		t.Fail()
	}
}

func TestMockUserManager_SearchUserList(t *testing.T) {
	var um MockUserManager
	man := um.GetNew()

	var mul UserList
	mul.Username = "tester"
	mul.FirstName = "tester"
	mul.LastName = "tester"

	var mull = []UserList{mul}
	um.MockUserList = &mull

	ul := man.SearchUserList(10)
	fmt.Println("ul: ", ul)
	if len(*ul) != 1 || (*ul)[0].Username != "tester" || (*ul)[0].FirstName != "tester" || (*ul)[0].LastName != "tester" {
		t.Fail()
	}
}

func TestMockUserManager_DeleteUser(t *testing.T) {
	var um MockUserManager
	man := um.GetNew()

	um.MockDeleteUserSuc = true

	suc := man.DeleteUser("tester", 55)
	if !suc {
		t.Fail()
	}
}
