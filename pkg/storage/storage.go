package storage

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate mockgen -source=./storage.go -destination=./mocks/storage_mocks.go -package=mocks

import (
	"time"

	"github.com/go-redis/redis"
)

// Redis - хранилище ключ-значение
type Redis interface {
	Ping() error
	Close() error
	Get(string) *redis.StringCmd
	Set(string, interface{}, time.Duration) *redis.StatusCmd
}