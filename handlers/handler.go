//Package handlers ...
package handlers

import "net/http"

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

//ResponseID ResponseID
type ResponseID struct {
	Success bool  `json:"success"`
	ID      int64 `json:"id"`
}

//Response Response
type Response struct {
	Success bool `json:"success"`
}

//Handler Handler
type Handler interface {
	AddRole(w http.ResponseWriter, r *http.Request)
	GetRole(w http.ResponseWriter, r *http.Request)
	GetRoleList(w http.ResponseWriter, r *http.Request)
	DeleteRole(w http.ResponseWriter, r *http.Request)

	AddUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	GetUserList(w http.ResponseWriter, r *http.Request)
	SearchUserList(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)

	LoginUser(w http.ResponseWriter, r *http.Request)

	ClientAddUser(w http.ResponseWriter, r *http.Request)
	ClientUpdateUser(w http.ResponseWriter, r *http.Request)
	ClientGetUser(w http.ResponseWriter, r *http.Request)
	ClientSearchUserList(w http.ResponseWriter, r *http.Request)
	ClientDeleteUser(w http.ResponseWriter, r *http.Request)
}
