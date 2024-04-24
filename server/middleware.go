package server

import (
	"errors"
	"net/http"
	"strings"

	"git.corp.zgcszkw.com/authhub/token"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "Authorization"
	ctxPayloadKey           = "ctx_payload"
	authorizationTypeBearer = "bearer"
	authorizationTypeAppKey = "appkey"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ResponseError(ctx, http.StatusUnauthorized, err.Error())
			return
		}

		fields := strings.Fields(strings.TrimPrefix(authorizationHeader, " "))
		if len(fields) < 2 {
			err := errors.New("authorization header is error")
			ResponseError(ctx, http.StatusUnauthorized, err.Error())
			return
		}

		authorizationType := strings.ToLower(fields[0])
		switch authorizationType {
		case authorizationTypeBearer:
			// 解析token
			payload, err := tokenMaker.VerifyToken(fields[1])
			if err != nil {
				err := errors.New("authorization parse error")
				ResponseError(ctx, http.StatusUnauthorized, err.Error())
				return
			}

			ctx.Set(ctxPayloadKey, payload)
		default:
			err := errors.New("authorization header type error")
			ResponseError(ctx, http.StatusUnauthorized, err.Error())
			return
		}

		// 解析token
		// payload, err := tokenMaker.VerifyToken(fields[1])
		// if err != nil {
		// 	err := errors.New("authorization parse error")
		// 	ResponseError(ctx, http.StatusUnauthorized, err.Error())
		// 	return
		// }

		// ctx.Set(ctxPayloadKey, payload)
		ctx.Next()
	}
}
