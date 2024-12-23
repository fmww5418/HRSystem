//go:build unit

package redis

import (
	"github.com/alicebob/miniredis/v2"
	goredis "github.com/go-redis/redis/v8"
)

type (
	CacheClient struct {
		*goredis.Client
	}
)

var (
	StubServer *miniredis.Miniredis
	DB         *CacheClient
)

func Init() *goredis.Client {
	var err error
	StubServer, err = miniredis.Run()
	if err != nil {
		panic(err)
	}

	DB = new(CacheClient)
	DB.Client = goredis.NewClient(&goredis.Options{
		Addr: StubServer.Addr(),
	})

	return DB.Client
}
