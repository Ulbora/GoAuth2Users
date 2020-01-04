//Package managers ...
package managers

import (
	"fmt"
	"time"

	db "github.com/Ulbora/GoAuth2Users/db"
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

//AddUser AddUser
func (m *UserManager) AddUser(us *db.User) bool {
	var rtn bool
	if us.Username != "" && us.Password != "" && us.ClientID > 0 && us.FirstName != "" && us.LastName != "" {
		suc, hpw := m.hashPassword(us.Password)
		if suc {
			us.Password = hpw
			us.Entered = time.Now()
			fmt.Println("user: ", us)
			usuc := m.UserDB.AddUser(us)
			if usuc {
				rtn = true
			}
		}
	}
	return rtn
}

//UpdateUserPassword UpdateUserPassword
func (m *UserManager) UpdateUserPassword(us *db.User) bool {
	var rtn bool
	if us.Username != "" && us.Password != "" && us.ClientID > 0 {
		u := m.UserDB.GetUser(us.Username, us.ClientID)
		suc, hpw := m.hashPassword(us.Password)
		if suc {
			u.Password = hpw
			fmt.Println("user: ", us)
			usuc := m.UserDB.UpdateUser(u)
			if usuc {
				rtn = true
			}
		}
	}
	return rtn
}

//UpdateUserEnabled UpdateUserEnabled
func (m *UserManager) UpdateUserEnabled(us *db.User) bool {
	var rtn bool
	if us.Username != "" && us.ClientID > 0 {
		u := m.UserDB.GetUser(us.Username, us.ClientID)
		u.Enabled = us.Enabled
		fmt.Println("user: ", us)
		usuc := m.UserDB.UpdateUser(u)
		if usuc {
			rtn = true
		}
	}
	return rtn
}

//UpdateUserInfo UpdateUserInfo
func (m *UserManager) UpdateUserInfo(us *db.User) bool {
	var rtn bool
	if us.Username != "" && us.ClientID > 0 && us.Email != "" && us.FirstName != "" && us.LastName != "" && us.RoleID > 0 {
		u := m.UserDB.GetUser(us.Username, us.ClientID)
		u.Email = us.Email
		u.FirstName = us.FirstName
		u.LastName = us.LastName
		u.RoleID = us.RoleID
		fmt.Println("user: ", us)
		usuc := m.UserDB.UpdateUser(u)
		if usuc {
			rtn = true
		}
	}
	return rtn
}

//GetUser GetUser
func (m *UserManager) GetUser(username string, clientID int64) *User {
	var rtn User
	if username != "" && clientID > 0 {
		u := m.UserDB.GetUser(username, clientID)
		rtn.Username = u.Username
		rtn.Enabled = u.Enabled
		rtn.Entered = u.Entered
		rtn.Email = u.Email
		rtn.FirstName = u.FirstName
		rtn.LastName = u.LastName
		rtn.RoleID = u.RoleID
		rtn.ClientID = u.ClientID
	}

	return &rtn
}

//GetUserList GetUserList
func (m *UserManager) GetUserList() *[]UserList {
	var rtn []UserList
	ul := m.UserDB.GetUserList()
	for _, u := range *ul {
		uu := processUserList(&u)
		rtn = append(rtn, *uu)
	}
	return &rtn
}

//SearchUserList SearchUserList
func (m *UserManager) SearchUserList(cid int64) *[]UserList {
	var rtn []UserList
	ul := m.UserDB.SearchUserList(cid)
	for _, u := range *ul {
		uu := processUserList(&u)
		rtn = append(rtn, *uu)
	}
	return &rtn
}

//DeleteUser DeleteUser
func (m *UserManager) DeleteUser(username string, clientID int64) bool {
	var rtn bool
	if username != "" && clientID > 0 {
		rtn = m.UserDB.DeleteUser(username, clientID)
	}
	return rtn
}

func processUserList(u *db.User) *UserList {
	var rtnl UserList
	rtnl.Username = u.Username
	rtnl.Enabled = u.Enabled
	rtnl.FirstName = u.FirstName
	rtnl.LastName = u.LastName
	rtnl.ClientID = u.ClientID
	return &rtnl
}
