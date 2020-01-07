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
