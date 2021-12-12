package big_cache_test

import (
	"testing"
	"time"

	"github.com/ssentinull/kumparan-article-service/pkg/utils/big_cache"
	"github.com/stretchr/testify/assert"
)

func TestSuccessfulGetCache(t *testing.T) {
	k, v := "key", "value"
	bc := big_cache.NewBigCache(big_cache.Config{EvictionTime: time.Duration(5)})
	bc.Set(k, []byte(v))

	c := big_cache.NewCacher(&big_cache.CacherConfig{BigCache: bc})
	repl, err := c.Get(k)

	assert.NoError(t, err)
	assert.NotNil(t, repl)
}

func TestFailedGetCache(t *testing.T) {
	bc := big_cache.NewBigCache(big_cache.Config{EvictionTime: time.Duration(5)})
	c := big_cache.NewCacher(&big_cache.CacherConfig{BigCache: bc})
	repl, err := c.Get("key")

	assert.Error(t, err)
	assert.Nil(t, repl)
}

func TestSuccessfulSetCache(t *testing.T) {
	bc := big_cache.NewBigCache(big_cache.Config{EvictionTime: time.Duration(5)})
	c := big_cache.NewCacher(&big_cache.CacherConfig{BigCache: bc})
	err := c.Put("key", []byte("value"))

	assert.NoError(t, err)
}
