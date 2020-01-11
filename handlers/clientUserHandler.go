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

//ClientAddUser ClientAddUser
func (h *UserHandler) ClientAddUser(w http.ResponseWriter, r *http.Request) {

	var addRoleURL = "/ulbora/rs/client/user/add"

	var causcl jv.Claim
	causcl.Role = "admin"
	causcl.URL = addRoleURL
	causcl.Scope = "write"
	//fmt.Println("client: ", h.Validator)
	auth := h.ValidatorClient.Authorize(r, &causcl, h.getValidationURL())
	if auth {
		h.SetContentType(w)
		caasURIContOk := h.CheckContent(r)
		fmt.Println("conOk: ", caasURIContOk)
		if !caasURIContOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var us db.User
			caussuc, uaerr := h.ProcessBody(r, &us)
			cidStr := r.Header.Get("clientId")
			cid, cidErr := strconv.ParseInt(cidStr, 10, 64)
			fmt.Println("aussuc: ", caussuc)
			fmt.Println("us: ", us)
			fmt.Println("uaerr: ", uaerr)
			if !caussuc && uaerr != nil {
				http.Error(w, uaerr.Error(), http.StatusBadRequest)
			} else if cid == 0 && cidErr != nil {
				http.Error(w, cidErr.Error(), http.StatusBadRequest)
			} else {
				us.ClientID = cid
				causSuc := h.Manager.AddUser(&us)
				fmt.Println("us after clientID: ", us)
				fmt.Println("causSuc: ", causSuc)
				var rtn Response
				if causSuc {
					rtn.Success = causSuc
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

//ClientUpdateUser ClientUpdateUser
func (h *UserHandler) ClientUpdateUser(w http.ResponseWriter, r *http.Request) {
	//url of this endpoint
	fmt.Println("inside UpdateUser------------------------------------")

	var upsAuURL = "/ulbora/rs/client/user/update"

	var cuprlcl jv.Claim
	cuprlcl.Role = "admin"
	cuprlcl.URL = upsAuURL
	cuprlcl.Scope = "write"
	//fmt.Println("client: ", h.Client)
	auth := h.ValidatorClient.Authorize(r, &cuprlcl, h.getValidationURL())
	fmt.Println("auth: ", auth)
	if auth {
		// w.Header().Set("Content-Type", "application/json")
		h.SetContentType(w)
		uPasURIContOk := h.CheckContent(r)
		fmt.Println("conOk: ", uPasURIContOk)
		if !uPasURIContOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var cupus db.User
			cupusbsuc, upuserr := h.ProcessBody(r, &cupus)
			cidStr := r.Header.Get("clientId")
			cid, cidErr := strconv.ParseInt(cidStr, 10, 64)
			fmt.Println("cupusbsuc: ", cupusbsuc)
			fmt.Println("upus: ", cupus)
			fmt.Println("upuserr: ", upuserr)
			if !cupusbsuc && upuserr != nil {
				http.Error(w, upuserr.Error(), http.StatusBadRequest)
			} else if cid == 0 && cidErr != nil {
				http.Error(w, cidErr.Error(), http.StatusBadRequest)
			} else {
				cupus.ClientID = cid
				fmt.Println("cupus after clientID: ", cupus)
				var cupussuc bool
				if cupus.Password != "" {
					fmt.Println("in password ")
					cupussuc = h.Manager.UpdateUserPassword(&cupus)
				} else if cupus.FirstName != "" && cupus.LastName != "" && cupus.Email != "" && cupus.RoleID != 0 {
					fmt.Println("in info ")
					cupussuc = h.Manager.UpdateUserInfo(&cupus)
				} else {
					fmt.Println("in enabled ")
					cupussuc = h.Manager.UpdateUserEnabled(&cupus)
				}
				fmt.Println("cupussuc: ", cupussuc)
				var rtn Response
				if cupussuc {
					rtn.Success = cupussuc
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(rtn)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var cfusuprtn Response
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(cfusuprtn)
		fmt.Fprint(w, string(resJSON))
	}
}

//ClientGetUser ClientGetUser
func (h *UserHandler) ClientGetUser(w http.ResponseWriter, r *http.Request) {
	var getcusURL = "/ulbora/rs/client/user/get"

	var cgusrcl jv.Claim
	cgusrcl.Role = "admin"
	cgusrcl.URL = getcusURL
	cgusrcl.Scope = "read"
	//fmt.Println("client: ", h.Client)
	auth := h.ValidatorClient.Authorize(r, &cgusrcl, h.getValidationURL())
	if auth {
		//var id string
		h.SetContentType(w)
		vars := mux.Vars(r)
		fmt.Println("vars: ", len(vars))
		if vars != nil && len(vars) == 1 {
			//var cidStr = vars["clientId"]
			ccidStr := r.Header.Get("clientId")
			var cusernm = vars["username"]
			fmt.Println("vars: ", vars)
			cid, cidErr := strconv.ParseInt(ccidStr, 10, 64)
			if cid != 0 && cidErr == nil && cusernm != "" {
				fmt.Println("cid: ", cid)
				getUsr := h.Manager.GetUser(cusernm, cid)
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

//ClientSearchUserList ClientSearchUserList
func (h *UserHandler) ClientSearchUserList(w http.ResponseWriter, r *http.Request) {
	var getCusURL = "/ulbora/rs/client/user/search"

	var cusrcl jv.Claim
	cusrcl.Role = "admin"
	cusrcl.URL = getCusURL
	cusrcl.Scope = "read"
	//fmt.Println("client: ", h.Client)
	auth := h.ValidatorClient.Authorize(r, &cusrcl, h.getValidationURL())
	if auth {
		h.SetContentType(w)
		clcidStr := r.Header.Get("clientId")
		clcid, cidErr := strconv.ParseInt(clcidStr, 10, 64)
		if clcid != 0 && cidErr == nil {
			fmt.Println("cid: ", clcid)
			sUsr := h.Manager.SearchUserList(clcid)
			fmt.Println("sUsr: ", sUsr)
			w.WriteHeader(http.StatusOK)
			resJSON, _ := json.Marshal(sUsr)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//ClientDeleteUser ClientDeleteUser
func (h *UserHandler) ClientDeleteUser(w http.ResponseWriter, r *http.Request) {
	var getcudURL = "/ulbora/rs/client/user/delete"

	var dcusrcl jv.Claim
	dcusrcl.Role = "admin"
	dcusrcl.URL = getcudURL
	dcusrcl.Scope = "write"
	//fmt.Println("client: ", h.Client)
	auth := h.ValidatorClient.Authorize(r, &dcusrcl, h.getValidationURL())
	if auth {
		//var id string
		h.SetContentType(w)
		vars := mux.Vars(r)
		fmt.Println("vars: ", len(vars))
		if vars != nil && len(vars) == 1 {
			//var dcidStr = vars["clientId"]
			var dusernm = vars["username"]
			fmt.Println("vars: ", vars)
			dccidStr := r.Header.Get("clientId")
			dccid, cidErr := strconv.ParseInt(dccidStr, 10, 64)
			if dccid != 0 && cidErr == nil && dusernm != "" {
				fmt.Println("cid: ", dccid)
				dUsr := h.Manager.DeleteUser(dusernm, dccid)
				fmt.Println("dUsr: ", dUsr)
				var rtn Response
				if dUsr {
					rtn.Success = dUsr
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
