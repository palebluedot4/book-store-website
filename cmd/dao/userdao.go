package dao

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"bookstore/cmd/model"
	"bookstore/cmd/utils"
)

func GetUserByUsernameAndPassword(username, password string) (*model.User, error) {
	const query = "SELECT id, username, password, email FROM users WHERE username = ?"
	stmt, err := utils.DB.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("Failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(username)

	user := &model.User{}
	err = row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("User not found")
	case err != nil:
		return nil, fmt.Errorf("Failed to scan user row: %v", err)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, fmt.Errorf("Invalid password")
	}
	return user, nil
}

func GetUserByUsername(username string) (*model.User, error) {
	const query = "SELECT id, username, password, email FROM users WHERE username = ?"
	row := utils.DB.QueryRow(query, username)

	user := &model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("User not found")
	case err != nil:
		return nil, fmt.Errorf("Failed to scan user row: %v", err)
	}

	return user, nil
}

func SaveUser(username, password, email string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("Failed to hash password: %v", err)
	}

	const query = "INSERT INTO users (username, password, email) VALUES (?, ?, ?)"
	stmt, err := utils.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("Failed to prepare statement: %v", err)
	}

	if _, err := stmt.Exec(username, hashedPassword, email); err != nil {
		return fmt.Errorf("Failed to execute statement: %v", err)
	}
	defer stmt.Close()
	return nil
}
