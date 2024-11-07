package go_cache

import (
	"context"
	"testing"
	"time"
)

func TestRedisGet(t *testing.T) {
	c, err := NewRedis(context.Background(), &Config{
		Username: "",
		Password: "",
		Database: 1,
		Host:     "127.0.0.1",
		Port:     "6379",
	}, "test")

	if err != nil {
		t.Fail()
	}

	c.Put("test", 1, NoExpiration)
	time.Sleep(time.Second * 3)

	if c.GetInt("test", 2) != 1 {
		t.Fail()
	}
}
