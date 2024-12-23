package auth

import "fmt"

const (
	userSession = "user-session:%v"
)

func RedisKeyUserSession[T uint | string](userID T) string {
	return fmt.Sprintf(userSession, userID)
}
