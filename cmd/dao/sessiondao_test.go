package dao

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"bookstore/cmd/model"
)

func TestSession(t *testing.T) {
	t.Run("Create session validation", testCreateSession)
	t.Run("Delete session validation", testDeleteSession)
	t.Run("Get session by ID validation", testGetSessionByID)
	t.Run("Is logged in validation", testIsLoggedIn)
}

func testCreateSession(t *testing.T) {
	session := &model.Session{
		SessionID: "12345678901234567890",
		Username:  "user1",
		UserID:    2,
	}

	if err := CreateSession(session); err != nil {
		t.Fatalf("CreateSession failed: %v", err)
	}
}

func testDeleteSession(t *testing.T) {
	sessionID := "12345678901234567890"

	if err := DeleteSession(sessionID); err != nil {
		t.Fatalf("DeleteSession failed: %v", err)
	}
}

func testGetSessionByID(t *testing.T) {
	sessionID := "18de498b-995b-46fe-ab1f-8ea15841b39d"

	session, err := GetSessionByID(sessionID)
	if err != nil {
		t.Fatalf("GetSessionByID failed: %v", err)
	}
	fmt.Printf("Information for session: %v", session)
}

func testIsLoggedIn(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	r.AddCookie(&http.Cookie{Name: "user", Value: "18de498b-995b-46fe-ab1f-8ea15841b39d"})

	loggedIn, session := IsLoggedIn(r)
	if !loggedIn {
		t.Fatal("Expected user to be logged in, but got false")
	}
	if session == nil {
		t.Fatal("Expected session to be non-nil, but got nil")
	}

	expectedUsername := "Bruce"
	if session.Username != expectedUsername {
		t.Fatalf("Expected username %s, but got %s", expectedUsername, session.Username)
	}
}
