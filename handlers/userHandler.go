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

//AddUser AddUser
func (h *UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {

	var addRoleURL = "/ulbora/rs/user/add"

	var auscl jv.Claim
	auscl.Role = "superAdmin"
	auscl.URL = addRoleURL
	auscl.Scope = "write"
	//fmt.Println("client: ", h.Validator)
	auth := h.ValidatorClient.Authorize(r, &auscl, h.getValidationURL())
	if auth {
		h.SetContentType(w)
		aasURIContOk := h.CheckContent(r)
		fmt.Println("conOk: ", aasURIContOk)
		if !aasURIContOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var us db.User
			aussuc, uaerr := h.ProcessBody(r, &us)
			fmt.Println("aussuc: ", aussuc)
			fmt.Println("us: ", us)
			fmt.Println("uaerr: ", uaerr)
			if !aussuc && uaerr != nil {
				http.Error(w, uaerr.Error(), http.StatusBadRequest)
			} else {
				ausSuc := h.Manager.AddUser(&us)
				fmt.Println("ausSuc: ", ausSuc)
				var rtn Response
				if ausSuc {
					rtn.Success = ausSuc
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(rtn)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var fausrtn Response
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(fausrtn)
		fmt.Fprint(w, string(resJSON))
	}
}

//UpdateUser UpdateUser
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	//url of this endpoint
	fmt.Println("inside UpdateUser------------------------------------")

	var upsAuURL = "/ulbora/rs/user/update"

	var uprlcl jv.Claim
	uprlcl.Role = "superAdmin"
	uprlcl.URL = upsAuURL
	uprlcl.Scope = "write"
	//fmt.Println("client: ", h.Client)
	auth := h.ValidatorClient.Authorize(r, &uprlcl, h.getValidationURL())
	fmt.Println("auth: ", auth)
	if auth {
		// w.Header().Set("Content-Type", "application/json")
		h.SetContentType(w)
		uPasURIContOk := h.CheckContent(r)
		fmt.Println("conOk: ", uPasURIContOk)
		if !uPasURIContOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var upus db.User
			upusbsuc, upuserr := h.ProcessBody(r, &upus)
			fmt.Println("upusbsuc: ", upusbsuc)
			fmt.Println("upus: ", upus)
			fmt.Println("upuserr: ", upuserr)
			if !upusbsuc && upuserr != nil {
				http.Error(w, upuserr.Error(), http.StatusBadRequest)
			} else {
				var upussuc bool
				if upus.Password != "" {
					fmt.Println("in password ")
					upussuc = h.Manager.UpdateUserPassword(&upus)
				} else if upus.FirstName != "" && upus.LastName != "" && upus.Email != "" && upus.RoleID != 0 {
					fmt.Println("in info ")
					upussuc = h.Manager.UpdateUserInfo(&upus)
				} else {
					fmt.Println("in enabled ")
					upussuc = h.Manager.UpdateUserEnabled(&upus)
				}
				fmt.Println("upussuc: ", upussuc)
				var rtn Response
				if upussuc {
					rtn.Success = upussuc
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(rtn)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var fusuprtn Response
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(fusuprtn)
		fmt.Fprint(w, string(resJSON))
	}
}

//GetUser GetUser
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	var getAuURL = "/ulbora/rs/user/get"

	var gusrcl jv.Claim
	gusrcl.Role = "superAdmin"
	gusrcl.URL = getAuURL
	gusrcl.Scope = "read"
	//fmt.Println("client: ", h.Client)
	auth := h.ValidatorClient.Authorize(r, &gusrcl, h.getValidationURL())
	if auth {
		//var id string
		h.SetContentType(w)
		vars := mux.Vars(r)
		fmt.Println("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			var cidStr = vars["clientId"]
			var usernm = vars["username"]
			fmt.Println("vars: ", vars)
			cid, cidErr := strconv.ParseInt(cidStr, 10, 64)
			if cid != 0 && cidErr == nil && usernm != "" {
				fmt.Println("cid: ", cid)
				getUsr := h.Manager.GetUser(usernm, cid)
				fmt.Println("getUsr: ", getUsr)
				w.WriteHeader(http.StatusOK)
				resJSON, _ := json.Marshal(getUsr)
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
