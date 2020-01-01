//Package db ...
package db

import (
	"fmt"
	"strconv"
	"time"
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
func (d *UserDB) AddUser(us *User) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, us.Username, us.Password, us.Enabled, us.Entered, us.Email, us.FirstName, us.LastName, us.RoleID, us.ClientID)
	suc, _ := d.DB.Insert(insertUser, a...)
	return suc
}

//UpdateUser UpdateUser
func (d *UserDB) UpdateUser(us *User) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, us.Password, us.Enabled, us.Email, us.FirstName, us.LastName, us.RoleID, us.Username, us.ClientID)
	suc := d.DB.Update(updateUser, a...)
	return suc
}

//GetUser GetUser
func (d *UserDB) GetUser(username string, clientID int64) *User {
	var rtn *User
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, username, clientID)
	row := d.DB.Get(getUser, a...)
	if row != nil && len(row.Row) != 0 {
		foundRow := row.Row
		rtn = parseUserRow(&foundRow)
	}
	return rtn
}

//GetUserList GetUserList
func (d *UserDB) GetUserList() *[]User {
	var rtn []User
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	rows := d.DB.GetList(getUserList, a...)
	fmt.Println("user rows: ", rows)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := parseUserRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//SearchUserList SearchUserList
func (d *UserDB) SearchUserList(cid int64) *[]User {
	var rtn []User
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, cid)
	rows := d.DB.GetList(searchUserList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := parseUserRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//DeleteUser DeleteUser
func (d *UserDB) DeleteUser(username string, clientID int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, username, clientID)
	suc := d.DB.Delete(deleteUser, a...)
	return suc
}

func parseUserRow(foundRow *[]string) *User {
	var rtn User
	if len(*foundRow) > 0 {
		int64Valr, err := strconv.ParseInt((*foundRow)[7], 10, 64)
		if err == nil {
			int64Valc, err := strconv.ParseInt((*foundRow)[8], 10, 64)
			if err == nil {
				cTime, err := time.Parse(TimeFormat, (*foundRow)[3])
				fmt.Println("time error:", err)
				if err == nil {
					enab, err := strconv.ParseBool((*foundRow)[2])
					if err == nil {
						rtn.Username = (*foundRow)[0]
						rtn.Password = (*foundRow)[1]
						rtn.Enabled = enab
						rtn.Entered = cTime
						rtn.Email = (*foundRow)[4]
						rtn.FirstName = (*foundRow)[5]
						rtn.LastName = (*foundRow)[6]
						rtn.RoleID = int64Valr
						rtn.ClientID = int64Valc
					}
				}
			}
		}
	}
	return &rtn
}
