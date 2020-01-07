//Package handlers ...
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	jv "github.com/Ulbora/GoAuth2JwtValidator"
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
