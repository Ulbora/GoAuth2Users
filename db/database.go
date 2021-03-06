//Package db ...
package db

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

import (
	"fmt"
	"log"
	"strconv"
	"time"

	dbi "github.com/Ulbora/dbinterface"
)

const (
	//TimeFormat TimeFormat
	TimeFormat = "2006-01-02"
)

//UserDatabase UserDatabase
type UserDatabase interface {
	AddRole(rl *Role) (bool, int64)
	GetRole(id int64) *Role
	GetRoleList() *[]Role
	DeleteRole(id int64) bool

	AddUser(us *User) bool
	UpdateUser(us *User) bool
	GetUser(username string, clientID int64) *User
	GetUserList() *[]User
	SearchUserList(cid int64) *[]User
	DeleteUser(username string, clientID int64) bool
}

//UserDB UserDB
type UserDB struct {
	DB dbi.Database
}

//Role Role
type Role struct {
	ID   int64  `json:"id"`
	Role string `json:"role"`
}

//User User
type User struct {
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Enabled   bool      `json:"enabled"`
	Entered   time.Time `json:"dateEntered"`
	Email     string    `json:"emailAddress"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	RoleID    int64     `json:"roleId"`
	ClientID  int64     `json:"clientId"`
}

func (d *UserDB) testConnection() bool {
	log.Println("in testConnection")
	var rtn = false
	var a []interface{}
	log.Println("d.DB: ", d.DB)
	rowPtr := d.DB.Test(connectTest, a...)
	fmt.Println("rowPtr", rowPtr)
	log.Println("after testConnection test", rowPtr)
	if len(rowPtr.Row) != 0 {
		foundRow := rowPtr.Row
		int64Val, err := strconv.ParseInt(foundRow[0], 10, 0)
		//log.Print("Records found during test ")
		//log.Println("Records found during test :", int64Val)
		if err != nil {
			log.Print(err)
		}
		if int64Val >= 0 {
			rtn = true
		}
	}
	return rtn
}

//GetNew GetNew
func (d *UserDB) GetNew() UserDatabase {
	var db UserDatabase
	db = d
	return db
}
