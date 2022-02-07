package lock

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

func TestNewLock(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	l := NewLock(client, "test")
	err := l.TryLock(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	err = l.TryLock(context.Background())
	if err != nil && !errors.Is(err, ErrLockFailed) {
		t.Fatal(err)
	}
	time.Sleep(ttl)
	err = l.Unlock(context.Background())
	if err != nil {
		t.Fatal(err)
	}
}
