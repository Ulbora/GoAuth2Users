//Package managers ...
package managers

import (
	"fmt"
	"testing"
)

var pw = "tester123"
var hpw1 string
var hpw2 string

func TestUserManager_hashPassword(t *testing.T) {
	var m UserManager
	suc, hpw := m.hashPassword(pw)
	hpw1 = hpw
	fmt.Println("hashed pw: ", hpw)
	valid := m.validatePassword(pw, hpw)
	if !suc || hpw == "" || !valid {
		t.Fail()
	}
}

func TestUserManager_hashPassword2(t *testing.T) {
	var m UserManager
	suc, hpw := m.hashPassword(pw)
	hpw2 = hpw
	fmt.Println("hashed pw: ", hpw)
	valid := m.validatePassword(pw, hpw)
	if !suc || hpw == "" || !valid {
		t.Fail()
	}
}

func TestUserManager_validatePassword(t *testing.T) {
	var m UserManager
	fmt.Println("validating: ", hpw1)
	valid := m.validatePassword(pw, hpw1)
	if !valid {
		t.Fail()
	}
}

func TestUserManager_validatePassword2(t *testing.T) {
	var m UserManager
	fmt.Println("validating: ", hpw2)
	valid := m.validatePassword(pw, hpw2)
	if !valid {
		t.Fail()
	}
}
