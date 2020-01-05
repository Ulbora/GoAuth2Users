//Package managers ...
package managers

import db "github.com/Ulbora/GoAuth2Users/db"

/*
 Copyright (C) 2019 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2019 Ken Williamson
 All rights reserved.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.
 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.
 You should have received a copy of the GNU General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.

*/

//MockUserManager MockUserManager
type MockUserManager struct {
	MockInsertRoleSuc bool
	MockInsertRoleID  int64
	MockRole          *db.Role
	MockRoleList      *[]db.Role
	MockDeleteRoleSuc bool

	MockInsertUserSuc     bool
	MockUpdatePasswordSuc bool
	MockUpdateEnableSuc   bool
	MockUpdateInfoSuc     bool
	MockUser              *User
	MockUserList          *[]UserList
	MockDeleteUserSuc     bool
}

//AddRole AddRole
func (m *MockUserManager) AddRole(rl *db.Role) (bool, int64) {
	return m.MockInsertRoleSuc, m.MockInsertRoleID
}

//GetRole GetRole
func (m *MockUserManager) GetRole(id int64) *db.Role {
	return m.MockRole
}

//GetRoleList GetRoleList
func (m *MockUserManager) GetRoleList() *[]db.Role {
	return m.MockRoleList
}

//DeleteRole DeleteRole
func (m *MockUserManager) DeleteRole(id int64) bool {
	return m.MockDeleteRoleSuc
}

//AddUser AddUser
func (m *MockUserManager) AddUser(us *db.User) bool {
	return m.MockInsertUserSuc
}

//UpdateUserPassword UpdateUserPassword
func (m *MockUserManager) UpdateUserPassword(us *db.User) bool {
	return m.MockUpdatePasswordSuc
}

//UpdateUserEnabled UpdateUserEnabled
func (m *MockUserManager) UpdateUserEnabled(us *db.User) bool {
	return m.MockUpdateEnableSuc
}

//UpdateUserInfo UpdateUserInfo
func (m *MockUserManager) UpdateUserInfo(us *db.User) bool {
	return m.MockUpdateInfoSuc
}

//GetUser GetUser
func (m *MockUserManager) GetUser(username string, clientID int64) *User {
	return m.MockUser
}

//GetUserList GetUserList
func (m *MockUserManager) GetUserList() *[]UserList {
	return m.MockUserList
}

//SearchUserList SearchUserList
func (m *MockUserManager) SearchUserList(cid int64) *[]UserList {
	return m.MockUserList
}

//DeleteUser DeleteUser
func (m *MockUserManager) DeleteUser(username string, clientID int64) bool {
	return m.MockDeleteUserSuc
}

//GetNew GetNew
func (m *MockUserManager) GetNew() Manager {
	var man Manager
	man = m
	return man
}
