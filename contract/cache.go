package cache

import (
	"context"
	"time"
)

type Cache interface {
	Driver
}

type Driver interface {
	Add(key string, value any, t time.Duration) bool
	Decrement(key string, value ...int64) (int64, error)
	Forever(key string, value any) bool
	Forget(key string) bool
	Flush() bool
	Get(key string, def ...any) any
	GetBool(key string, def ...bool) bool
	GetInt(key string, def ...int) int
	GetInt64(key string, def ...int64) int64
	GetString(key string, def ...string) string
	Has(key string) bool
	Increment(key string, value ...int64) (int64, error)
	Lock(key string, t ...time.Duration) Lock
	Put(key string, value any, t time.Duration) error
	Pull(key string, def ...any) any
	Remember(key string, ttl time.Duration, callback func() (any, error)) (any, error)
	RememberForever(key string, callback func() (any, error)) (any, error)
	WithContext(ctx context.Context) Driver
}

type Lock interface {
	Block(t time.Duration, callback ...func()) bool
	Get(callback ...func()) bool
	Release() bool
	ForceRelease() bool
}
