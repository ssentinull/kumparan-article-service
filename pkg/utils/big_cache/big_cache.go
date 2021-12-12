package big_cache

import (
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/sirupsen/logrus"
)

type Config struct {
	EvictionTime time.Duration
}

func NewBigCache(cfg Config) *bigcache.BigCache {
	bigCache, err := bigcache.NewBigCache(bigcache.DefaultConfig(cfg.EvictionTime * time.Minute))
	if err != nil {
		logrus.Fatal(err)
	}

	return bigCache
}

type CacherConfig struct {
	BigCache *bigcache.BigCache
}

type Cacher struct {
	*CacherConfig
}

func NewCacher(cfg *CacherConfig) *Cacher {
	return &Cacher{cfg}
}

func (c *Cacher) Get(key string) ([]byte, error) {
	entry, err := c.BigCache.Get(key)
	if err != nil {
		return nil, err
	}

	return entry, nil
}

func (c *Cacher) Put(key string, entry []byte) error {
	err := c.BigCache.Set(key, entry)
	if err != nil {
		return err
	}

	return nil
}
