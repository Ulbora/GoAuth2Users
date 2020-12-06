//Package handlers ...
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

//LoginUser LoginUser
func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {

	h.SetContentType(w)
	logURIContOk := h.CheckContent(r)
	h.Log.Debug("conOk: ", logURIContOk)
	if !logURIContOk {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		var us db.User
		valsuc, uaerr := h.ProcessBody(r, &us)
		h.Log.Debug("valsuc: ", valsuc)
		h.Log.Debug("us: ", us)
		h.Log.Debug("uaerr: ", uaerr)
		if !valsuc && uaerr != nil {
			http.Error(w, uaerr.Error(), http.StatusBadRequest)
		} else {
			var rtn LoginResponse
			if us.Username != "" && us.Password != "" && us.ClientID != 0 {
				valid := h.Manager.ValidateUser(us.Username, us.Password, us.ClientID)
				h.Log.Debug("valid: ", valid)
				if valid {
					rtn.Valid = valid
					rtn.Code = strconv.FormatInt(us.ClientID, 10)
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusUnauthorized)
				}
			} else {
				w.WriteHeader(http.StatusUnauthorized)
			}
			resJSON, _ := json.Marshal(rtn)
			fmt.Fprint(w, string(resJSON))
		}
	}
}
