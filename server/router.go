package server

import "github.com/gin-gonic/gin"

func (s *Server) newRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery(), CORSMiddleware(), gin.Logger())

	test := router.Group("/test")
	test.POST("/appkey", s.GenerateAppKeySignature)

	v1 := router.Group("/v1")
	v1.Use(authMiddlewareAppKey(s.appKeyValidator))

	v1.POST("/recognize_text", s.RecognizeTextAuto)
	v1.POST("/open_ocr", s.OpenOcr)
	v1.POST("/generate_img", s.GenerateImg)

	return router
}
