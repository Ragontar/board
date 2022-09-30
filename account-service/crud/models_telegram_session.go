package crud

import (
	"account-service/database"
	"context"
)

type TelegramSession struct {
	SessionID string
	UserID    string
}

func (ts TelegramSession) Insert(ctx context.Context) error {
	if ts.UserID == "" || ts.SessionID == "" {
		return &ErrorBadInsertData{}
	}

	row, err := database.GetDB().Query(
		ctx,
		SQL.INSERT_TELEGRAM_SESSION,
		ts.SessionID,
		ts.UserID,
	)
	if err != nil {
		return err
	}
	defer row.Close()
	if !row.Next() {
		return &ErrorBadInsertData{}
	}

	return nil
}

func (ts TelegramSession) DeleteById(ctx context.Context) error {
	if ts.SessionID == "" {
		return &ErrorBadInsertData{}
	}

	row, err := database.GetDB().Query(
		ctx,
		SQL.DELETE_TELEGRAM_SESSION_BY_ID,
		ts.SessionID,
	)
	if err != nil {
		return err
	}
	defer row.Close()
	if !row.Next() {
		return &ErrorBadInsertData{}
	}

	return row.Scan(&ts)
}