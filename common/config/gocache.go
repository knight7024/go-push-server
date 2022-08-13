package config

import (
	"github.com/patrickmn/go-cache"
	"time"
)

const (
	DefaultExpirationTime = 7 * 24 * time.Hour // 7 days
	PurgeInterval         = 24 * time.Hour     // 1 day
)

var MemCache *cache.Cache

func InitCache() {
	MemCache = cache.New(DefaultExpirationTime, PurgeInterval)
}
