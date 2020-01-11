//Package handlers ...
package handlers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	jv "github.com/Ulbora/GoAuth2JwtValidator"
	m "github.com/Ulbora/GoAuth2Users/managers"
)

func TestUserHandler_ClientAddUser(t *testing.T) {
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
	r.Header.Set("clientId", "10")
	w := httptest.NewRecorder()

	h.ClientAddUser(w, r)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestUserHandler_ClientAddUserNotAuth(t *testing.T) {
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
	r.Header.Set("clientId", "10")
	w := httptest.NewRecorder()

	h.ClientAddUser(w, r)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 401 {
		t.Fail()
	}
}

func TestUserHandler_ClientAddUserFail(t *testing.T) {
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
	r.Header.Set("clientId", "10")
	w := httptest.NewRecorder()

	h.ClientAddUser(w, r)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 500 {
		t.Fail()
	}
}

func TestUserHandler_ClientAddUserClientHeader(t *testing.T) {
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
	//r.Header.Set("clientId", "10")
	w := httptest.NewRecorder()

	h.ClientAddUser(w, r)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 400 {
		t.Fail()
	}
}

func TestUserHandler_ClientAddUserBadBody(t *testing.T) {
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
	r.Header.Set("clientId", "10")
	w := httptest.NewRecorder()

	h.ClientAddUser(w, r)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 400 {
		t.Fail()
	}
}

func TestUserHandler_ClientAddUserMedia(t *testing.T) {
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
	r.Header.Set("clientId", "10")
	w := httptest.NewRecorder()

	h.ClientAddUser(w, r)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 415 {
		t.Fail()
	}
}

func TestUserHandler_ClientUpdateUserPw(t *testing.T) {
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
	r.Header.Set("clientId", "10")
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.ClientUpdateUser(w, r)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestUserHandler_ClientUpdateUserPwClientId(t *testing.T) {
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
	//r.Header.Set("clientId", "10")
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.ClientUpdateUser(w, r)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 400 {
		t.Fail()
	}
}
func TestUserHandler_ClientUpdateUserPwNotAuth(t *testing.T) {
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
	r.Header.Set("clientId", "10")
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.ClientUpdateUser(w, r)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 401 {
		t.Fail()
	}
}

func TestUserHandler_ClientUpdateUserPwFail(t *testing.T) {
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
	r.Header.Set("clientId", "10")
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.ClientUpdateUser(w, r)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 500 {
		t.Fail()
	}
}

func TestUserHandler_ClientUpdateUserPwBody(t *testing.T) {
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
	r.Header.Set("clientId", "10")
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.ClientUpdateUser(w, r)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 400 {
		t.Fail()
	}
}

func TestUserHandler_ClientUpdateUserPwMidia(t *testing.T) {
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
	r.Header.Set("clientId", "10")
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.ClientUpdateUser(w, r)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 415 {
		t.Fail()
	}
}

func TestUserHandler_ClientUpdateUserInfo(t *testing.T) {
	var uh UserHandler
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	um.MockUpdatePasswordSuc = true
	//um.MockInsertRoleID = 12
	uh.Manager = um.GetNew()

	h := uh.GetNew()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"username":"tester", "password":"","enabled":true, "emailAddress":"tester11@tester.com","firstName":"tester","lastName":"tester", "roleId": 4, "clientId": 444}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)
	r.Header.Set("clientId", "10")
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.ClientUpdateUser(w, r)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 500 {
		t.Fail()
	}
}

func TestUserHandler_ClientUpdateUserEnabled(t *testing.T) {
	var uh UserHandler
	var mc jv.MockOauthClient
	mc.MockValidate = true
	uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	um.MockUpdatePasswordSuc = true
	//um.MockInsertRoleID = 12
	uh.Manager = um.GetNew()

	h := uh.GetNew()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"username":"tester", "password":"","enabled":true, "emailAddress":"","firstName":"","lastName":"", "roleId": 4, "clientId": 444}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("PUT", "/ffllist", aJSON)
	r.Header.Set("clientId", "10")
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.ClientUpdateUser(w, r)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 500 {
		t.Fail()
	}
}
