package model

import (
	"database/sql"
	"errors"
	"eventManagement/db"
	"eventManagement/utils"
	"fmt"
)

type User struct {
	UserId   int64  `json:"user_id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u User) Save() error {
	query := `
		INSERT INTO users(email, password)
		VALUES(?,?)`

	prepare, err := db.DB.Prepare(query)

	if err != nil {
		panic(fmt.Sprintf("insert into users error: %s", err))
	}

	hashedPassword, err := utils.ToHashPassword(u.Password)

	if err != nil {
		panic(fmt.Sprintf("hashing password error: %s", err))
	}

	result, err := prepare.Exec(u.Email, hashedPassword)

	if err != nil {
		panic(fmt.Sprintf("exec create user error: %s", err))
	}

	defer func(prepare *sql.Stmt) {
		err := prepare.Close()
		if err != nil {
			panic(fmt.Sprintf("close users table error: %s", err))
		}
	}(prepare)

	id, _ := result.LastInsertId()
	u.UserId = id
	return err
}

func (u *User) ValidateUserCredential() error {
	query := `SELECT id, password FROM users WHERE email = ?`

	var retrievedPassword string
	row := db.DB.QueryRow(query, u.Email)
	err := row.Scan(&u.UserId, &retrievedPassword)

	if err != nil {
		return err
	}

	isValidate := utils.CompareCredential(retrievedPassword, u.Password)

	if !isValidate {
		return errors.New("invalid credential")
	}

	return nil
}

func ValidateUserByEmail(email string) (*User, error) {
	query := `SELECT email, user_id FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, email)

	var user User
	err := row.Scan(&user)

	if err != nil {
		return nil, err
	}

	if user.Email != "" {
		return nil, errors.New("email is not available")
	}

	return &user, nil
}

func ResetUserPassword(userId int64, newPassword string) error {
	query := `UPDATE users SET password = ? WHERE  user_id = ?`
	prepare, err := db.DB.Prepare(query)

	if err != nil {
		panic(fmt.Sprintf("prepare update user table error: %s", err))
	}

	defer func(prepare *sql.Stmt) {
		err := prepare.Close()
		if err != nil {
			panic(fmt.Sprintf("close events table error: %s", err))
		}
		return
	}(prepare)

	_, err = prepare.Exec(prepare, userId, newPassword)
	return err
}
