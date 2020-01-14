//Package managers ...
package managers

import (
	"fmt"

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

//AddRole AddRole
func (m *UserManager) AddRole(rl *db.Role) (bool, int64) {
	var rtn bool
	var id int64
	if rl.Role != "" {
		fmt.Println("role: ", rl)
		suc, rid := m.UserDB.AddRole(rl)
		if suc {
			rtn = true
			id = rid
		}
	}
	return rtn, id
}

//GetRole GetRole
func (m *UserManager) GetRole(id int64) *db.Role {
	u := m.UserDB.GetRole(id)
	return u
}

//GetRoleList GetRoleList
func (m *UserManager) GetRoleList() *[]db.Role {
	rl := m.UserDB.GetRoleList()
	return rl
}

//DeleteRole DeleteRole
func (m *UserManager) DeleteRole(id int64) bool {
	rtn := m.UserDB.DeleteRole(id)
	return rtn
}
