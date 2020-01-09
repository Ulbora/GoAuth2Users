//Package handlers ...
package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	jv "github.com/Ulbora/GoAuth2JwtValidator"
	m "github.com/Ulbora/GoAuth2Users/managers"
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

const (
	validationServiceLocal = "http://localhost:3000/rs/token/validate"
)

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
	// GetUserList(w http.ResponseWriter, r *http.Request)
	// SearchUserList(w http.ResponseWriter, r *http.Request)
	// DeleteUser(w http.ResponseWriter, r *http.Request)

	// ClientAddUser(w http.ResponseWriter, r *http.Request)
	// ClientUpdateUser(w http.ResponseWriter, r *http.Request)
	// ClientGetUser(w http.ResponseWriter, r *http.Request)
	// ClientSearchUserList(w http.ResponseWriter, r *http.Request)
	// ClientDeleteUser(w http.ResponseWriter, r *http.Request)

	// LoginUser(w http.ResponseWriter, r *http.Request)
}

//UserHandler UserHandler
type UserHandler struct {
	ValidatorClient jv.Client
	Manager         m.Manager
}

//GetNew GetNew
func (h *UserHandler) GetNew() Handler {
	var hd Handler
	hd = h
	return hd
}

//CheckContent CheckContent
func (h *UserHandler) CheckContent(r *http.Request) bool {
	var rtn bool
	cType := r.Header.Get("Content-Type")
	if cType == "application/json" {
		rtn = true
	}
	return rtn
}

//SetContentType SetContentType
func (h *UserHandler) SetContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

//ProcessBody ProcessBody
func (h *UserHandler) ProcessBody(r *http.Request, obj interface{}) (bool, error) {
	var suc bool
	var err error
	//fmt.Println("r.Body: ", r.Body)
	if r.Body != nil {
		decoder := json.NewDecoder(r.Body)
		//fmt.Println("decoder: ", decoder)
		err = decoder.Decode(obj)
		//fmt.Println("decoder: ", decoder)
		if err != nil {
			log.Println("Decode Error: ", err.Error())
		} else {
			suc = true
		}
	} else {
		err = errors.New("Bad Body")
	}

	return suc, err
}

func (h *UserHandler) getValidationURL() string {
	var url string
	if os.Getenv("VALIDATION_SERVICE") != "" {
		url = os.Getenv("VALIDATION_SERVICE")
	} else {
		url = validationServiceLocal
	}
	return url
}
