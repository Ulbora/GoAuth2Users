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
	lg "github.com/Ulbora/Level_Logger"
	"github.com/gorilla/mux"
)

func TestUserHandler_AddUser(t *testing.T) {
	var uh UserHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
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

func TestUserHandler_AddUser2(t *testing.T) {
	var uh UserHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
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
	r.Header.Set("appId", "app1")
	r.Header.Set("clientId", "50")
	r.Header.Set("role", "admin")
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
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
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
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
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
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
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
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
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
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
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

func TestUserHandler_UpdateUserPw2(t *testing.T) {
	var uh UserHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
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
	r.Header.Set("appId", "app1")
	r.Header.Set("clientId", "50")
	r.Header.Set("role", "admin")
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
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
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
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
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
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
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
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
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
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
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
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
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
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
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

func TestUserHandler_GetUser2(t *testing.T) {

	var uh UserHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
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
	r.Header.Set("appId", "app1")
	r.Header.Set("clientId", "50")
	r.Header.Set("role", "admin")
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
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
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
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
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
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
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

func TestUserHandler_GetUserList(t *testing.T) {

	var uh UserHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	var usr m.UserList
	usr.Username = "tester"
	usr.Enabled = true
	usr.FirstName = "tester"

	var usrlst = []m.UserList{usr}
	um.MockUserList = &usrlst

	uh.Manager = um.GetNew()

	h := uh.GetNew()

	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"username":"tester", "password":"somepw","enabled":true, "emailAddress":"tester11@tester.com","firstName":"tester","lastName":"tester", "roleId": 4, "clientId": 444}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	// vars := map[string]string{
	// 	"username": "tester",
	// 	"clientId": "5",
	// }
	// r = mux.SetURLVars(r, vars)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.GetUserList(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy []m.UserList
	json.Unmarshal(body, &bdy)
	fmt.Println("bdy: ", bdy)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || bdy[0].Username != "tester" || !bdy[0].Enabled {
		t.Fail()
	}
}

func TestUserHandler_GetUserListNotAuth(t *testing.T) {

	var uh UserHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
	var mc jv.MockOauthClient
	//mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	var usr m.UserList
	usr.Username = "tester"
	usr.Enabled = true
	usr.FirstName = "tester"

	var usrlst = []m.UserList{usr}
	um.MockUserList = &usrlst

	uh.Manager = um.GetNew()

	h := uh.GetNew()

	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"username":"tester", "password":"somepw","enabled":true, "emailAddress":"tester11@tester.com","firstName":"tester","lastName":"tester", "roleId": 4, "clientId": 444}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	// vars := map[string]string{
	// 	"username": "tester",
	// 	"clientId": "5",
	// }
	// r = mux.SetURLVars(r, vars)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.GetUserList(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy []m.UserList
	json.Unmarshal(body, &bdy)
	fmt.Println("bdy: ", bdy)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 401 {
		t.Fail()
	}
}

func TestUserHandler_SearchUserList(t *testing.T) {

	var uh UserHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	var usr m.UserList
	usr.Username = "tester"
	usr.Enabled = true
	usr.FirstName = "tester"

	var usrlst = []m.UserList{usr}
	um.MockUserList = &usrlst

	uh.Manager = um.GetNew()

	h := uh.GetNew()

	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"username":"tester", "password":"somepw","enabled":true, "emailAddress":"tester11@tester.com","firstName":"tester","lastName":"tester", "roleId": 4, "clientId": 444}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		//"username": "tester",
		"clientId": "5",
	}
	r = mux.SetURLVars(r, vars)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.SearchUserList(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy []m.UserList
	json.Unmarshal(body, &bdy)
	fmt.Println("bdy: ", bdy)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || bdy[0].Username != "tester" || !bdy[0].Enabled {
		t.Fail()
	}
}

func TestUserHandler_SearchUserList2(t *testing.T) {

	var uh UserHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	var usr m.UserList
	usr.Username = "tester"
	usr.Enabled = true
	usr.FirstName = "tester"

	var usrlst = []m.UserList{usr}
	um.MockUserList = &usrlst

	uh.Manager = um.GetNew()

	h := uh.GetNew()

	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"username":"tester", "password":"somepw","enabled":true, "emailAddress":"tester11@tester.com","firstName":"tester","lastName":"tester", "roleId": 4, "clientId": 444}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		//"username": "tester",
		"clientId": "5",
	}
	r = mux.SetURLVars(r, vars)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("appId", "app1")
	r.Header.Set("clientId", "50")
	r.Header.Set("role", "admin")
	w := httptest.NewRecorder()

	h.SearchUserList(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy []m.UserList
	json.Unmarshal(body, &bdy)
	fmt.Println("bdy: ", bdy)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || bdy[0].Username != "tester" || !bdy[0].Enabled {
		t.Fail()
	}
}

func TestUserHandler_SearchUserListNotAuth(t *testing.T) {

	var uh UserHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
	var mc jv.MockOauthClient
	//mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	var usr m.UserList
	usr.Username = "tester"
	usr.Enabled = true
	usr.FirstName = "tester"

	var usrlst = []m.UserList{usr}
	um.MockUserList = &usrlst

	uh.Manager = um.GetNew()

	h := uh.GetNew()

	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"username":"tester", "password":"somepw","enabled":true, "emailAddress":"tester11@tester.com","firstName":"tester","lastName":"tester", "roleId": 4, "clientId": 444}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		//"username": "tester",
		"clientId": "5",
	}
	r = mux.SetURLVars(r, vars)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.SearchUserList(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy []m.UserList
	json.Unmarshal(body, &bdy)
	fmt.Println("bdy: ", bdy)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 401 {
		t.Fail()
	}
}

func TestUserHandler_SearchUserListNoParam(t *testing.T) {

	var uh UserHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	var usr m.UserList
	usr.Username = "tester"
	usr.Enabled = true
	usr.FirstName = "tester"

	var usrlst = []m.UserList{usr}
	um.MockUserList = &usrlst

	uh.Manager = um.GetNew()

	h := uh.GetNew()

	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"username":"tester", "password":"somepw","enabled":true, "emailAddress":"tester11@tester.com","firstName":"tester","lastName":"tester", "roleId": 4, "clientId": 444}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		//"username": "tester",
		//"clientId": "5",
	}
	r = mux.SetURLVars(r, vars)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.SearchUserList(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy []m.UserList
	json.Unmarshal(body, &bdy)
	fmt.Println("bdy: ", bdy)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 400 {
		t.Fail()
	}
}

func TestUserHandler_SearchUserListBadParam(t *testing.T) {

	var uh UserHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	var usr m.UserList
	usr.Username = "tester"
	usr.Enabled = true
	usr.FirstName = "tester"

	var usrlst = []m.UserList{usr}
	um.MockUserList = &usrlst

	uh.Manager = um.GetNew()

	h := uh.GetNew()

	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"username":"tester", "password":"somepw","enabled":true, "emailAddress":"tester11@tester.com","firstName":"tester","lastName":"tester", "roleId": 4, "clientId": 444}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		//"username": "tester",
		"clientId": "5a",
	}
	r = mux.SetURLVars(r, vars)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.SearchUserList(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy []m.UserList
	json.Unmarshal(body, &bdy)
	fmt.Println("bdy: ", bdy)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 400 {
		t.Fail()
	}
}

func TestUserHandler_DeleteUser(t *testing.T) {

	var uh UserHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	um.MockDeleteUserSuc = true
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

	h.DeleteUser(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy Response
	json.Unmarshal(body, &bdy)
	fmt.Println("bdy: ", bdy)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || !bdy.Success {
		t.Fail()
	}
}

func TestUserHandler_DeleteUser2(t *testing.T) {

	var uh UserHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	um.MockDeleteUserSuc = true
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
	r.Header.Set("appId", "app1")
	r.Header.Set("clientId", "50")
	r.Header.Set("role", "admin")
	w := httptest.NewRecorder()

	h.DeleteUser(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy Response
	json.Unmarshal(body, &bdy)
	fmt.Println("bdy: ", bdy)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" || !bdy.Success {
		t.Fail()
	}
}

func TestUserHandler_DeleteUserNoAuth(t *testing.T) {

	var uh UserHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
	var mc jv.MockOauthClient
	//mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	um.MockDeleteUserSuc = true
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

	h.DeleteUser(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy Response
	json.Unmarshal(body, &bdy)
	fmt.Println("bdy: ", bdy)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 401 {
		t.Fail()
	}
}

func TestUserHandler_DeleteUserNoParam(t *testing.T) {

	var uh UserHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	um.MockDeleteUserSuc = true
	uh.Manager = um.GetNew()

	h := uh.GetNew()

	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"username":"tester", "password":"somepw","enabled":true, "emailAddress":"tester11@tester.com","firstName":"tester","lastName":"tester", "roleId": 4, "clientId": 444}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		//"username": "tester",
		//"clientId": "5",
	}
	r = mux.SetURLVars(r, vars)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.DeleteUser(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy Response
	json.Unmarshal(body, &bdy)
	fmt.Println("bdy: ", bdy)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 400 {
		t.Fail()
	}
}

func TestUserHandler_DeleteUserBadParam(t *testing.T) {

	var uh UserHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	um.MockDeleteUserSuc = true
	uh.Manager = um.GetNew()

	h := uh.GetNew()

	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"username":"tester", "password":"somepw","enabled":true, "emailAddress":"tester11@tester.com","firstName":"tester","lastName":"tester", "roleId": 4, "clientId": 444}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"username": "tester",
		"clientId": "5a",
	}
	r = mux.SetURLVars(r, vars)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.DeleteUser(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy Response
	json.Unmarshal(body, &bdy)
	fmt.Println("bdy: ", bdy)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 400 {
		t.Fail()
	}
}

func TestUserHandler_DeleteUserFail(t *testing.T) {

	var uh UserHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	//um.MockDeleteUserSuc = true
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

	h.DeleteUser(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var bdy Response
	json.Unmarshal(body, &bdy)
	fmt.Println("bdy: ", bdy)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 500 {
		t.Fail()
	}
}
