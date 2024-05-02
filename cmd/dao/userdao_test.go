package dao

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	// t.Run("Login validation", testLogin)
	// t.Run("Register validation", testRegister)
	// t.Run("Save user validation", testSave)
}

func testLogin(t *testing.T) {
	user, err := GetUserByUsernameAndPassword("Bruce", "palebluedot4")
	if err != nil {
		t.Fatalf("GetUserByUsernameAndPassword failed: %v", err)
	}
	fmt.Println("Fetched user information:", user)
}

func testRegister(t *testing.T) {
	user, err := GetUserByUsername("Bruce")
	if err != nil {
		t.Fatalf("GetUserByUsername failed: %v", err)
	}
	fmt.Println("Fetched user information:", user)
}

func testSave(t *testing.T) {
	if err := SaveUser("Bruce2", "1234", "Bruce2@gmail.com"); err != nil {
		t.Fatalf("SaveUser failed: %v", err)
	}
}
