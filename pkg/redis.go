package pkg

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
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

	REDIS_ADDR := os.Getenv("REDIS_ADDR")
	REDIS_PASSWORD := os.Getenv("REDIS_PASSWORD")
	REDIS_DBstr := os.Getenv("REDIS_DB")
	REDIS_DB, err := strconv.Atoi(REDIS_DBstr)
	if err != nil {
		log.Fatal("Error converting REDIS_DB to an integer:", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     REDIS_ADDR,     //"localhost:6379", // Redis server address
		Password: REDIS_PASSWORD, // No password
		DB:       REDIS_DB,       // Default database
	})

	_, err = rdb.Ping(context.Background()).Result()
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
