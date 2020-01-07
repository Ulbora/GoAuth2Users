//Package handlers ...
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	jv "github.com/Ulbora/GoAuth2JwtValidator"
	db "github.com/Ulbora/GoAuth2Users/db"
	"github.com/gorilla/mux"
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
func (h *UserHandler) AddRole(w http.ResponseWriter, r *http.Request) {

	var addRoleURL = "/ulbora/rs/role/add"

	var arlcl jv.Claim
	arlcl.Role = "superAdmin"
	arlcl.URL = addRoleURL
	arlcl.Scope = "write"
	//fmt.Println("client: ", h.Validator)
	auth := h.ValidatorClient.Authorize(r, &arlcl, h.getValidationURL())
	if auth {
		h.SetContentType(w)
		aasURIContOk := h.CheckContent(r)
		fmt.Println("conOk: ", aasURIContOk)
		if !aasURIContOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var rl db.Role
			rasuc, raerr := h.ProcessBody(r, &rl)
			fmt.Println("rasuc: ", rasuc)
			fmt.Println("rl: ", rl)
			fmt.Println("raerr: ", raerr)
			if !rasuc && raerr != nil {
				http.Error(w, raerr.Error(), http.StatusBadRequest)
			} else {
				arlSuc, arlID := h.Manager.AddRole(&rl)
				fmt.Println("arlSuc: ", arlSuc)
				fmt.Println("arlID: ", arlID)
				var rtn ResponseID
				if arlSuc && arlID != 0 {
					rtn.Success = arlSuc
					rtn.ID = arlID
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(rtn)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var farlrtn ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(farlrtn)
		fmt.Fprint(w, string(resJSON))
	}
}

//GetRole GetRole
func (h *UserHandler) GetRole(w http.ResponseWriter, r *http.Request) {
	var getAuURL = "/ulbora/rs/role/get"

	var grlcl jv.Claim
	grlcl.Role = "superAdmin"
	grlcl.URL = getAuURL
	grlcl.Scope = "read"
	//fmt.Println("client: ", h.Client)
	auth := h.ValidatorClient.Authorize(r, &grlcl, h.getValidationURL())
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		fmt.Println("vars: ", len(vars))
		if vars != nil && len(vars) != 0 {
			var idStr = vars["id"]
			fmt.Println("vars: ", vars)
			id, idErr := strconv.ParseInt(idStr, 10, 64)
			if id != 0 && idErr == nil {
				fmt.Println("id: ", id)
				getRl := h.Manager.GetRole(id)
				fmt.Println("getRl: ", getRl)
				w.WriteHeader(http.StatusOK)
				resJSON, _ := json.Marshal(getRl)
				fmt.Fprint(w, string(resJSON))
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//GetRoleList GetRoleList
func (h *UserHandler) GetRoleList(w http.ResponseWriter, r *http.Request) {
	var getAulURL = "/ulbora/rs/role/list"

	var grllcl jv.Claim
	grllcl.Role = "admin"
	grllcl.URL = getAulURL
	grllcl.Scope = "read"
	auth := h.ValidatorClient.Authorize(r, &grllcl, h.getValidationURL())
	if auth {
		h.SetContentType(w)
		getRll := h.Manager.GetRoleList()
		fmt.Println("getRll: ", getRll)
		w.WriteHeader(http.StatusOK)
		resJSON, _ := json.Marshal(getRll)
		fmt.Fprint(w, string(resJSON))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//DeleteRole DeleteRole
func (h *UserHandler) DeleteRole(w http.ResponseWriter, r *http.Request) {
	var getAudURL = "/ulbora/rs/role/delete"

	var rldcl jv.Claim
	rldcl.Role = "superAdmin"
	rldcl.URL = getAudURL
	rldcl.Scope = "write"
	//fmt.Println("client: ", h.Client)
	auth := h.ValidatorClient.Authorize(r, &rldcl, h.getValidationURL())
	if auth {
		//var id string
		h.SetContentType(w)
		vars := mux.Vars(r)
		fmt.Println("vars: ", len(vars))
		if vars != nil && len(vars) != 0 {
			var idStr = vars["id"]
			fmt.Println("vars delete: ", vars)
			id, idErr := strconv.ParseInt(idStr, 10, 64)
			fmt.Println("id delete: ", id)
			if id != 0 && idErr == nil {
				fmt.Println("id: ", id)
				rld := h.Manager.DeleteRole(id)
				var rtn Response
				if rld {
					rtn.Success = rld
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(rtn)
				fmt.Fprint(w, string(resJSON))
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
