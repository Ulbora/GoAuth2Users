//Package db ...
package db

import "strconv"

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
func (d *UserDB) AddRole(rl *Role) (bool, int64) {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, rl.Role)
	suc, id := d.DB.Insert(insertRole, a...)
	return suc, id
}

//GetRole GetRole
func (d *UserDB) GetRole(id int64) *Role {
	var rtn *Role
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	row := d.DB.Get(getRole, a...)
	if row != nil && len(row.Row) != 0 {
		foundRow := row.Row
		rtn = parseRoleRow(&foundRow)
	}
	return rtn
}

//GetRoleList GetRoleList
func (d *UserDB) GetRoleList() *[]Role {
	var rtn []Role
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	rows := d.DB.GetList(getRoleList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := parseRoleRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//DeleteRole DeleteRole
func (d *UserDB) DeleteRole(id int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	suc := d.DB.Delete(deleteRole, a...)
	return suc
}

func parseRoleRow(foundRow *[]string) *Role {
	var rtn Role
	if len(*foundRow) > 0 {
		int64Val, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		if err == nil {
			rtn.ID = int64Val
			rtn.Role = (*foundRow)[1]
		}
	}
	return &rtn
}
