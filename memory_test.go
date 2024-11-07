package go_cache

import (
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	c, _ := NewMemory()
	c.Put("test", 1, NoExpiration)
	time.Sleep(time.Second * 5)
	if c.GetInt("test") != 1 {
		t.Fail()
	}
}

func TestExpiration(t *testing.T) {
	c, _ := NewMemory()
	c.Put("test", 1, time.Second)
	time.Sleep(time.Second * 5)
	if c.GetInt("test") == 1 {
		t.Fail()
	}
}
