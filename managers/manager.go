//Package managers ...
package managers

import (
	"time"

	db "github.com/Ulbora/GoAuth2Users/db"
	lg "github.com/Ulbora/Level_Logger"
	"golang.org/x/crypto/bcrypt"
)

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

//Manager Manager
type Manager interface {
	AddRole(rl *db.Role) (bool, int64)
	GetRole(id int64) *db.Role
	GetRoleList() *[]db.Role
	DeleteRole(id int64) bool

	AddUser(us *db.User) bool
	UpdateUserPassword(us *db.User) bool
	UpdateUserEnabled(us *db.User) bool
	UpdateUserInfo(us *db.User) bool
	GetUser(username string, clientID int64) *User
	GetUserList() *[]UserList
	SearchUserList(cid int64) *[]UserList
	DeleteUser(username string, clientID int64) bool

	ValidateUser(username string, password string, clientID int64) bool
}

//UserManager UserManager
type UserManager struct {
	UserDB db.UserDatabase
	Log    *lg.Logger
}

//User User
type User struct {
	Username  string    `json:"username"`
	Enabled   bool      `json:"enabled"`
	Entered   time.Time `json:"dateEntered"`
	Email     string    `json:"emailAddress"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	RoleID    int64     `json:"roleId"`
	ClientID  int64     `json:"clientId"`
}

//UserList UserList
type UserList struct {
	Username  string `json:"username"`
	Enabled   bool   `json:"enabled"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	ClientID  int64  `json:"clientId"`
}

func (m *UserManager) hashPassword(pw string) (bool, string) {
	var suc bool
	var rtn string
	hashedPw, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err == nil {
		rtn = string(hashedPw)
		suc = true
	}
	return suc, rtn
}

func (m *UserManager) validatePassword(pw string, hpw string) bool {
	var suc bool
	err := bcrypt.CompareHashAndPassword([]byte(hpw), []byte(pw))
	if err == nil {
		suc = true
	}
	return suc
}

//GetNew GetNew
func (m *UserManager) GetNew() Manager {
	var man Manager
	man = m
	return man
}
