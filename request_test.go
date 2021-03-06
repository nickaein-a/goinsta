package goinsta

import (
	"os"
	"testing"
)

func TestRequest(t *testing.T) {
	username := os.Getenv("INSTA_USERNAME")
	password := os.Getenv("INSTA_PASSWORD")
	insta := New(username, password)
	err := insta.Login()
	if err != nil {
		t.Fatal(err)
		return
	}

	_, err = insta.sendRequest("accounts/logout/", "", false)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log("status : ok")
}
