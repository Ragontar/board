package crud

import (
	"account-service/database"
	"context"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
)

var redisKeyTTL = 14 * 24 * time.Hour

func UserRegistration(u User) (s Session, err error) {
	u, err = CreateUser(u)
	if err != nil {
		return
	}

	s.User.UserID = u.UserID
	s.User.Email = u.Email
	s.Token = uuid.New().String()

	pipe := database.GetTokenStorage().Pipeline()
	set_res := pipe.Set(
		context.Background(),
		s.Token,          // Token
		uuid.NewString(), // Secret
		redisKeyTTL,
	)
	get_res := pipe.Get(
		context.Background(),
		s.Token,
	)
	_, err = pipe.Exec(context.Background())
	if err != nil && err != redis.Nil {
		return
	}
	if set_res.Err() != nil {
		err = set_res.Err()
		return
	}
	if get_res.Err() != nil {
		err = set_res.Err()
		return
	}

	res, err := get_res.Result()
	if err != nil {
		return
	}
	uuidSecret, err := uuid.Parse(res)
	if err != nil {
		return s, err
	}
	s.Secret = uuidSecret.String()

	return s, err
}

func AuthenticateRequest(req AuthRequest) AuthResponse {
	return AuthResponse{
		Authenticated: req.CheckPayload(),
	}
}

func AuthenticateUser(u User) (s Session, err error) {
	s.User, err = SelectUserByCreds(u)
	if err != nil {
		return
	}
	s.User.Credentials = ""
	s.Token = uuid.New().String()

	pipe := database.GetTokenStorage().Pipeline()
	set_res := pipe.Set(
		context.Background(),
		s.Token,          // Token
		uuid.NewString(), // Secret
		redisKeyTTL,
	)
	get_res := pipe.Get(
		context.Background(),
		s.Token,
	)
	_, err = pipe.Exec(context.Background())
	if err != nil && err != redis.Nil {
		return
	}
	if set_res.Err() != nil {
		err = set_res.Err()
		return
	}
	if get_res.Err() != nil {
		err = set_res.Err()
		return
	}

	res, err := get_res.Result()
	if err != nil {
		return
	}
	uuidSecret, err := uuid.Parse(res)
	if err != nil {
		return s, err
	}
	s.Secret = uuidSecret.String()

	return s, err
}
