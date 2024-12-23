//go:build unit

package auth

import (
	"HRSystem/config"
	rlib "HRSystem/src/lib/redis"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	config.LoadConfig(nil)

	rlib.Init()
	os.Exit(m.Run())
}
