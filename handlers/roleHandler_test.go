//Package handlers ...
package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	jv "github.com/Ulbora/GoAuth2JwtValidator"
	db "github.com/Ulbora/GoAuth2Users/db"
	m "github.com/Ulbora/GoAuth2Users/managers"
	"github.com/gorilla/mux"
)

func TestUserHandler_AddRole(t *testing.T) {
	var uh UserHandler
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	um.MockInsertRoleSuc = true
	um.MockInsertRoleID = 12
	uh.Manager = um.GetNew()

	h := uh.GetNew()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":0, "role":"test"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddRole(w, r)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestUserHandler_AddRoleNotAuth(t *testing.T) {
	var uh UserHandler
	var mc jv.MockOauthClient
	mc.MockValidate = false
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	um.MockInsertRoleSuc = true
	um.MockInsertRoleID = 12
	uh.Manager = um.GetNew()

	h := uh.GetNew()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":0, "role":"test"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddRole(w, r)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 401 {
		t.Fail()
	}
}

func TestUserHandler_AddRoleFailAdd(t *testing.T) {
	var uh UserHandler
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	um.MockInsertRoleSuc = false
	um.MockInsertRoleID = 0
	uh.Manager = um.GetNew()

	h := uh.GetNew()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":0, "role":"test"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddRole(w, r)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 500 {
		t.Fail()
	}
}

func TestUserHandler_AddRoleBadBody(t *testing.T) {
	var uh UserHandler
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	um.MockInsertRoleSuc = true
	um.MockInsertRoleID = 12
	uh.Manager = um.GetNew()

	h := uh.GetNew()

	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":0, "role":"test"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", nil)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddRole(w, r)
	hd := w.Header()
	fmt.Println("w content type", hd.Get("Content-Type"))
	fmt.Println("code: ", w.Code)
	if w.Code != 400 {
		t.Fail()
	}
}

func TestUserHandler_AddRoleBadMedia(t *testing.T) {
	var uh UserHandler
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	um.MockInsertRoleSuc = true
	um.MockInsertRoleID = 12
	uh.Manager = um.GetNew()

	h := uh.GetNew()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":0, "role":"test"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddRole(w, r)
	hd := w.Header()
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 415 {
		t.Fail()
	}
}

func TestUserHandler_GetRole(t *testing.T) {
	var uh UserHandler
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	//um.MockInsertRoleSuc = true
	//um.MockInsertRoleID = 12
	uh.Manager = um.GetNew()
	var rl db.Role
	rl.ID = 5
	rl.Role = "test"
	um.MockRole = &rl

	h := uh.GetNew()

	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"id": "5",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	h.GetRole(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy db.Role
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	fmt.Println("code: ", w.Code)
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || bdy.ID == 0 {
		t.Fail()
	}
}

func TestUserHandler_GetRoleVars(t *testing.T) {
	var uh UserHandler
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	//um.MockInsertRoleSuc = true
	//um.MockInsertRoleID = 12
	uh.Manager = um.GetNew()
	var rl db.Role
	rl.ID = 5
	rl.Role = "test"
	um.MockRole = &rl

	h := uh.GetNew()

	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		//"id": "5",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	h.GetRole(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy db.Role
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	fmt.Println("code: ", w.Code)
	if w.Code != 400 {
		t.Fail()
	}
}

func TestUserHandler_GetRoleNotAuth(t *testing.T) {
	var uh UserHandler
	var mc jv.MockOauthClient
	//mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	//um.MockInsertRoleSuc = true
	//um.MockInsertRoleID = 12
	uh.Manager = um.GetNew()
	var rl db.Role
	rl.ID = 5
	rl.Role = "test"
	um.MockRole = &rl

	h := uh.GetNew()

	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"id": "5",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	h.GetRole(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy db.Role
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	fmt.Println("code: ", w.Code)
	if w.Code != 401 {
		t.Fail()
	}
}

func TestUserHandler_GetRoleBadParam(t *testing.T) {
	var uh UserHandler
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	//um.MockInsertRoleSuc = true
	//um.MockInsertRoleID = 12
	uh.Manager = um.GetNew()
	var rl db.Role
	rl.ID = 5
	rl.Role = "test"
	um.MockRole = &rl

	h := uh.GetNew()

	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"id": "a",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	h.GetRole(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy db.Role
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	fmt.Println("code: ", w.Code)
	if w.Code != 400 {
		t.Fail()
	}
}

func TestUserHandler_GetRoleList(t *testing.T) {
	var uh UserHandler
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	//um.MockInsertRoleSuc = true
	//um.MockInsertRoleID = 12
	uh.Manager = um.GetNew()
	var rl db.Role
	rl.ID = 5
	rl.Role = "test"
	var rlist = []db.Role{rl}

	um.MockRoleList = &rlist

	h := uh.GetNew()
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	w := httptest.NewRecorder()
	h.GetRoleList(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy []db.Role
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	fmt.Println("len(bdy): ", len(bdy))
	fmt.Println("code: ", w.Code)
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || len(bdy) != 1 {
		t.Fail()
	}

}

func TestUserHandler_GetRoleListNotAuth(t *testing.T) {
	var uh UserHandler
	var mc jv.MockOauthClient
	//mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	//um.MockInsertRoleSuc = true
	//um.MockInsertRoleID = 12
	uh.Manager = um.GetNew()
	var rl db.Role
	rl.ID = 5
	rl.Role = "test"
	var rlist = []db.Role{rl}

	um.MockRoleList = &rlist

	h := uh.GetNew()
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	w := httptest.NewRecorder()
	h.GetRoleList(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy []db.Role
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	fmt.Println("len(bdy): ", len(bdy))
	fmt.Println("code: ", w.Code)
	if w.Code != 401 {
		t.Fail()
	}

}

func TestUserHandler_DeleteRole(t *testing.T) {
	var uh UserHandler
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	//um.MockInsertRoleSuc = true
	//um.MockInsertRoleID = 12
	uh.Manager = um.GetNew()
	um.MockDeleteRoleSuc = true

	h := uh.GetNew()

	r, _ := http.NewRequest("DELETE", "/ffllist", nil)
	vars := map[string]string{
		"id": "5",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	h.DeleteRole(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	fmt.Println("Code: ", w.Code)
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || bdy.Success != true {
		t.Fail()
	}
}

func TestUserHandler_DeleteRoleNoAuth(t *testing.T) {
	var uh UserHandler
	var mc jv.MockOauthClient
	//mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	//um.MockInsertRoleSuc = true
	//um.MockInsertRoleID = 12
	uh.Manager = um.GetNew()
	um.MockDeleteRoleSuc = true

	h := uh.GetNew()

	r, _ := http.NewRequest("DELETE", "/ffllist", nil)
	vars := map[string]string{
		"id": "5",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	h.DeleteRole(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	fmt.Println("Code: ", w.Code)
	if w.Code != 401 {
		t.Fail()
	}
}

func TestUserHandler_DeleteRoleNoParm(t *testing.T) {
	var uh UserHandler
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	//um.MockInsertRoleSuc = true
	//um.MockInsertRoleID = 12
	uh.Manager = um.GetNew()
	um.MockDeleteRoleSuc = true

	h := uh.GetNew()

	r, _ := http.NewRequest("DELETE", "/ffllist", nil)
	vars := map[string]string{
		//"id": "5",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	h.DeleteRole(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	fmt.Println("Code: ", w.Code)
	if w.Code != 400 {
		t.Fail()
	}
}

func TestUserHandler_DeleteRoleBadParm(t *testing.T) {
	var uh UserHandler
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	//um.MockInsertRoleSuc = true
	//um.MockInsertRoleID = 12
	uh.Manager = um.GetNew()
	um.MockDeleteRoleSuc = true

	h := uh.GetNew()

	r, _ := http.NewRequest("DELETE", "/ffllist", nil)
	vars := map[string]string{
		"id": "a",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	h.DeleteRole(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	fmt.Println("Code: ", w.Code)
	if w.Code != 400 {
		t.Fail()
	}
}

func TestUserHandler_DeleteRoleFail(t *testing.T) {
	var uh UserHandler
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	//um.MockInsertRoleSuc = true
	//um.MockInsertRoleID = 12
	uh.Manager = um.GetNew()
	//um.MockDeleteRoleSuc = true

	h := uh.GetNew()

	r, _ := http.NewRequest("DELETE", "/ffllist", nil)
	vars := map[string]string{
		"id": "5",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	h.DeleteRole(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy Response
	json.Unmarshal(body, &bdy)
	fmt.Println("body: ", string(body))
	fmt.Println("Code: ", w.Code)
	if w.Code != 500 {
		t.Fail()
	}
}
