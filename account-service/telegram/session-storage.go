package telegram

import (
	"account-service/database"
	"context"
	"log"
	"sync"

	redis "github.com/go-redis/redis/v9"
	"github.com/pkg/errors"
)

type RedisStorage struct {
	SessionID string
	mux       sync.Mutex
}

func (r *RedisStorage) LoadSession(ctx context.Context) ([]byte, error) {
	if r == nil {
		return nil, errors.New("nil session storage is invalid")
	}

	r.mux.Lock()
	defer r.mux.Unlock()

	status := database.GetSessionStorage().Get(ctx, r.SessionID)
	session, err := status.Result()
	if err != redis.Nil {
		println("LoadSession err", err.Error())
		return nil, err
	}

	log.Printf("Loaded: %s", session)

	return []byte(session), nil
}

func (r *RedisStorage) StoreSession(ctx context.Context, data []byte) error {
	if r == nil {
		return errors.New("nil session storage is invalid")
	}

	log.Printf("Storing: %s", string(data))

	r.mux.Lock()
	defer r.mux.Unlock()
	// r.SessionID = uuid.New().String()
	log.Printf("\n\n\n Storing to key %s\n\n\nContent: %s\n\n\n", r.SessionID, string(data))
	status := database.GetSessionStorage().Set(ctx, r.SessionID, string(data), 0)
	err := status.Err()
	if err != nil {
		print("StoreSession err: ", err.Error())
	}
	// if err != redis.Nil {
	// 	println("StoreSession err", err)
	// 	return err
	// }
	return nil
}
