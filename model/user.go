package model

import (
	"database/sql"
	"errors"
	"eventManagement/db"
	"eventManagement/utils"
	"fmt"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u User) Save() error {
	query := `
		INSERT INTO users(email, password)
		VALUES(?,?)`

	prepare, err := db.DB.Prepare(query)

	if err != nil {
		panic(fmt.Sprintf("insert into usera error: %s", err))
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
	u.ID = id
	return err
}

func (u *User) ValidateUserCredential() error {
	query := `SELECT id, password FROM useres WHERE email = ?`

	var retrievedPassword string
	row := db.DB.QueryRow(query, u.Email)
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return err
	}

	isValidate := utils.CompareCredential(retrievedPassword, u.Password)

	if !isValidate {
		return errors.New("invalid credential")
	}

	return nil
}
