package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"git.corp.zgcszkw.com/authhub/config"
	"git.corp.zgcszkw.com/authhub/service"
	"git.corp.zgcszkw.com/authhub/token"
)

// Server 为服务提供 http 请求
type Server struct {
	// store           db.Store //  使用接口，方便 mock
	tokenMaker      token.Maker
	appKeyValidator *token.AppKeyValidator
	config          *config.Config // 配置文件
	httpServer      *http.Server
	service         *service.Service
}

// NewServer 创建一个 HTTP server 并提供所有 API 路由
func NewServer(config *config.Config, service *service.Service) (*Server, error) {
	if config.AppKey == "" || config.AppKeySecret == "" {
		return nil, fmt.Errorf("app key and app key secret must be provided")
	}
	keyMap := map[string]string{
		config.AppKey: config.AppKeySecret,
	}
	appKeyValidator, err := token.NewAppKeyValidator(config.AppKeyDuration, keyMap)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %s", err.Error())
	}

	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %s", err.Error())
	}

	server := &Server{
		tokenMaker:      tokenMaker,
		config:          config,
		appKeyValidator: appKeyValidator,
		service:         service,
	}

	// 注册验证器
	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	// 	v.RegisterValidation("currency", valiadCurrency)
	// }

	router := server.newRouter()
	server.swaggerRouter(router)

	server.setRouter(router)
	return server, nil
}

func (s *Server) setRouter(router *gin.Engine) {
	s.httpServer = &http.Server{
		Handler: router,
	}
}

// Start 在指定的地址启动服务
func (s *Server) Start(addr string) error {
	s.httpServer.Addr = addr
	return s.httpServer.ListenAndServe()
}

// 在指定的地址启动服务
func (s *Server) GracefulShutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

// errorResponse 返回错误
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
