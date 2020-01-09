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
	m "github.com/Ulbora/GoAuth2Users/managers"
	"github.com/gorilla/mux"
)

func TestUserHandler_AddUser(t *testing.T) {
	var uh UserHandler
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	um.MockInsertUserSuc = true
	//um.MockInsertRoleID = 12
	uh.Manager = um.GetNew()

	h := uh.GetNew()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"username":"tester", "password":"somepw","enabled":true, "emailAddress":"tester11@tester.com","firstName":"tester","lastName":"tester", "roleId": 4, "clientId": 444}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddUser(w, r)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestUserHandler_AddUserNotAuth(t *testing.T) {
	var uh UserHandler
	var mc jv.MockOauthClient
	//mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	um.MockInsertUserSuc = true
	//um.MockInsertRoleID = 12
	uh.Manager = um.GetNew()

	h := uh.GetNew()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"username":"tester", "password":"somepw","enabled":true, "emailAddress":"tester11@tester.com","firstName":"tester","lastName":"tester", "roleId": 4, "clientId": 444}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddUser(w, r)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 401 {
		t.Fail()
	}
}

func TestUserHandler_AddUserFail(t *testing.T) {
	var uh UserHandler
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	//um.MockInsertUserSuc = true
	//um.MockInsertRoleID = 12
	uh.Manager = um.GetNew()

	h := uh.GetNew()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"username":"tester", "password":"somepw","enabled":true, "emailAddress":"tester11@tester.com","firstName":"tester","lastName":"tester", "roleId": 4, "clientId": 444}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddUser(w, r)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 500 {
		t.Fail()
	}
}

func TestUserHandler_AddUserBadBody(t *testing.T) {
	var uh UserHandler
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	um.MockInsertUserSuc = true
	//um.MockInsertRoleID = 12
	uh.Manager = um.GetNew()

	h := uh.GetNew()

	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"username":"tester", "password":"somepw","enabled":true, "emailAddress":"tester11@tester.com","firstName":"tester","lastName":"tester", "roleId": 4, "clientId": 444}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", nil)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddUser(w, r)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 400 {
		t.Fail()
	}
}

func TestUserHandler_AddUserMedia(t *testing.T) {
	var uh UserHandler
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	um.MockInsertUserSuc = true
	//um.MockInsertRoleID = 12
	uh.Manager = um.GetNew()

	h := uh.GetNew()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"username":"tester", "password":"somepw","enabled":true, "emailAddress":"tester11@tester.com","firstName":"tester","lastName":"tester", "roleId": 4, "clientId": 444}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.AddUser(w, r)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 415 {
		t.Fail()
	}
}

func TestUserHandler_UpdateUserPw(t *testing.T) {
	var uh UserHandler
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	um.MockUpdatePasswordSuc = true
	//um.MockInsertRoleID = 12
	uh.Manager = um.GetNew()

	h := uh.GetNew()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"username":"tester", "password":"somepw","enabled":true, "emailAddress":"tester11@tester.com","firstName":"tester","lastName":"tester", "roleId": 4, "clientId": 444}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.UpdateUser(w, r)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestUserHandler_UpdateUserPwNoAuth(t *testing.T) {
	var uh UserHandler
	var mc jv.MockOauthClient
	//mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	um.MockUpdatePasswordSuc = true
	//um.MockInsertRoleID = 12
	uh.Manager = um.GetNew()

	h := uh.GetNew()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"username":"tester", "password":"somepw","enabled":true, "emailAddress":"tester11@tester.com","firstName":"tester","lastName":"tester", "roleId": 4, "clientId": 444}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.UpdateUser(w, r)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 401 {
		t.Fail()
	}
}

func TestUserHandler_UpdateUserPwFailed(t *testing.T) {
	var uh UserHandler
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	//um.MockUpdatePasswordSuc = true
	//um.MockInsertRoleID = 12
	uh.Manager = um.GetNew()

	h := uh.GetNew()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"username":"tester", "password":"somepw","enabled":true, "emailAddress":"tester11@tester.com","firstName":"tester","lastName":"tester", "roleId": 4, "clientId": 444}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.UpdateUser(w, r)
	hd := w.Header()
	fmt.Println("failed code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 500 {
		t.Fail()
	}
}

func TestUserHandler_UpdateUserPwBody(t *testing.T) {
	var uh UserHandler
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	um.MockUpdatePasswordSuc = true
	//um.MockInsertRoleID = 12
	uh.Manager = um.GetNew()

	h := uh.GetNew()

	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"username":"tester", "password":"somepw","enabled":true, "emailAddress":"tester11@tester.com","firstName":"tester","lastName":"tester", "roleId": 4, "clientId": 444}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", nil)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.UpdateUser(w, r)
	hd := w.Header()
	fmt.Println("body code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 400 {
		t.Fail()
	}
}

func TestUserHandler_UpdateUserPwMedia(t *testing.T) {
	var uh UserHandler
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	um.MockUpdatePasswordSuc = true
	//um.MockInsertRoleID = 12
	uh.Manager = um.GetNew()

	h := uh.GetNew()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"username":"tester", "password":"somepw","enabled":true, "emailAddress":"tester11@tester.com","firstName":"tester","lastName":"tester", "roleId": 4, "clientId": 444}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.UpdateUser(w, r)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 415 {
		t.Fail()
	}
}

func TestUserHandler_UpdateUserInfo(t *testing.T) {
	var uh UserHandler
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	um.MockUpdateInfoSuc = true
	//um.MockInsertRoleID = 12
	uh.Manager = um.GetNew()

	h := uh.GetNew()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"username":"tester", "password":"","enabled":true, "emailAddress":"tester11@tester.com","firstName":"tester","lastName":"tester", "roleId": 4, "clientId": 444}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUt", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.UpdateUser(w, r)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestUserHandler_UpdateUserEnabled(t *testing.T) {
	var uh UserHandler
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	um.MockUpdateEnableSuc = true
	//um.MockInsertRoleID = 12
	uh.Manager = um.GetNew()

	h := uh.GetNew()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"username":"tester", "password":"","enabled":true, "emailAddress":"","firstName":"","lastName":"", "roleId": 4, "clientId": 444}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUt", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.UpdateUser(w, r)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestUserHandler_GetUser(t *testing.T) {

	var uh UserHandler
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	var usr m.User
	usr.Username = "tester"
	usr.Enabled = true
	usr.FirstName = "tester"
	um.MockUser = &usr
	uh.Manager = um.GetNew()

	h := uh.GetNew()

	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"username":"tester", "password":"somepw","enabled":true, "emailAddress":"tester11@tester.com","firstName":"tester","lastName":"tester", "roleId": 4, "clientId": 444}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"username": "tester",
		"clientId": "5",
	}
	r = mux.SetURLVars(r, vars)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.GetUser(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.User
	json.Unmarshal(body, &bdy)
	fmt.Println("bdy: ", bdy)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || bdy.Username != "tester" || !bdy.Enabled {
		t.Fail()
	}
}

func TestUserHandler_GetUserNoAuth(t *testing.T) {

	var uh UserHandler
	var mc jv.MockOauthClient
	//mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	var usr m.User
	usr.Username = "tester"
	usr.Enabled = true
	usr.FirstName = "tester"
	um.MockUser = &usr
	uh.Manager = um.GetNew()

	h := uh.GetNew()

	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"username":"tester", "password":"somepw","enabled":true, "emailAddress":"tester11@tester.com","firstName":"tester","lastName":"tester", "roleId": 4, "clientId": 444}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"username": "tester",
		"clientId": "5",
	}
	r = mux.SetURLVars(r, vars)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.GetUser(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.User
	json.Unmarshal(body, &bdy)
	fmt.Println("bdy: ", bdy)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 401 {
		t.Fail()
	}
}

func TestUserHandler_GetUserNoParam(t *testing.T) {

	var uh UserHandler
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	var usr m.User
	usr.Username = "tester"
	usr.Enabled = true
	usr.FirstName = "tester"
	um.MockUser = &usr
	uh.Manager = um.GetNew()

	h := uh.GetNew()

	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"username":"tester", "password":"somepw","enabled":true, "emailAddress":"tester11@tester.com","firstName":"tester","lastName":"tester", "roleId": 4, "clientId": 444}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"username": "tester",
		//"clientId": "5",
	}
	r = mux.SetURLVars(r, vars)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.GetUser(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.User
	json.Unmarshal(body, &bdy)
	fmt.Println("bdy: ", bdy)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 400 {
		t.Fail()
	}
}

func TestUserHandler_GetUserBadParam(t *testing.T) {

	var uh UserHandler
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	var usr m.User
	usr.Username = "tester"
	usr.Enabled = true
	usr.FirstName = "tester"
	um.MockUser = &usr
	uh.Manager = um.GetNew()

	h := uh.GetNew()

	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"username":"tester", "password":"somepw","enabled":true, "emailAddress":"tester11@tester.com","firstName":"tester","lastName":"tester", "roleId": 4, "clientId": 444}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"username": "tester",
		"clientId": "a",
	}
	r = mux.SetURLVars(r, vars)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.GetUser(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy m.User
	json.Unmarshal(body, &bdy)
	fmt.Println("bdy: ", bdy)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 400 {
		t.Fail()
	}
}
