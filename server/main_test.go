package server

import (
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"

	"git.corp.zgcszkw.com/authhub/config"
	"git.corp.zgcszkw.com/authhub/util"
)

// newTestServer 创建一个测试服务器
func newTestServer(t *testing.T) *Server {
	config := &config.Config{
		TokenSymmetricKey: util.NewPasetoSymmetricKey(),
		TokenDuration:     time.Minute,
		AppKey:            "test_key",
		AppKeySecret:      "test_secret",
		AppKeyDuration:    time.Duration(600),
	}

	server, err := NewServer(config, nil)
	require.NoError(t, err)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
