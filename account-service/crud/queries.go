package crud

import (
	"account-service/database"
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
)

// CreateUser inserts User into database returning User with generated UserID field on success. Also returns
// ErrorBadInsertData error in case of empty strings in Email or Credentials fields.
// Returns ErrorAlreadyRegistered if email or credentials are not unique.
func CreateUser(u User) (User, error) {
	if u.Email == "" || u.Credentials == "" {
		return u, &ErrorBadInsertData{}
	}
	u.UserID = uuid.New().String()

	row, err := database.GetDB().Query(
		context.Background(),
		SQL.INSERT_USER,
		u.UserID,
		u.Email,
		u.Credentials,
	)
	if err != nil {
		return u, err
	}
	defer row.Close()
	if !row.Next() {
		return u, &ErrorAlreadyRegistered{}
	}

	return u, err
}

func DeleteUserByID(u User) (User, error) {
	returnedUser := User{}
	if u.UserID == uuid.Nil.String() {
		return u, &ErrorBadInsertData{}
	}

	row, err := database.GetDB().Query(
		context.Background(),
		SQL.DELETE_USER_BY_ID,
		u.UserID,
	)
	if err != nil {
		return u, err
	}
	defer row.Close()

	log.Println(row.RawValues())
	if !row.Next() {
		return returnedUser, fmt.Errorf("row.Next error")
	}
	err = row.Scan(&returnedUser.UserID, &returnedUser.Email, &returnedUser.Credentials)

	return returnedUser, err
}

func SelectUserByCreds(u User) (User, error) {
	if u.UserID == uuid.Nil.String() {
		return u, &ErrorBadInsertData{}
	}

	row, err := database.GetDB().Query(
		context.Background(),
		SQL.SELECT_USER_BY_CREDS,
		u.Credentials,
	)
	if err != nil {
		return u, err
	}
	defer row.Close()

	if !row.Next() {
		return u, fmt.Errorf("row.Next error or user not found")
	}
	err = row.Scan(&u.UserID, &u.Email, &u.Credentials)

	return u, err
}
