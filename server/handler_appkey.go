package server

import (
	"net/http"

	"git.corp.zgcszkw.com/authhub/token"
	"github.com/gin-gonic/gin"
)

type SignatureHeader struct {
	AppKey    string `json:"appKey" binding:"required"`
	Created   string `json:"created" binding:"required"`
	RequestId string `json:"requestId" binding:"required"`
}

// 测试用： 生成签名
func (s *Server) GenerateAppKeySignature(c *gin.Context) {
	req := &SignatureHeader{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	appKey := token.NewAppKey(req.RequestId, req.Created, req.AppKey)
	signature, err := s.appKeyValidator.CreateToken(appKey)
	if err != nil {
		ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	type Signature struct {
		Signature string `json:"secret"`
	}

	ResponseSuccess(c, &Signature{
		Signature: signature,
	})
}
