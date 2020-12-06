//Package handlers ...
package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	lg "github.com/Ulbora/Level_Logger"
	"net/http"
	"os"
	"testing"
)

type testObj struct {
	Valid bool   `json:"valid"`
	Code  string `json:"code"`
}

func TestUserHandler_ProcessBody(t *testing.T) {
	var uh UserHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
	var robj testObj
	robj.Valid = true
	robj.Code = "3"
	// var res http.Response
	// res.Body = ioutil.NopCloser(bytes.NewBufferString(`{"valid":true, "code":"1"}`))
	var sURL = "http://localhost/test"
	aJSON, _ := json.Marshal(robj)
	r, _ := http.NewRequest("POST", sURL, bytes.NewBuffer(aJSON))
	var obj testObj
	suc, _ := uh.ProcessBody(r, &obj)
	if !suc || obj.Valid != true || obj.Code != "3" {
		t.Fail()
	}
}

func TestUserHandler_ProcessBodyBadBody(t *testing.T) {
	var uh UserHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
	var robj testObj
	robj.Valid = true
	robj.Code = "3"
	// var res http.Response
	// res.Body = ioutil.NopCloser(bytes.NewBufferString(`{"valid":true, "code":"1"}`))
	var sURL = "http://localhost/test"
	aJSON, _ := json.Marshal(robj)
	r, _ := http.NewRequest("POST", sURL, bytes.NewBuffer(aJSON))
	//var obj testObj
	suc, _ := uh.ProcessBody(r, nil)
	if suc {
		t.Fail()
	}
}

func TestUserHandler_getValidationURL(t *testing.T) {
	var uh UserHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
	os.Setenv("VALIDATION_SERVICE", "testsys")
	//var obj testObj
	url := uh.getValidationURL()
	fmt.Println("url: ", url)
	if url != "testsys" {
		t.Fail()
	}
}

func TestUserHandler_getValidationURL2(t *testing.T) {
	var uh UserHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	uh.Log = &l
	os.Unsetenv("VALIDATION_SERVICE")
	//var obj testObj
	url := uh.getValidationURL()
	fmt.Println("url: ", url)
	if url == "testsys" {
		t.Fail()
	}
}

func TestUserHandler_getRole(t *testing.T) {
	var h UserHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	h.Log = &l
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	r.Header.Set("appId", "app1")
	r.Header.Set("clientId", "50")
	r.Header.Set("role", "admin")
	role := h.getRole(r)
	if role != "app150admin" {
		t.Fail()
	}

}
