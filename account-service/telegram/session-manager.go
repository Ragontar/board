package telegram

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"

	// tg "github.com/xelaj/mtproto/telegram"
	telegram "github.com/gotd/td/telegram"
	auth "github.com/gotd/td/telegram/auth"
	tg "github.com/gotd/td/tg"
)

const (
	MTPROTO_SERVER_ADDR = "149.154.167.40:443"
	INIT_WARN_CHANNEL   = true
)

var (
	ENV_APP_ID          string
	ENV_APP_HASH        string
	ENV_PUBLIC_KEY_FILE string
	ENV_STORAGE_PATH    string
)

type TelegramSession struct {
	UserID   string
	Phone    string
	Code     string
	CodeHash string
	Storage  RedisStorage
	CodeChan chan (string)
	// SessionFile     string
	AccountPassword *tg.AccountPassword

	TelegramClient *telegram.Client
}

func NewTelegramSession(userID string, phone string) (*TelegramSession, error) {
	if userID == "" || phone == "" {
		return &TelegramSession{}, fmt.Errorf("error: cannot create session with \nUser:%v \nPhone: %v", userID, phone)
	}
	ts := TelegramSession{
		UserID: userID,
		Phone:  phone,
	}

	if err := ts.NewClient(); err != nil {
		return &TelegramSession{}, err
	}

	return &ts, nil
}

func (t *TelegramSession) NewClient() error {
	t.Storage = RedisStorage{
		SessionID: uuid.NewString(),
	}

	println("NewClient")
	client := telegram.NewClient(
		appIDtoInt(),
		ENV_APP_HASH,
		telegram.Options{
			SessionStorage: &t.Storage,
		},
	)

	println("NewClientFinished")
	t.TelegramClient = client
	return nil
}

func (t *TelegramSession) RefreshClient() {
	t.TelegramClient = telegram.NewClient(
		appIDtoInt(),
		ENV_APP_HASH,
		telegram.Options{
			SessionStorage: &t.Storage,
		},
	)
}

func appIDtoInt() int {
	appid, err := strconv.Atoi(ENV_APP_ID)
	if err != nil {
		panic("CANNOT CONVERT APPID TO INTEGER")
	}

	return appid
}

func (t *TelegramSession) SendCode() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)

	println("SendCode")
	println(t.Phone, ENV_APP_HASH, ENV_APP_ID)

	go func(c context.Context) {
		t.TelegramClient.Run(c, func(ctx context.Context) error {
			println("RUN")
			defer cancel()
			authSentCode, err := t.TelegramClient.Auth().SendCode(c, t.Phone, auth.SendCodeOptions{})
			if err != nil {
				println(err)
				return err
			}

			t.CodeHash = authSentCode.PhoneCodeHash
			println("CODE HASH", t.CodeHash)
			t.CodeChan = make(chan string)
			t.Code = <-t.CodeChan

			_, err = t.TelegramClient.Auth().SignIn(c, t.Phone, t.Code, t.CodeHash)
			if err != nil {
				if err == auth.ErrPasswordAuthNeeded {
					log.Println("PASSWORD REQUIRED - NOT IMPLEMENTED")
				}
				println(err.Error())
				return err
			}
			println("OUT RUN CONFIRM")
			return nil
		})
	}(ctx)

	return nil

	// return nil
}

func (t *TelegramSession) ConfirmCode(code string) (*tg.AuthAuthorization, error) {
	println("CONFIRM")
	t.CodeChan <- code
	println("OUT CONFIRM")
	return nil, nil
}

// func (t *TelegramSession) ConfirmPassword(pwd string) (tg.AuthAuthorization, error) {
// 	inputCheck, err := telegram.GetInputCheckPassword(pwd, t.AccountPassword)
// 	if err != nil {
// 		return nil, err
// 	}

// 	auth, err := t.TelegramClient.AuthCheckPassword(inputCheck)

// 	return auth, err
// }

// func (t *TelegramSession) GetAccountPassword() error {
// 	var err error
// 	t.AccountPassword, err = t.TelegramClient.AccountGetPassword()
// 	return err
// }

// SaveSession saves TelegramSession to redis storage and returns SessionID key.
// func (t *TelegramSession) SaveSession() (string, error) {

// 	t.TelegramClient.Auth()
// 	return sessionId, status.Err()
// }
