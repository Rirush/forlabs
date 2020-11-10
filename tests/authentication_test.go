package tests

import (
	"github.com/Rirush/forlabs"
	"testing"
	"time"
)

func TestAuthentication(t *testing.T) {
	_, err := forlabs.Authenticate(ValidUsername, ValidPassword)
	if err != nil {
		t.Errorf("Successful login failed, %s", err)
	}
	time.Sleep(time.Second)
	_, err = forlabs.Authenticate("a", "b")
	if err == nil {
		t.Errorf("Invalid login failed")
		return
	}
	if err.Error() != "invalid password" {
		t.Errorf("Invalid login failed, %s", err)
	}
}