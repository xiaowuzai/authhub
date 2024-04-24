package server

import (
	"errors"
	"net/http"
	"strings"

	"git.corp.zgcszkw.com/authhub/token"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

const (
	headerRequestId = "RequestId"
	headerAppKey    = "AppKey"
	headerCreated   = "Created"
)

func authMiddlewareAppKey(tokenValidator *token.AppKeyValidator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		fields := strings.Fields(strings.TrimPrefix(authorizationHeader, " "))
		if len(fields) < 2 {
			err := errors.New("authorization header is error")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		switch authorizationType {
		case authorizationTypeAppKey:
			requestId, key, created := getParas(ctx)
			// 解析token
			appKey := token.NewAppKey(requestId, created, key)
			err := tokenValidator.VerifyToken(appKey, fields[1])
			if err != nil {
				log.Error().
					Err(err).
					Str("requestId", requestId).
					Str("appKey", key).
					Str("created", created).
					Msgf("VerifyToken failed")
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
				return
			}

			// ctx.Set(ctxPayloadKey, payload)
		default:
			err := errors.New("authorization header type error")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		ctx.Next()
	}
}

func getParas(ctx *gin.Context) (string, string, string) {
	requestId := ctx.GetHeader("RequestId")
	appKey := ctx.GetHeader("AppKey")
	created := ctx.GetHeader("Created")
	return requestId, appKey, created
}
