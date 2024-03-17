package pkg

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type Conn struct {
	Client *redis.Client
}

const (
	cacheNil string = `redis: nil`
)

func NewClient() *Conn {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password
		DB:       0,                // Default database
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		// panic(errors.Wrap(err, "redis connection"))
	}
	return &Conn{
		Client: rdb,
	}
}

func (r *Conn) Set(key string, value interface{}, exp time.Duration) error {
	ctx := context.Background()
	err := r.Client.Set(ctx, key, value, exp)
	if err != nil {
		return err.Err()
	}
	return nil
}

func (r *Conn) Get(ctx context.Context, key string) ([]byte, error) {
	cmd := r.Client.Get(ctx, key)
	fmt.Println("CMD ==> ", cmd)
	b, e := cmd.Bytes()

	if e != nil {
		if e.Error() == cacheNil {
			return b, nil
		}
	}

	return b, e
}

func (r *Conn) Close() error {
	return r.Client.Close()
}
