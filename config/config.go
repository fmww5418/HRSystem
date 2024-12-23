package config

import (
	"os"
	"regexp"
	"sync"

	"github.com/jinzhu/configor"
)

var Config = struct {
	ServerPort    int    `json:"server_port" env:"server_port" default:"8080" yaml:"serverPort"`
	DBHost        string `json:"db_host" env:"db_host" default:"db" yaml:"dbHost"`
	DBPort        int    `json:"db_port" env:"db_port" default:"3306" yaml:"dbPort"`
	DBUsername    string `json:"db_username" env:"db_username" default:"root" yaml:"dbUsername"`
	DBPassword    string `json:"db_password" env:"db_password" default:"password" yaml:"dbPassword"`
	DBName        string `json:"db_name" env:"db_name" default:"hr_system" yaml:"dbName"`
	RedisHost     string `json:"redis_host" env:"redis_host" default:"redis" yaml:"redisHost"`
	RedisPort     int    `json:"redis_port" env:"redis_port" default:"6379" yaml:"redisPort"`
	RedisPassword string `json:"redis_password" env:"redis_password" default:"" yaml:"redisPassword"`
	JWTSecret     []byte `json:"jwt_secret" env:"jwt_secret" default:"93f554c87a3d93e40259d4096648aef16ceec090ddac6bf6a16cecdb58423f657d432b046cd706191a99fb7c2fe5078254e673acdf13e78c49b0d9b8ac230cb25e842de11002fcfa0d1cabe4ba742cd9fb2efb2a07ade3d863f17ddfd83a2bb43db1aedd1365459cc369c51ab6a52cb5c7e38c7faafe5738ce1037f63e3bdb0839eaa8d05e4ac8a36acca3e239a7abdeb09e0e5169734c894c2dd6241e461707a330f4cd63be25bd633fb05007cd50b50a0791eb28322f6df07674eaf9a6c453374ee5d4967dd9c8b1f065bf4827b8ee5102f66b384eeb3ab27ba0d08cf87bc33b66c91bad8c3793b5b52beed6ffd2bb85e6903f175b0a31cf7534d34449146a" yaml:"jwtSecret"`
}{}

var (
	once sync.Once
)

func LoadConfig(filepath *string) {
	once.Do(func() {
		loadConfig(filepath)
	})
}

var (
	testRegexp  = regexp.MustCompile("_test|(\\.test$)")
	defaultPath = "config/config.json"
)

func loadConfig(filepath *string) {
	if filepath == nil {
		filepath = &defaultPath
	}

	if testRegexp.MatchString(os.Args[0]) {
		if err := configor.Load(&Config); err != nil {
			panic(err)
		}
	} else {
		if err := configor.Load(&Config, *filepath); err != nil {
			panic(err)
		}
	}
}
