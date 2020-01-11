//Package handlers ...
package handlers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	m "github.com/Ulbora/GoAuth2Users/managers"
)

func TestUserHandler_LoginUser(t *testing.T) {
	var uh UserHandler
	//var mc jv.MockOauthClient
	//mc.MockValidate = true
	//uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	um.MockUserLoginSuc = true
	//um.MockInsertRoleID = 12
	uh.Manager = um.GetNew()

	h := uh.GetNew()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"username":"tester", "password":"somepw", "clientId": 444}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	//r.Header.Set("clientId", "10")
	w := httptest.NewRecorder()

	h.LoginUser(w, r)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 200 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestUserHandler_LoginUserNoUser(t *testing.T) {
	var uh UserHandler
	//var mc jv.MockOauthClient
	//mc.MockValidate = true
	//uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	um.MockUserLoginSuc = true
	//um.MockInsertRoleID = 12
	uh.Manager = um.GetNew()

	h := uh.GetNew()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"username":"", "password":"somepw", "clientId": 444}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	//r.Header.Set("clientId", "10")
	w := httptest.NewRecorder()

	h.LoginUser(w, r)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 401 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestUserHandler_LoginUserNotValid(t *testing.T) {
	var uh UserHandler
	//var mc jv.MockOauthClient
	//mc.MockValidate = true
	//uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	//um.MockUserLoginSuc = true
	//um.MockInsertRoleID = 12
	uh.Manager = um.GetNew()

	h := uh.GetNew()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"username":"tester", "password":"somepw", "clientId": 444}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	//r.Header.Set("clientId", "10")
	w := httptest.NewRecorder()

	h.LoginUser(w, r)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 401 || w.Header().Get("Content-Type") != "application/json" {
		t.Fail()
	}
}

func TestUserHandler_LoginUserMidia(t *testing.T) {
	var uh UserHandler
	//var mc jv.MockOauthClient
	//mc.MockValidate = true
	//uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	um.MockUserLoginSuc = true
	//um.MockInsertRoleID = 12
	uh.Manager = um.GetNew()

	h := uh.GetNew()

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"username":"tester", "password":"somepw", "clientId": 444}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	//r.Header.Set("Content-Type", "application/json")
	//r.Header.Set("clientId", "10")
	w := httptest.NewRecorder()

	h.LoginUser(w, r)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 415 {
		t.Fail()
	}
}

func TestUserHandler_LoginUserBody(t *testing.T) {
	var uh UserHandler
	//var mc jv.MockOauthClient
	//mc.MockValidate = true
	//uh.ValidatorClient = mc.GetNewClient()
	var um m.MockUserManager
	um.MockUserLoginSuc = true
	//um.MockInsertRoleID = 12
	uh.Manager = um.GetNew()

	h := uh.GetNew()

	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"username":"tester", "password":"somepw", "clientId": 444}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", nil)
	//r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")
	//r.Header.Set("clientId", "10")
	w := httptest.NewRecorder()

	h.LoginUser(w, r)
	hd := w.Header()
	fmt.Println("code: ", w.Code)
	fmt.Println("w content type", hd.Get("Content-Type"))
	if w.Code != 400 {
		t.Fail()
	}
}
