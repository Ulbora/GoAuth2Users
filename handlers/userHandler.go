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
	auscl.Role = h.getRole(r)
	auscl.URL = addRoleURL
	auscl.Scope = "write"
	//fmt.Println("client: ", h.Validator)
	auth := h.ValidatorClient.Authorize(r, &auscl, h.getValidationURL())
	h.SetContentType(w)
	if auth {
		aasURIContOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", aasURIContOk)
		if !aasURIContOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var us db.User
			aussuc, uaerr := h.ProcessBody(r, &us)
			h.Log.Debug("aussuc: ", aussuc)
			h.Log.Debug("us: ", us)
			h.Log.Debug("uaerr: ", uaerr)
			if !aussuc && uaerr != nil {
				http.Error(w, uaerr.Error(), http.StatusBadRequest)
			} else {
				var cidStr string
				if auscl.Role != superAdmin {
					cidStr = r.Header.Get("clientId")
					cid, _ := strconv.ParseInt(cidStr, 10, 64)
					us.ClientID = cid
				}
				ausSuc := h.Manager.AddUser(&us)
				h.Log.Debug("ausSuc: ", ausSuc)
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
	h.Log.Debug("inside UpdateUser------------------------------------")

	var upsAuURL = "/ulbora/rs/user/update"

	var uprlcl jv.Claim
	uprlcl.Role = h.getRole(r)
	uprlcl.URL = upsAuURL
	uprlcl.Scope = "write"
	//fmt.Println("client: ", h.Client)
	auth := h.ValidatorClient.Authorize(r, &uprlcl, h.getValidationURL())
	h.SetContentType(w)
	h.Log.Debug("auth: ", auth)
	if auth {
		// w.Header().Set("Content-Type", "application/json")
		uPasURIContOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", uPasURIContOk)
		if !uPasURIContOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var upus db.User
			upusbsuc, upuserr := h.ProcessBody(r, &upus)
			h.Log.Debug("upusbsuc: ", upusbsuc)
			h.Log.Debug("upus: ", upus)
			h.Log.Debug("upuserr: ", upuserr)
			if !upusbsuc && upuserr != nil {
				http.Error(w, upuserr.Error(), http.StatusBadRequest)
			} else {
				var cidStr string
				if uprlcl.Role != superAdmin {
					cidStr = r.Header.Get("clientId")
					cid, _ := strconv.ParseInt(cidStr, 10, 64)
					upus.ClientID = cid
				}
				var upussuc bool
				if upus.Password != "" {
					h.Log.Debug("in password ")
					upussuc = h.Manager.UpdateUserPassword(&upus)
				} else if upus.FirstName != "" && upus.LastName != "" && upus.Email != "" && upus.RoleID != 0 {
					h.Log.Debug("in info ")
					upussuc = h.Manager.UpdateUserInfo(&upus)
				} else {
					h.Log.Debug("in enabled ")
					upussuc = h.Manager.UpdateUserEnabled(&upus)
				}
				h.Log.Debug("upussuc: ", upussuc)
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
	h.Log.Debug("getAuURL: ", getAuURL)
	var gusrcl jv.Claim
	gusrcl.Role = h.getRole(r)
	gusrcl.URL = getAuURL
	gusrcl.Scope = "read"
	//fmt.Println("client: ", h.Client)
	auth := h.ValidatorClient.Authorize(r, &gusrcl, h.getValidationURL())
	h.SetContentType(w)
	if auth {
		//var id string
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			var cidStr string
			if gusrcl.Role == superAdmin {
				cidStr = vars["clientId"]
			} else {
				cidStr = r.Header.Get("clientId")
			}
			var usernm = vars["username"]
			h.Log.Debug("vars: ", vars)
			cid, cidErr := strconv.ParseInt(cidStr, 10, 64)
			if cid != 0 && cidErr == nil && usernm != "" {
				h.Log.Debug("cid: ", cid)
				getUsr := h.Manager.GetUser(usernm, cid)
				h.Log.Debug("getUsr: ", getUsr)
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

//GetUserList GetUserList
func (h *UserHandler) GetUserList(w http.ResponseWriter, r *http.Request) {
	var getAuURL = "/ulbora/rs/user/list"

	var gusrlcl jv.Claim
	gusrlcl.Role = superAdmin
	gusrlcl.URL = getAuURL
	gusrlcl.Scope = "read"
	//fmt.Println("client: ", h.Client)
	auth := h.ValidatorClient.Authorize(r, &gusrlcl, h.getValidationURL())
	h.SetContentType(w)
	if auth {
		getUsrl := h.Manager.GetUserList()
		h.Log.Debug("getUsrl: ", getUsrl)
		w.WriteHeader(http.StatusOK)
		resJSON, _ := json.Marshal(getUsrl)
		fmt.Fprint(w, string(resJSON))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//SearchUserList SearchUserList
func (h *UserHandler) SearchUserList(w http.ResponseWriter, r *http.Request) {
	var getAuURL = "/ulbora/rs/user/search"

	var susrcl jv.Claim
	susrcl.Role = h.getRole(r)
	susrcl.URL = getAuURL
	susrcl.Scope = "read"
	//fmt.Println("client: ", h.Client)
	auth := h.ValidatorClient.Authorize(r, &susrcl, h.getValidationURL())
	h.SetContentType(w)
	if auth {
		//var id string
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) != 0 {
			var cidStr string
			if susrcl.Role == superAdmin {
				cidStr = vars["clientId"]
			} else {
				cidStr = r.Header.Get("clientId")
			}
			//var usernm = vars["username"]
			h.Log.Debug("vars: ", vars)
			cid, cidErr := strconv.ParseInt(cidStr, 10, 64)
			if cid != 0 && cidErr == nil {
				h.Log.Debug("cid: ", cid)
				sUsr := h.Manager.SearchUserList(cid)
				h.Log.Debug("sUsr: ", sUsr)
				w.WriteHeader(http.StatusOK)
				resJSON, _ := json.Marshal(sUsr)
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

//DeleteUser DeleteUser
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var getAuURL = "/ulbora/rs/user/delete"

	var dusrcl jv.Claim
	dusrcl.Role = h.getRole(r)
	dusrcl.URL = getAuURL
	dusrcl.Scope = "write"
	//fmt.Println("client: ", h.Client)
	auth := h.ValidatorClient.Authorize(r, &dusrcl, h.getValidationURL())
	h.SetContentType(w)
	if auth {
		//var id string
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			var cidStr string
			if dusrcl.Role == superAdmin {
				cidStr = vars["clientId"]
			} else {
				cidStr = r.Header.Get("clientId")
			}
			var dusernm = vars["username"]
			h.Log.Debug("vars: ", vars)
			dcid, cidErr := strconv.ParseInt(cidStr, 10, 64)
			if dcid != 0 && cidErr == nil && dusernm != "" {
				h.Log.Debug("cid: ", dcid)
				dUsr := h.Manager.DeleteUser(dusernm, dcid)
				h.Log.Debug("dUsr: ", dUsr)
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

// add pupulate User to db.User
