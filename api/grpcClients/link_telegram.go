package grpcclients

import (
	"api/grpc/AccountService"
	"api/models"
	"context"
	"fmt"
)

func (asc *AccountServiceClient) LinkTelegram(ts models.TelegramSession, userID string) (models.TelegramConfirmationCode, error) {
	tcc := models.TelegramConfirmationCode{}
	if ts.Phone == nil {
		return tcc, fmt.Errorf("error: no phone number provided")
	}

	// var opts grpc.CallOption
	ctx, cancel := context.WithTimeout(context.Background(), timeoutAuth)
	defer cancel()

	resp, err := asc.client.LinkTelegram(
		ctx,
		&AccountService.TelegramSession{
			Phone: *ts.Phone,
			User: &AccountService.User{
				UserId: userID,
			},
		})

	tcc.PhoneCode = &resp.PhoneCode
	tcc.Key = &resp.Key

	return tcc, err
}

func (asc *AccountServiceClient) LinkTelegramConfirmCode(tcc models.TelegramConfirmationCode) (models.LinkStatus, error) {
	if tcc.Key == nil || tcc.PhoneCode == nil {
		return models.LinkStatus{}, fmt.Errorf("error: no key or confirmation code provided")
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeoutAuth)
	defer cancel()
	resp, err := asc.client.LinkTelegramConfirmCode(ctx, &AccountService.TelegramConfirmationCode{
		PhoneCode: *tcc.PhoneCode,
		Key: *tcc.Key,
	})
	if err != nil {
		return models.LinkStatus{}, err
	}

	return models.LinkStatus{
		Ok: resp.GetOk(),
		PasswordRequired: resp.GetPasswordRequired(),
	}, nil
}
