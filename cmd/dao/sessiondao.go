package dao

import (
	"database/sql"
	"fmt"
	"net/http"

	"bookstore/cmd/model"
	"bookstore/cmd/utils"
)

func CreateSession(session *model.Session) error {
	const query = "INSERT INTO sessions VALUES (?, ?, ?)"
	stmt, err := utils.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("Failed to prepare statement: %v", err)
	}

	if _, err := stmt.Exec(session.SessionID, session.Username, session.UserID); err != nil {
		return fmt.Errorf("Failed to execute statement: %v", err)
	}
	defer stmt.Close()
	return nil
}

func DeleteSession(sessionID string) error {
	const query = "DELETE FROM sessions WHERE session_id = ?"

	if _, err := utils.DB.Exec(query, sessionID); err != nil {
		return fmt.Errorf("Failed to execute statement: %v", err)
	}
	return nil
}

func GetSessionByID(sessionID string) (*model.Session, error) {
	const query = "SELECT session_id, username, user_id FROM sessions WHERE session_id = ?"
	row := utils.DB.QueryRow(query, sessionID)

	session := &model.Session{}
	err := row.Scan(&session.SessionID, &session.Username, &session.UserID)
	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("Session not found")
	case err != nil:
		return nil, fmt.Errorf("Failed to scan session row: %v", err)
	}

	return session, nil
}

func IsLoggedIn(r *http.Request) (bool, *model.Session) {
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		session, _ := GetSessionByID(cookieValue)
		if session != nil && session.UserID > 0 {
			return true, session
		}
	}
	return false, nil
}
