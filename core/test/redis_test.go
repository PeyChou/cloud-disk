package test

import (
	"context"
	"github.com/redis/go-redis/v9"
	"testing"
	"time"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func TestSetValue(t *testing.T) {
	err := rdb.Set(ctx, "handsome", "peychou", time.Second*10).Err()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetValue(t *testing.T) {
	val, err := rdb.Get(ctx, "handsome").Result()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(val)
}
